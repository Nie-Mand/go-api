FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./

COPY internal ./internal

COPY pkg ./pkg

RUN mkdir bin

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server

FROM alpine:latest

COPY --from=0 /bin/server /bin/server

EXPOSE 1323

CMD ["/bin/server"]
