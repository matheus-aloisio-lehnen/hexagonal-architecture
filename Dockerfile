FROM golang:1.24

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

USER root
RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    sqlite3 \
    libsqlite3-dev \
 && rm -rf /var/lib/apt/lists/*

ENV GOPROXY=https://goproxy.io,direct
ENV GOSUMDB=sum.golang.org
ENV CGO_ENABLED=1

RUN go install github.com/golang/mock/mockgen@v1.6.0 \
 && go install github.com/spf13/cobra-cli@latest

RUN usermod -u 1000 www-data \
 && mkdir -p /var/www/.cache \
 && chown -R www-data:www-data /go /var/www/.cache

USER www-data

CMD ["tail", "-f", "/dev/null"]
