    FROM golang:alpine
    ENV GO111MODULE=on
    ENV envport = 9002
    ENV envtitle="Shamim"
    ENV envmsg="will show here" 
    WORKDIR /home/shamimhello
    COPY . .
    RUN go build
    RUN go install -v ./...
    EXPOSE 9001
    CMD ["./web"]