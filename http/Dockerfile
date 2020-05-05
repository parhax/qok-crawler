FROM golang:1.13-alpine3.11


WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .


RUN go build  -o main .

WORKDIR /dist
RUN cp /app/main  .
CMD ["/dist/main"]