FROM golang:1.13-alpine3.11


WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application

RUN go build  -o consumer .

WORKDIR /distt
RUN cp /app/consumer  .
CMD ["/distt/consumer"]