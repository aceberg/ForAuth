FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/ForAuth/ && CGO_ENABLED=0 go build -o /ForAuth .


FROM scratch

WORKDIR /data/ForAuth
WORKDIR /app

COPY --from=builder /ForAuth /app/

ENTRYPOINT ["./ForAuth"]