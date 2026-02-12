FROM golang:1.25-alpine AS build

WORKDIR /src
COPY script.go ./

RUN go build -o /app ./script.go

FROM alpine:3.20

WORKDIR /app
COPY --from=build /app /app/app

ENTRYPOINT ["/app/app"]
