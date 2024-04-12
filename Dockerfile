# use official Golang image
FROM golang:1.22-alpine3.18

# set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go mod download

# Build the Go app
RUN go build -o todo .

#EXPOSE the port
EXPOSE 8080

# Run the executable
CMD ["./todo"]