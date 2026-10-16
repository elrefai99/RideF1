import { Tool } from '@modelcontextprotocol/sdk/types.js'

export const GET_LAST_USERS: Tool = {
     name: "get_last_users",
     description: "Fetch the most recently created users from the database",
     inputSchema: {
          type: "object",
          properties: {
               limit: {
                    type: "number",
                    description: "Maximum number of users to return (default 10, maximum 100)",
                    minimum: 1,
                    maximum: 100,
                    multipleOf: 1
               }
          },
          additionalProperties: false
     }
} as const
