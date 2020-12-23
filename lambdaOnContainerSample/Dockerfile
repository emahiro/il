FROM public.ecr.aws/lambda/go:1.x

WORKDIR /src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
