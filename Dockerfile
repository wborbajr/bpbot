# 1st - Dependencies
#
FROM golang:alpine as dependencies
WORKDIR /go/app
COPY go.mod .
COPY go.sum .
RUN go mod download

# 2nd - Build
#
FROM dependencies as builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
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

# RUN mkdir /usr/local/share/ca-certificates
# COPY certs/server.crt /usr/local/share/ca-certificates
# COPY certs/server.csr /usr/local/share/ca-certificates
# COPY certs/server.key /usr/local/share/ca-certificates
# COPY certs/ca.crt /usr/local/share/ca-certificates
# COPY certs/ca.key /usr/local/share/ca-certificates
# COPY certs/ca.srl /usr/local/share/ca-certificates

RUN update-ca-certificates

# 4th - Final
#
FROM distrosetup as Final

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
    version="V0.0.1.0" \
    description="Docker Application | bplusbot:latest"

WORKDIR /bplusbot

COPY --from=distrosetup /bin/bash /bin/bash
COPY --from=distrosetup /var/spool/cron /var/spool/cron
COPY --from=distrosetup /usr/bin/crontab /usr/bin/crontab
COPY --from=distrosetup /etc/localtime /etc/localtime
COPY --from=distrosetup /etc/timezone /etc/timezone
COPY --from=distrosetup /usr/bin/htop /usr/bin/htop
COPY --from=distrosetup /usr/local/share/ca-certificates /usr/local/share/ca-certificates
#
COPY --from=builder go/app/app .
COPY --from=builder go/app/certs ./certs
#
COPY .env .

EXPOSE 3443

# Run the service on container startup.
CMD ["./app"]
