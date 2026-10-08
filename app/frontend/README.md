# PBE Frontend

This Nuxt app contains the email composer interface. It loads a sample draft from the Go API and submits provider selections to a mock send endpoint. No email is delivered.

## Development

```sh
pnpm install
pnpm dev
```

Open the local URL printed by Nuxt, usually `http://localhost:3000`.

Start the API in a second terminal from `app/api` with `go run .`. Set `NUXT_PUBLIC_API_BASE` if it is not running at `http://localhost:8080`.

## Production

```sh
pnpm build
pnpm preview
```
