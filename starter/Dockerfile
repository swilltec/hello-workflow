FROM golang:1.16-alpine3.13

# Set the working directory to /start
WORKDIR /app

COPY ./go.* ./

# RUN go mod download (For some reason this gets stucked here.)

COPY starter ./starter
COPY helloworkflow ./helloworkflow

WORKDIR /app/starter
RUN go build

CMD ["./starter"]