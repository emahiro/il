import { ServerRequest } from "https://deno.land/std@0.54.0/http/server.ts";
import { Status, STATUS_TEXT } from "https://deno.land/std@0.54.0/http/http_status.ts";

export async function entryHandler(req: ServerRequest): Promise<void> {
    switch (req.method) {
        case "GET":
            return await getEntryHandler(req)
        default:
            return req.respond({status: Status.MethodNotAllowed, body: STATUS_TEXT.get(Status.MethodNotAllowed)})
    }
}

async function getEntryHandler(req: ServerRequest): Promise<void> {
    return req.respond({status: Status.OK, body: "OK! This is entory handler."})
}
