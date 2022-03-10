# -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM golang:1.16 AS stage-compile

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./cmd/wildfire
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/wildfire

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM alpine:latest 

WORKDIR /root/
COPY --from=stage-compile /go/src/app/wildfire ./

RUN apk --no-cache add ca-certificates

ENTRYPOINT [ "./wildfire" ]  