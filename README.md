# RideF1

RideF1 is an early stage ride-hailing project. The repository currently contains a Go API, GORM schema models and migrations, and supporting container configuration.

## Project layout

- `app/Backend/api` contains the Gin API, authentication registration flow, database connection, middleware, and in-memory job queue.
- `app/Schema` contains the PostgreSQL models and schema migration entry point.
- `app/Backend/chat` contains a Node.js package manifest for the chat component; application source and run scripts are not present yet.
- `docker` contains supporting container definitions.

## Requirements

- Go 1.26.2
- PostgreSQL

## Configure and run

The API and schema migration read `DATABASE_URL` from the environment. The API also accepts `PORT`; it defaults to `8080`. Both commands load a `.env` file from their current working directory.

Create `app/Schema/.env` with your local PostgreSQL connection string, then run the schema migration:

```sh
cd app/Schema
go run .
```

Create `app/Backend/api/.env` with `DATABASE_URL` and optionally `PORT`, then start the API:

```sh
cd app/Backend/api
go run ./cmd
```

Use local placeholder values in `.env` and do not commit credentials.

## API routes

- `GET /health` returns the API health status.
- `POST /api/auth/register` registers a user. The JSON request requires `fullname`, `email`, `phone`, and `password` (minimum 8 characters).

Example registration body:

```json
{
  "fullname": "Sam Lee",
  "email": "sam@example.com",
  "phone": "+15551234567",
  "password": "replace-with-a-local-password"
}
```

The API listens on port `8080` by default. Set `PORT` to use another port.
