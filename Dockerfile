FROM golang:1.24-alpine

WORKDIR /app

# Install Python and required build dependencies
RUN apk add --no-cache python3 py3-pip gcc musl-dev python3-dev libffi-dev openssl-dev

# Create and activate a virtual environment
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

# Copy requirements.txt and install Python dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the facts directory
COPY facts/ ./facts/

# Copy Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"] 