ARG BASE=worldmaomao/hd-manager-base:1.0.0

FROM ${BASE} AS builder

WORKDIR $GOPATH/src/app

ARG MAKE='make build'

COPY . .

RUN $MAKE

FROM scratch

LABEL Name=web-hellword Version=${VERSION}

COPY --from=builder $GOPATH/src/app/cmd /web-helloworld

EXPOSE 48080

WORKDIR /web-helloworld/

CMD ["./web-helloworld"]