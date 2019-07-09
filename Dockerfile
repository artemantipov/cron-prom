FROM golang:1.12-alpine as builder
RUN apk add --no-cache git
RUN adduser -D -u 10001 appuser
ADD . /go
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cron-prom .

FROM scratch
COPY --from=builder /go/cron-prom /cron-prom
COPY --from=builder /etc/passwd /etc/passwd
USER appuser
CMD ["/cron-prom"]