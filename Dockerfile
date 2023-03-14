ARG GO_VERSION=1.19.4

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

RUN CGO_ENABLED=0 go build -o /main ./cmd/api

# Multi-Stage production build
FROM golang:${GO_VERSION}-alpine AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Install timezone database
RUN apk --no-cache add tzdata

COPY .env .
COPY --from=builder /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]
