FROM golang:1.20-rc-alpine AS build
LABEL maintainer="sebastian@sommerfeld.io"

COPY go /app
WORKDIR /app

RUN go mod download \
    && go test ./... \
    && go build -o target/jiracli .


FROM alpine:3.17.1 AS run
LABEL maintainer="sebastian@sommerfeld.io"

COPY --from=build /app/target/jiracli /bin/jiracli

ARG USER=jiracli_user
RUN adduser -D $USER

USER $USER

ENTRYPOINT [ "/bin/jiracli" ]
