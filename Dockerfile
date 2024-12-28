FROM alpine:latest

COPY CaptainFeedHook "/"
RUN apk add libc6-compat && mkdir /config && echo [] > /config/save.json

CMD ["/CaptainFeedHook"]
