FROM golang:latest

WORKDIR /app

COPY . .

LABEL Project="ascii-art-web"

RUN go build -o ascii-art-web .

CMD [ "./ascii-art-web" ]