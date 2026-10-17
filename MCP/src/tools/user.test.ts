import { describe, expect, it } from "vitest";
import { GET_LAST_USERS } from "./user.js";

describe("GET_LAST_USERS", () => {
     it("defines the get_last_users tool and its limit constraints", () => {
          expect(GET_LAST_USERS.name).toBe("get_last_users");
          expect(GET_LAST_USERS.inputSchema).toMatchObject({
               type: "object",
               properties: {
                    limit: {
                         type: "number",
                         minimum: 1,
                         maximum: 100,
                         multipleOf: 1
                    }
               },
               additionalProperties: false
          });
     });
});
