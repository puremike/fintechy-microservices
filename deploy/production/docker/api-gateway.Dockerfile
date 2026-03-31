FROM golang:1.26 AS builder
WORKDIR /app

# 1. Copy ONLY the dependency files first
COPY go.mod go.sum ./
# 2. Download dependencies (this layer is cached until go.mod changes)
RUN go mod download

# 3. Now copy the source code
COPY . .

# 4. Build the specific service
WORKDIR /app/services/api-gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway .

# <---- Run Stage ----->
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/services/api-gateway/api-gateway .
# Use EXPOSE to document the port for Tilt/K8s
EXPOSE 8070 
CMD ["./api-gateway"]