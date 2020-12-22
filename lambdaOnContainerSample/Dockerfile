FROM golang:1.15.6-alpine

WORKDIR /src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
