FROM alpine:3.6

COPY go-hello-world /usr/local/bin/app
EXPOSE 3000
ENTRYPOINT ["app"]
