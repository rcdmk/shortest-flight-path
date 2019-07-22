#################################
#           BUILD API           #
#################################
FROM golang:1.12.7-alpine3.10

ARG WORKDIR=/go/src/github.com/rcdmk/shortest-flight-path

RUN apk add --no-cache \
    ca-certificates \
    gcc \
    g++

RUN mkdir -p ${WORKDIR}
ADD . ${WORKDIR}
WORKDIR ${WORKDIR}

RUN go build -tags netgo -installsuffix netgo -o api

#################################
#            RUN API            #
#################################
FROM alpine:3.10.1

ARG BUILD_PATH=/go/src/github.com/rcdmk/shortest-flight-path
ARG WORKDIR=/opt/app

RUN apk update && apk add ca-certificates

RUN mkdir -p ${WORKDIR}

COPY --from=0 ${BUILD_PATH}/api ${WORKDIR}/api
COPY --from=0 ${BUILD_PATH}/config.toml ${WORKDIR}/config.toml

WORKDIR ${WORKDIR}
EXPOSE 5000
CMD ["./api"]