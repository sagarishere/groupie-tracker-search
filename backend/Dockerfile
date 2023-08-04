# Stage 1: Building the Go application
FROM golang:alpine AS builder

WORKDIR /app

# This is done separately from copying all source files to leverage Docker cache and speed up build time
COPY go.mod ./
RUN go mod download && go mod verify

# Copy all source files from your host to the current location in your image
COPY . .

# Now, let's build the application
# The reason we mount the cache folder is to speed up Docker builds
# This is a great solution if your application takes a long time to compile
RUN --mount=type=cache,target=/root/.cache/go-build go build -o /bin/server .

# Stage 2: Running the application
# We're using the alpine version as it is a much smaller image compared to the others
FROM alpine:latest

# Copy the binary from the builder stage to the production stage
COPY --from=builder /bin/server /bin/server

# Set the working directory in the container to app
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]
