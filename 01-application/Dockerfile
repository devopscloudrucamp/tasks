FROM golang:alpine AS builder

WORKDIR /build
ADD go.mod .
ADD main.go main.go
RUN go build -o app ./main.go

FROM alpine

WORKDIR /app
COPY --from=builder /build/app /app/app
COPY public public

CMD ["./app"]