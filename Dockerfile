FROM golang:1.14.2-alpine
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
RUN apk add --no-cache libstdc++ gcc g++ make git ca-certificates linux-headers wget curl jq libsass

WORKDIR /go/src/github.com/statping/statping
ADD . .

RUN go mod download && \
    go build -o emailer ./ && \
    mv emailer /usr/local/bin/

WORKDIR /app
VOLUME /app
EXPOSE 8000

ENV HOST ""
ENV USERNAME ""
ENV PASSWORD ""
ENV PORT 587

CMD emailer
