FROM golang:1.20.1

WORKDIR /app

COPY . .

LABEL Project="ascii-art-web"

RUN go build -o ascii-art-web .

CMD [ "./ascii-art-web" ]