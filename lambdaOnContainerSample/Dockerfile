FROM public.ecr.aws/lambda/go:latest

WORKDIR /src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
