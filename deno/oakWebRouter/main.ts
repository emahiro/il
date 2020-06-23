import {
    Application,
    Router
} from "./deps.ts";
import {
    green,
    cyan,
    bold
} from "./deps.ts"

const port = 8080;
const app = new Application()

// Logger
app.use(async (ctx, next) => {
    await next();
    const rt = ctx.request.headers.get("X-Response-Time");
    console.log(`${green(ctx.request.method)} ${cyan(ctx.request.url.pathname)} - ${ bold(String(rt) )}`)
})

const r = new Router()
r
    .get("/", (ctx) => {
        ctx.response.body = "Hello oak app!"
    })
    .get("/articles", (ctx) => {
        ctx.response.body = "article handler"
    });

app.use(r.routes())
app.addEventListener("listen", ({ hostname, port }) => {
    console.log(`Staring server on ${hostname}:${port}`);
});
await app.listen({ hostname: "localhost", port: port });
console.log(bold("Close server"));
