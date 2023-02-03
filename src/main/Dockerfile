# ---------------------------------------------------------
# Build
# ---------------------------------------------------------
FROM golang:1.20-rc-alpine AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY go /app
WORKDIR /app

RUN go mod download \
    && go test ./... \
    && go build -o target/jiracli .

# ---------------------------------------------------------
# Run
# ---------------------------------------------------------
FROM alpine:3.17.1 AS run
LABEL maintainer="sebastian@sommerfeld.io"

# To run integration tests with a proper shell
RUN apk add --no-cache bash=5.2.15-r0

COPY --from=build /app/target/jiracli /usr/bin/jiracli

ARG USER=jiracli
RUN adduser -D $USER

USER $USER

ENTRYPOINT [ "/usr/bin/jiracli" ]
