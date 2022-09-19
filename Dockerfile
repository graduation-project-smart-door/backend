FROM golang:1.19.1

WORKDIR /opt/smart-door

ADD ./app /opt/smart-door/app
COPY ./migrations /opt/smart-door/
COPY ./go.mod /opt/smart-door/
COPY ./go.sum /opt/smart-door/
COPY ./.env /opt/smart-door/

RUN go mod tidy
RUN go build -o /app/build/app ./app/cmd/app/main.go
CMD /app/build/app
