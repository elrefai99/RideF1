# RideF1 API

The Go API uses Gin and PostgreSQL. It currently exposes a health check and user registration endpoint.

## Requirements

- Go 1.26.2
- PostgreSQL with the RideF1 schema initialized

## Configuration

Create a `.env` file in this directory:

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/ridef1?sslmode=disable
PORT=8080
```

Replace the example database values with your local configuration. `PORT` is optional and defaults to `8080`. Do not commit credentials.

Initialize the schema from `app/Schema` before starting the API:

```sh
cd ../../Schema
go run .
cd ../Backend/api
go run ./cmd
```

## Routes

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
