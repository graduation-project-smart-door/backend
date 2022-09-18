FROM golang:1.19.1

WORKDIR /opt/smart-door

COPY ./app /opt/smart-door

RUN go build -o /app/build/app ./cmd/app/main.go
CMD /app/build/app
