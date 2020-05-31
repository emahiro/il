FROM emahiro/deno

EXPOSE 8080

WORKDIR /app

USER deno

COPY deps.ts .
RUN deno cache deps.ts

COPY . .

RUN deno cache main.ts

CMD ["run", "--allow-net", "main.ts"]

