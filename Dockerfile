FROM golang:latest AS build-stage

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -ldflags "-linkmode external -extldflags -static"
RUN echo [] > save.json

FROM scratch
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage /app/CaptainFeedHook "/"
COPY --from=build-stage /app/save.json "/config/"
CMD ["/CaptainFeedHook"]