FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 go build -o shipment-calculator main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/shipment-calculator .
EXPOSE 8080

ENTRYPOINT ["./shipment-calculator"]
