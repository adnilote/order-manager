# syntax=docker/dockerfile:experimental

FROM golang:1.14-buster AS binary
RUN apt-get update && \
    apt-get install -y git ca-certificates build-essential && \
    update-ca-certificates && \
    apt clean
ARG LIBRDKAFKA_VERSION=1.2.1
RUN git clone -b v${LIBRDKAFKA_VERSION} --single-branch --depth 1 https://github.com/edenhill/librdkafka.git
WORKDIR librdkafka
RUN ./configure && \
        make && \
        make install && \
        rm -rf librdkafka
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
ENV GO111MODULE="auto"
ADD ./ $GOPATH/src/github.com/adnilote/order-manager
WORKDIR $GOPATH/src/github.com/adnilote/order-manager/app
RUN go get
RUN go build -o bin

FROM debian:buster-20190910
COPY --from=binary /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=binary /usr/lib/pkgconfig /usr/lib/pkgconfig
COPY --from=binary /usr/lib/librdkafka* /usr/lib/
COPY --from=binary /go/src/github.com/adnilote/order-manager/app/bin /app/bin
COPY --from=binary /go/src/github.com/adnilote/order-manager/config.yaml /app/config.yaml
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait
WORKDIR /app
CMD /wait && /app/bin