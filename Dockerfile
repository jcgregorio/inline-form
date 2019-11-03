# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang as builder

# Copy local code to the container image.
WORKDIR /go/src/cloudrun/server
COPY . .

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM alpine:latest

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/cloudrun/server/server /server

# Copy template
COPY index.html /index.html

# Run the web service on container startup.
CMD ["/server"]