FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./

COPY internal ./internal

COPY pkg ./pkg

RUN mkdir bin

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main

FROM alpine:latest

COPY --from=0 /bin/main /bin/main

EXPOSE 1323

CMD ["/bin/main"]
