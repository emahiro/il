import { listenAndServe, ServerRequest} from "https://deno.land/std@0.54.0/http/server.ts";
import { Status, STATUS_TEXT } from "https://deno.land/std@0.54.0/http/http_status.ts";
import { entryHandler } from "./handler.ts";

const port = 8080;
const options = {port: port}
listenAndServe(options, (request: ServerRequest) => {
    // request.url には /pathName が入ってくる
    const s = request.url.split('?')
    if (s.length == 0) {
        return request.respond({ status: 400, body: 'Unexpected url' })
    }
    const normalizedUrl = s[0]
    // let rawQueryParams = ''
    // if (s.length >= 2){
    //     // query param あり
    //     rawQueryParams = s[1]
    // }

    // https://localhost:8080/path => ['', 'path',..]
    const re = /^\/[a-zA-Z0-9]+$/
    const r = normalizedUrl.match(re)
    if (r === null) {
        if (normalizedUrl === '/') {
            return request.respond({ status: Status.OK, body: 'OK! This is root handler!' })
        }
        return request.respond({ status: Status.NotFound, body: 'status not found: url is ' + normalizedUrl })
    }

    const path = r[0]
    switch (path) {
        case '/entries':
            return entryHandler(request)
        default:
            return request.respond({ status: Status.NotFound, body: 'status not found: url is ' + normalizedUrl})
    }
})
