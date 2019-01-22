# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test-app .

FROM golang:1.10.1-alpine3.7 as builder
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test-app /app/*.go

FROM scratch
COPY --from=builder /go/test-app .
EXPOSE 8080
CMD ["./test-app"]
