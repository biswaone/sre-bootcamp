FROM golang:1.24.4-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/sre-bootcamp ./cmd 


FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/bin/sre-bootcamp /

EXPOSE 8080

CMD ["/sre-bootcamp"]



