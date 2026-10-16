import { Pool } from "pg";

const pool = new Pool({
     connectionString: process.env.DATABASE_URL,
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
