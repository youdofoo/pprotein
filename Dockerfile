# --------------------------------------------------
ARG GO_VERSION=1.23.3
FROM golang:${GO_VERSION} AS pprotein

RUN apt update && apt install -y npm make

WORKDIR $GOPATH/src/app
COPY . .

RUN make build

# --------------------------------------------------

FROM golang:${GO_VERSION} AS alp

RUN go install github.com/tkuchiki/alp/cmd/alp@latest

# --------------------------------------------------

FROM golang:${GO_VERSION} AS slp

RUN go install github.com/tkuchiki/slp/cmd/slp@latest

# --------------------------------------------------

FROM debian:bookworm-slim

RUN apt update && apt install -y graphviz percona-toolkit

COPY --from=pprotein /go/src/app/pprotein /usr/local/bin/
COPY --from=pprotein /go/src/app/pprotein-agent /usr/local/bin/
COPY --from=alp /go/bin/alp /usr/local/bin/
COPY --from=slp /go/bin/slp /usr/local/bin/

RUN mkdir -p /opt/pprotein
WORKDIR /opt/pprotein

ENTRYPOINT ["pprotein"]
