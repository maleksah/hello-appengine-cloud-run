FROM golang:1.19.0 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

# Use the official Alpine image for a lean production container.
FROM alpine:3.17.2

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY index.html ./index.html
COPY assets/ ./assets/

# Run the web service on container startup.
CMD ["/server"]