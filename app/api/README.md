# PBE API

The API serves demo email composer data and accepts mock send requests. The send endpoint does not deliver email.

## Development

From this directory, run:

```sh
go run ./src/cmd
```

The server listens on port `8080` by default. Set `PORT` to use another port.

## Endpoints

- `GET /api/emails/demo` returns the sample draft used by the frontend.
- `POST /api/emails/send` accepts a draft and a `gmail` or `resend` provider, then returns a mock accepted response.
