# Dockerfile
FROM golang:1.22.2-alpine

# mockery is a lib to generate mocks
RUN go install github.com/vektra/mockery/v2@latest

# swag is a lib to generate swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8