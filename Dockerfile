# Image Source of Golang
FROM golang:1.19

# Create Work Directory in Docker
WORKDIR /app

# Copy go.mod and go.sum from apps to docker (literally all)
COPY . .

# Download the go.mod and go.sum that we copied
RUN go mod download
RUN go mod tidy

# copy the file generated after download
COPY *.go ./

# Build golang app
RUN go build -o /docker-omreca-service-api

# Expose port
EXPOSE 8081

# ENTRYPOINT ["/docker-omreca-service-api"]
# Run
CMD [ "/docker-omreca-service-api" ]
