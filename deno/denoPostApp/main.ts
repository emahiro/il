import { listenAndServe, ServerRequest} from "https://deno.land/std@0.54.0/http/server.ts";
import { Status } from "https://deno.land/std@0.54.0/http/http_status.ts";
import { entryHandler } from "./handler.ts";

const port = 8080;
const options = {port: port}
listenAndServe(options, (request: ServerRequest) => {
    const s = request.url.split('?')[0]
    if (s.length == 0) {
        request.respond({status: 400, body: 'Unexpected url'})
    }

    const normalizedUrl = s[0]
    // let rawQueryParams = ''
    // if (s.length >= 2){
    //     // query param あり
    //     rawQueryParams = s[1]
    // }

    // https://localhost:8080/:path => ['https:','','localhost:8080', ':path',..]
    const path = normalizedUrl.split('/')[3]
    switch (path) {
        case 'entries':
            entryHandler(request)
            break;
        default:
            request.respond({ status: Status.NotFound})
            break;
    }
})
