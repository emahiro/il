import { serve} from "http/server.ts";

const port = 8080;
const s = serve({ port: port });
console.log("Starting server port: ", port);
for await (const req of s) {
    req.respond({ body: "Hello World\n" });
}
