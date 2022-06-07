FROM golang:1.18.3-alpine

LABEL maintainer="Dang Hoang Phuc <13364457+phuchptty@users.noreply.github.com>"

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