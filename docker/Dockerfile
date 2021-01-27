from golang:latest

COPY web_server.go /
RUN go build -o /web_server /web_server.go
EXPOSE 8000
ENTRYPOINT ["/web_server"]
