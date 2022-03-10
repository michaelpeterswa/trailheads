# # -=-=-=-=-=-=- Compile Image -=-=-=-=-=-=-

FROM node AS stage-ember-build

WORKDIR /app
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend/ .

RUN npm run build

# -=-=-=-=-=-=- Compile Go Binary -=-=-=-=-=-=-

FROM golang:1.17 AS stage-go-build

WORKDIR /go/src/app
COPY backend/ .

RUN go get -d -v ./cmd/trailheads
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/trailheads

# -=-=-=-=-=-=- Final Image -=-=-=-=-=-=-

FROM alpine:latest 

WORKDIR /root/

# Copy Go Binary
COPY --from=stage-go-build /go/src/app/trailheads .
# Copy Ember Dist
COPY --from=stage-ember-build /app/dist/ ./dist

RUN apk --no-cache add ca-certificates

ENTRYPOINT [ "./trailheads" ]  