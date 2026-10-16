import { Pool } from "pg";

const databaseUrl = process.env.DATABASE_URL;

if (!databaseUrl) {
     throw new Error("DATABASE_URL is missing. Add it to MCP/.env or the MCP server environment.");
}

const databasePassword = new URL(databaseUrl).password;

if (!databasePassword) {
     throw new Error("DATABASE_URL has no password. Add the PostgreSQL password to MCP/.env or the MCP server environment.");
}

const pool = new Pool({
     connectionString: databaseUrl,
});

export const allUser = async (payload: { limit?: number }) => {
     const requestedLimit = Number(payload?.limit ?? 10);
     const limit = Number.isInteger(requestedLimit)
          ? Math.min(Math.max(requestedLimit, 1), 100)
          : 10;

     const result = await pool.query(
          "SELECT id, username FROM users ORDER BY created_at DESC LIMIT $1",
          [limit]
     );

     return result.rows;
};
