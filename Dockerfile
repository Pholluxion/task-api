# Use base golang image from Docker Hub
FROM golang:latest AS build

WORKDIR /app

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . .

# Compile the application to /app.
RUN echo "Building the application..."
RUN go build -o server ./cmd/server

FROM gcr.io/distroless/static-debian11

# Set environment variables
ENV JWT_SECRET='$JWT_SECRET'
ENV PORT='$PORT'
ENV DB_CONNECTION='$DB_CONNECTION'
ENV TOKEN_EXPIRE_TIME='$TOKEN_EXPIRE_TIME'

# Copy template & assets
WORKDIR /app
COPY --from=build app/server .
# Expose port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["./server"]
