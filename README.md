# RideF1

RideF1 is an early stage ride-hailing project. The repository currently contains a Go API, GORM schema models and migrations, and supporting container configuration.

## Project layout

- `app/api` contains the Gin API, authentication registration flow, database connection, middleware, and in-memory job queue.
- `Database` contains the PostgreSQL models and schema migration entry point.
- `appchat` contains a Node.js package manifest for the chat component; application source and run scripts are not present yet.
- `docker` contains supporting container definitions.

## Configure and run

The API and schema migration read `DATABASE_URL` from the environment. The API also accepts `PORT`; it defaults to `8080`. Both commands load a `.env` file from their current working directory.

Create `Database/.env` with your local PostgreSQL connection string, then run the schema migration:

```sh
cd Database
go run .
```

Create `app/api/.env` with `DATABASE_URL` and optionally `PORT`, then start the API:

```sh
cd app/api
go run ./cmd
```

Use local placeholder values in `.env` and do not commit credentials.

The API listens on port `8080` by default. Set `PORT` to use another port.
