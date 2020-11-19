# 1st - Dependencies
#
FROM golang:alpine as dependencies
WORKDIR /go/app
COPY go.mod go.sum ./
RUN go mod download all

# 2nd - Build
#
FROM dependencies as builder
# Create the user and group files that will be used in the running
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
COPY . .
COPY --from=dependencies /go/pkg/mod/cache .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -v -o app main.go

# 3rd - Distro Setup
#
FROM alpine as distrosetup
ENV LANG=en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8
ENV TZ America/Sao_Paulo
RUN apk add --update --no-cache \
    tzdata \
    htop \
    apk-cron \
    bash \
    ca-certificates \
    && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo "${TZ}" > /etc/timezone

RUN update-ca-certificates

# 4th - Final
#
FROM distrosetup as Final

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
    version="V0.0.1.0" \
    description="Docker Application | bplusbot:latest"

WORKDIR /bplusbot

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/
#
# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=distrosetup /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=distrosetup /bin/bash /bin/bash
COPY --from=distrosetup /var/spool/cron /var/spool/cron
COPY --from=distrosetup /usr/bin/crontab /usr/bin/crontab
COPY --from=distrosetup /etc/localtime /etc/localtime
COPY --from=distrosetup /etc/timezone /etc/timezone
COPY --from=distrosetup /usr/bin/htop /usr/bin/htop
#
COPY --from=builder go/app/app .
COPY --from=builder go/app/certs ./certs
#
COPY .env .

EXPOSE 3443

# Perform any further action as an unprivileged user.
# USER nobody:nobody

# Run the service on container startup.
ENTRYPOINT ["./app"]
# CMD ["./app"]
