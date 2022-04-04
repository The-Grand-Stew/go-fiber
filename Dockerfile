#TODO: add other base images especially alpine linux for aws, centos blah blah
FROM golang:1.18-buster as builder
COPY . /app
WORKDIR /app
# build the binary
RUN GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o appBinary cmd/main.go
# copy the binary file the root directory
RUN cp appBinary /appBinary
# remove source code folder
RUN rm -rf /app


FROM alpine:3.14 as deploy
RUN apk add libc6-compat
COPY --from=builder /appBinary .
CMD ["./appBinary"]