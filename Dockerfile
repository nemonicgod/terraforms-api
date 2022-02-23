FROM golang:1.15-alpine

RUN mkdir -p /go/src/github.com/nemonicgod/terraforms-api
WORKDIR /go/src/github.com/nemonicgod/terraforms-api

RUN apk add --no-cache build-base

COPY go.mod .
COPY go.sum .

RUN go mod download && \
    go mod tidy && \
    go mod vendor

ADD . /go/src/github.com/nemonicgod/terraforms-api

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go build -o bin/api pkg/api/main.go
RUN go build -o bin/client pkg/client/main.go
RUN go build -o bin/worker pkg/worker/main.go

# Second layer, only for execution -----------------------------------
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root

# COPY --from=0 /usr/bin/node /usr/bin/node
# COPY --from=0 /usr/local/bin/perp /usr/local/bin/perp

# COPY --from=0 /go/src/github.com/nemonicgod/terraforms-api/bin/api .
# COPY --from=0 /go/src/github.com/nemonicgod/terraforms-api/bin/job .
# COPY --from=0 /go/src/github.com/nemonicgod/terraforms-api/seeds.json .

# ENTRYPOINT /root/api
