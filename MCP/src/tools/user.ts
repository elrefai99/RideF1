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

export const GET_FIRST_USERS: Tool = {
     name: "get_first_users",
     description: "Fetch the most recently users created from the database",
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

export const GET_COUNT_USERS: Tool = {
     name: "get_total_users",
     description: "Fetch the most recently total users from the database",
     inputSchema: {
          type: "object",
          properties: {
          },
          additionalProperties: false
     }
} as const
