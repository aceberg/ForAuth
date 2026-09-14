FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx

FROM --platform=$BUILDPLATFORM golang:alpine AS builder

COPY --from=xx / /

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETPLATFORM
RUN CGO_ENABLED=0 xx-go build -trimpath -ldflags='-w -s' -o /ForAuth ./cmd/ForAuth


FROM alpine

WORKDIR /data/ForAuth
WORKDIR /app

COPY --from=builder /ForAuth /app/

ENTRYPOINT ["./ForAuth"]