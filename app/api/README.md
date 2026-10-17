# RideF1 API

The Go API uses Gin and PostgreSQL. It currently exposes a health check and user registration endpoint.

## Requirements

- Go 1.26.2
- PostgreSQL with the RideF1 schema initialized

## Configuration

Create a `.env` file in this directory:

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/ridef1?sslmode=disable
REDIS_URL=redis://default:password@localhost:6379
PORT=8080
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
CLOUDINARY_UPLOAD_PRESET=your_signed_upload_preset
```

Replace the example database values with your local configuration. `PORT` is optional and defaults to `8080`. Do not commit credentials.

`CLOUDINARY_UPLOAD_PRESET` is optional for signed uploads. If you configure one, create a **signed** upload preset in Cloudinary and restrict it to JPG, JPEG, PNG, and WEBP images with a 5 MB maximum file size.

Initialize the schema from `Database` before starting the API:

```sh
cd ../../Schema
go run .
cd ..api
go run ./cmd
```

## Routes

### `GET /metrics`

Exposes Prometheus metrics for HTTP requests and server CPU, memory, and disk usage. Prometheus should scrape this endpoint every 15 seconds. Set `MONITORING_DISK_PATH` to the mount point to monitor; it defaults to `/` on Linux and `C:\` on Windows.

### `GET /health`

Returns a JSON health status.

### `POST /api/auth/register`

Creates a user account. The request requires a full name, valid email, phone number, and password with at least eight characters.

```json
{
  "fullname": "Sam Lee",
  "email": "sam@example.com",
  "phone": "+15551234567",
  "password": "replace-with-a-local-password"
}
```

Successful registration returns HTTP `201`. Invalid input or registration errors return HTTP `400`.

### Avatar upload

`POST /api/user/avatar/signature` requires authentication and returns signed upload fields for the current user's avatar. Send the image directly to the returned `upload_url` as `multipart/form-data`, including the returned `api_key`, `timestamp`, `folder`, `public_id`, `overwrite`, and `signature` fields, plus `upload_preset` if it is configured, and the image in a `file` field. Do not set the multipart `Content-Type` header manually in browser clients.

After Cloudinary responds successfully, send its `public_id`, `version`, and `signature` fields as JSON to `PUT /api/user/avatar`. The API verifies Cloudinary's response and stores the avatar URL for the authenticated user.

Browser example:

```js
async function uploadAvatar(file, token) {
  const signatureResponse = await fetch("/api/user/avatar/signature", {
    method: "POST",
    headers: { Authorization: `Bearer ${token}` },
  });
  const { data: signed } = await signatureResponse.json();

  const form = new FormData();
  form.append("file", file);
  form.append("api_key", signed.api_key);
  form.append("timestamp", String(signed.timestamp));
  form.append("folder", signed.folder);
  form.append("public_id", signed.public_id);
  if (signed.upload_preset) form.append("upload_preset", signed.upload_preset);
  form.append("overwrite", String(signed.overwrite));
  form.append("signature", signed.signature);

  const uploadResponse = await fetch(signed.upload_url, {
    method: "POST",
    body: form,
  });
  const uploaded = await uploadResponse.json();
  if (!uploadResponse.ok) throw new Error(uploaded.error?.message ?? "Cloudinary upload failed");

  const savedResponse = await fetch("/api/user/avatar", {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      public_id: uploaded.public_id,
      version: uploaded.version,
      signature: uploaded.signature,
    }),
  });
  if (!savedResponse.ok) throw new Error("Could not save avatar");

  return savedResponse.json();
}
```
