FROM golang AS builder
COPY alphatree .
RUN export GOPATH=/go/src/alphatree
RUN go get -v ./src/alphatree
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/alphatree

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir ./static
COPY --from=builder /go/main .
COPY --from=builder /go/static ./static
RUN \
  mkdir ./logs && \
  chmod -R +rwx ./logs && \
  echo "Creating log file" > ./logs/main_log.txt && \
  chmod +rwx ./logs/main_log.txt
CMD ./main
