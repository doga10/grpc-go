FROM golang:1.14-alpine as builder
LABEL maintainer="Douglas Dennys <douglasdennys@yahoo.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN cp .env.example .env
RUN go build ./src/framework/cmd/server.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
#COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./server"]