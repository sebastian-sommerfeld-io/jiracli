FROM alpine:3.17.2 AS run
LABEL maintainer="sebastian@sommerfeld.io"

# To run reindex commands with a proper shell
RUN apk add --no-cache bash=5.2.15-r0 \
    && apk add --no-cache curl=7.87.0-r2

ARG USER=indexer
RUN adduser -D $USER

USER $USER
