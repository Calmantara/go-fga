# scenario 3: multistage docker
# STAGE1: BUILD BINARY
FROM golang:latest as builder

# membuat work dir untuk builder
RUN mkdir -p /app
WORKDIR /app

# copy all assets
COPY . .

# build binary
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARC=amd64 go build  -o go-fga .

# STAGE2: FINAL IMAGE
# final image
FROM alpine:latest

# membuat work dir untuk final image
RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /app/go-fga .
COPY --from=builder /app/.env .

CMD [ "./go-fga" ]