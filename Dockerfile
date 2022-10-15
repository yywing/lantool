FROM alpine:latest

COPY release/lantool_linux_amd64 /usr/bin/lantool
ENTRYPOINT ["lantool"]