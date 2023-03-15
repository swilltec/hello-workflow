FROM golang:1.16-alpine3.13

WORKDIR /app

COPY ./go.* ./

# RUN go mod download (For some reason this gets stucked here.)

COPY worker ./worker
COPY helloworkflow ./helloworkflow

WORKDIR /app/worker
RUN go build

CMD ["./worker"]

