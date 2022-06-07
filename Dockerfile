FROM golang:1.18.3-alpine

ENV PORT=8080

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /kma-bin

EXPOSE 8080

CMD [ "/kma-bin" ]