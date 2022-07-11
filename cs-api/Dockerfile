FROM golang:1.18 as build

WORKDIR /app

COPY . .

RUN go build -o server main.go

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build /app/server /app/server
COPY --from=build /app/db/migrations /app/db/migrations

CMD ["./server", "server"]