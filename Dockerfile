FROM golang:1.24.5-alpine

COPY . .

ENTRYPOINT ["go", "run", "main.go", "-rtp=1.0"]