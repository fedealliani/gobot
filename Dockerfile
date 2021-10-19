FROM golang:latest

# Add a work directory
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /gobot

EXPOSE 8080

CMD [ "/gobot" ]
