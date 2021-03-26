FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o go-openapi-client .
EXPOSE 8080

CMD ["/go/src/app/go-openapi-client"]
