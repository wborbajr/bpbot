
# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.15-alpine as builder

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
      version="V0.0.1.0" \
      description="Docker Application | bplusbot:latest"

# Copy local code to the container image.
WORKDIR /go/app
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -v -o app main.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates
# FROM gcr.io/distroless/base

WORKDIR /bplusbot

# Copy the binary to the production image from the builder stage.
COPY --from=builder go/app/app .
COPY .env .

# Run the web service on container startup.
CMD ["./app"]