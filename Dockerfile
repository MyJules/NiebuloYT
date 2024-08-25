# syntax=docker/dockerfile:1

##
## Build the application from source
##

FROM golang:1.19 AS build-stage

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/niebulo/*.go ./niebulo/
COPY src/yt/*.go ./yt/
COPY src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /NiebuloYT

##
## Run the tests in the container
##

FROM build-stage AS run-test-stage
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /home/nonroot

COPY --from=build-stage /NiebuloYT ./NiebuloYT
COPY config/config.yaml ./config.yaml
COPY bin/yt-dlp-linux ./yt-dlp.exe

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./NiebuloYT"]
