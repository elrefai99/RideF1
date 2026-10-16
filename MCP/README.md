# RideF1 MCP Server

This package provides a local stdio MCP server with a `get_last_users` tool. It reads `DATABASE_URL` from `.env` in the current working directory.

## Requirements

- Node.js
- pnpm
- PostgreSQL connection string in `DATABASE_URL`

## Install and run

From this directory, install dependencies and build the server:

```sh
pnpm install
pnpm build
```

Create `MCP/.env` with your database connection:

```env
DATABASE_URL=postgres://USER:PASSWORD@HOST:5432/DATABASE
```

Start the server with `pnpm start`. MCP clients should launch it directly over stdio using `node` and the built entry point. Keep stdout reserved for MCP messages; server status is written to stderr.

## Connect Codex

Add this to the Codex MCP configuration. Replace the paths with the absolute path to this repository on your machine:

```toml
[mcp_servers.ridef1]
command = "node"
args = ["E:/path/to/pbe/MCP/dist/src/index.js"]
cwd = "E:/path/to/pbe/MCP"
```

The `.env` file is loaded from `cwd`. Alternatively, set `DATABASE_URL` in the server's `env` configuration. Restart Codex and run `codex mcp list` to check that the server is registered.

## Connect Claude Code

From `MCP/`, register the server at project scope so its `.env` file is found from the working directory:

```sh
claude mcp add --scope project ridef1 -- node "E:/path/to/pbe/MCP/dist/src/index.js"
```

Start Claude Code from `MCP/` so the server finds `.env`. Approve the project MCP server when Claude Code prompts you. Use `claude mcp list` to verify it is registered.

## Available tool

- `get_last_users`: returns the latest users. The optional `limit` argument defaults to 10 and accepts integers from 1 to 100.

## Publishing

The package is configured for public npm publishing. Its current package name is `mcp`; npm requires the name to be available before publishing. Build before publishing:

```sh
pnpm build
npm publish
```
