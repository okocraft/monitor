import { defineConfig } from "orval";

export default defineConfig({
    api: {
        input: {
            target: "../schema/openapi/monitor-api.yml",
        },
        output: {
            mode: "tags-split",
            target: "./src/api/api.ts",
            schemas: "src/api/model",
            clean: true,
            client: "react-query",
        },
    },
});
