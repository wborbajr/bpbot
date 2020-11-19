# 1st - Dependencies
#
FROM golang:alpine

WORKDIR /go/app

COPY . .

RUN go mod download

EXPOSE 3443

# Run the service on container startup.
ENTRYPOINT ["go", "run", "main.go"]
# CMD ["./app"]
