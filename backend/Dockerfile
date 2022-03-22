# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM golang:1.17 AS stage-compile

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/trailheads
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/trailheads

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM alpine:latest 

WORKDIR /root/
COPY --from=stage-compile /go/src/app/trailheads ./

RUN apk --no-cache add ca-certificates

ENTRYPOINT [ "./trailheads" ]  