# --- Build Stage ---
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# ติดตั้ง git สำหรับ go get / mod
RUN apk add --no-cache git

# โหลด dependency
COPY go.mod go.sum ./
RUN go mod download

# คัดลอก source code
COPY . .

# Build binary
RUN go build -o app ./cmd

# --- Run Stage ---
FROM alpine:latest

WORKDIR /app

# ติดตั้ง lib ที่จำเป็น (สำหรับ SSL connection กับ Postgres)
RUN apk add --no-cache ca-certificates

# คัดลอก binary และ .env จาก builder
COPY --from=builder /app/app .
COPY --from=builder /app/.env .env

# เปิดพอร์ตที่ Fiber ใช้
EXPOSE 3001

# สั่งให้รันแอป
CMD ["./app"]
