# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test-app .

FROM scratch

COPY . /

EXPOSE 8080

CMD ["./test-app"]
