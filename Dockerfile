FROM golang:1.16-alpine as build

ARG TARGETPLATFORM
ENV TARGETPLATFORM=${TARGETPLATFORM:-linux/amd64}

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN \
    export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) \
    && \
    export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) \
    && \
    GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1} \
    && \
    go mod download \
    && \
    go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o hajimari /build/main.go \
    && \
    chmod +x hajimari

FROM alpine:3.14

RUN \
    apk add --no-cache \
        tzdata \
        tini \
    && \
    addgroup -S hajimari \
    && \
    adduser -S hajimari -G hajimari

COPY --from=build /build/hajimari /usr/local/bin/hajimari

USER hajimari:hajimari
ENTRYPOINT [ "/sbin/tini", "--" ]
CMD [ "hajimari" ]

LABEL org.opencontainers.image.source https://github.com/toboshii/hajimari
