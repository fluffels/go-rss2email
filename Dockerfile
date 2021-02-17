FROM golang:1.16.0-alpine3.13 AS builder

WORKDIR /home
COPY . .
RUN go build -o re


FROM alpine:3.13
COPY --from=builder /home/re /usr/bin
CMD [ "re" ]
