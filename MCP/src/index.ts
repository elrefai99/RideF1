import './config/dotenv'
import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { CallToolRequestSchema, ListToolsRequestSchema } from "@modelcontextprotocol/sdk/types.js";
import { allUser, countAllUsers, firstUser, GET_COUNT_USERS, GET_FIRST_USERS, GET_LAST_USERS } from './gen-import';

const server = new Server({
     name: "RideF1",
     version: "1.0.0",
}, {
     capabilities: { tools: {} }
});

server.setRequestHandler(ListToolsRequestSchema, async () => ({
     tools: [
          GET_LAST_USERS,
          GET_FIRST_USERS,
          GET_COUNT_USERS,
     ]
}));

server.setRequestHandler(CallToolRequestSchema, async (request) => {
     if (request.params.name === "get_last_users") {
          const limit = Number(request.params.arguments?.limit ?? 10);

          const user = await allUser({ limit: limit })
          return {
               content: [{
                    type: "text",
                    text: JSON.stringify(user)
               }]
          };
     }
     if (request.params.name === "get_first_users") {
          const limit = Number(request.params.arguments?.limit ?? 10);

          const user = await firstUser({ limit: limit })
          return {
               content: [{
                    type: "text",
                    text: JSON.stringify(user)
               }]
          };
     }
     if (request.params.name === "get_total_users") {
          const user = await countAllUsers()
          return {
               content: [{
                    type: "text",
                    text: JSON.stringify(user)
               }]
          };
     }

     return {
          isError: true,
          content: [{
               type: "text",
               text: `Unknown tool: ${request.params.name}`
          }]
     };
});

async function main() {
     const transport = new StdioServerTransport();
     await server.connect(transport);
     console.error("MCP of F1 backend started");
}

main().catch((err: unknown) => {
     console.error(err instanceof Error ? err.message : err);
     process.exit(1);
});
