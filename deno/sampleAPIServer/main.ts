import { Application } from "./deps.ts";

const port:number = 8080;
const app = new Application()

app.use((ctx) => {
    ctx.response.body = "Hello oak app !!!"
})
app.addEventListener("listen", ({ hostname, port}) => {
    console.log(`Start server on ${hostname}:${port}`);
  });
await app.listen({hostname: "127.0.0.1", port: port})
console.log("Shutdown server.")
