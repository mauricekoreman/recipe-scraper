FROM golang:1.23

# Set the working directory inside the container for the build process
WORKDIR /app

# Copy the Go module files (go.mod, go.sum) first to leverage Docker's build cache.
# If these files don't change, subsequent builds will reuse this layer.
COPY go.mod go.sum ./

# Download the Go module dependencies.
RUN go mod download

# Copy the rest of the application code.
COPY . .

# Build the application.
RUN go build -v -o /usr/local/bin/recipe-scraper .

CMD ["recipe-scraper"]