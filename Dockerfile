FROM docker.io/node:22.2-alpine AS build-frontend

WORKDIR /build

COPY . .

WORKDIR /build/frontend

RUN npm install

RUN npm run build

FROM docker.io/golang:1.22.3-alpine as build

ARG TARGETPLATFORM
ENV TARGETPLATFORM=${TARGETPLATFORM:-linux/amd64}

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

COPY --from=build-frontend /build/frontend/build /build/frontend/build

RUN \
    export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) \
    && \
    export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) \
    && \
    GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1} \
    && \
    go mod download \
    && \
    go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o hajimari /build/cmd/hajimari/main.go \
    && \
    chmod +x hajimari

FROM docker.io/alpine:3.19

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

LABEL org.opencontainers.image.source https://github.com/ullbergm/hajimari
