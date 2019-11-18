# Go Example Webapp

This is a simple web-app that prints out the current time, system hostname,
build time, and build commit hash. This is mainly used for testing pipeline
builds for an Infrastrucutre Operations class CI/CD Assignment.

## Building and Running

For local development this requires you have Golang installed (last tested with
1.12.6) and can be easily built and tested using the following commands.

```bash
go build -ldflags "-X main.BuildTime=$(date +"%m-%d-%y-%H%M%S") -X main.CommitSHA=$(git rev-list -1 HEAD)"
./go-example-webapp
```

For use with docker, it can be built and run using

```bash
docker build . -t go-example-webapp:latest
docker run -d -p8080:8080 go-example-webapp
```