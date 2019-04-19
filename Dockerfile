FROM golang:latest
RUN mkdir app/src/
WORKDIR app/src/
COPY . app/src/
RUN go build main.go
CMD ["/app/src/main"]

