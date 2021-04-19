# pull official base image
FROM golang as builder
ENV GO111MODULE=on
# Working directory
WORKDIR /app
# Copy files
COPY go.mod .
COPY go.sum .
# Install app dependencies
RUN go mod tidy
RUN go mod vendor
# Add src app
COPY . .
# Build app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app

# final stage
FROM alpine
WORKDIR /
# Copy binary from builder
COPY --from=builder /app/app /app
COPY --from=builder /app/config.toml /config.toml
RUN cat config.toml
ENTRYPOINT [ "/app" ] 
# Export grpc, metric, http ports
EXPOSE 8080