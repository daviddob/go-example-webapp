FROM golang:alpine AS build-env
RUN apk add --no-cache git mercurial
COPY . /go/src/go-example-webapp
RUN cd /go/src/go-example-webapp && go build -ldflags "-X main.BuildTime=$(date +"%m-%d-%y-%H%M%S") -X main.CommitSHA=$(git rev-list -1 HEAD)"

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/go-example-webapp/go-example-webapp /app
COPY --from=build-env /go/src/go-example-webapp/templates /app/templates
COPY --from=build-env /go/src/go-example-webapp/static /app/static

EXPOSE 8080
ENTRYPOINT [ "./go-example-webapp" ]