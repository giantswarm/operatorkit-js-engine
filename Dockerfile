FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./operatorkit-js-engine /operatorkit-js-engine

ENTRYPOINT ["/operatorkit-js-engine"]
