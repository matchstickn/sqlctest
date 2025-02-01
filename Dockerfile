# # syntax=docker/dockerfile:1

# FROM alpine:3.20

# # Set destination for COPY
# WORKDIR /app

# RUN apk add --no-cache \
#     ca-certificates \
#     git \
#     && wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz

# RUN tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz 

# RUN rm go1.23.3.linux-amd64.tar.gz

# ENV PATH="/usr/local/go/bin:${PATH}"
# ENV GOPATH="/go"
# ENV GOBIN="/go/bin"

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/reference/dockerfile/#copy
# RUN mkdir cmd
# COPY ./cmd/main.go ./cmd

# RUN mkdir -p assets/db
# RUN cd assets/db

# COPY ./assets/db *.go ./

# RUN cd ..
# RUN cd ..

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o app/docker-sqlc

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/reference/dockerfile/#expose
# EXPOSE 4000

# RUN ls -l

# # Run
# RUN chmod +x app/docker-sqlc
# ENTRYPOINT ["app/docker-sqlc"]
# CMD ["app/docker-sqlc"]




# syntax=docker/dockerfile:1

# FROM debian:bullseye-slim

# # Set destination for COPY
# WORKDIR /app

# # Install necessary packages
# RUN apt-get update && apt-get install -y \
#     ca-certificates \
#     git \
#     wget \
#     && rm -rf /var/lib/apt/lists/*

# # Download Go
# RUN wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz

# RUN tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz 

# RUN rm go1.23.3.linux-amd64.tar.gz

# ENV PATH="/usr/local/go/bin:${PATH}"
# ENV GOPATH="/go"
# ENV GOBIN="/go/bin"

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code
# RUN mkdir cmd
# COPY ./cmd/main.go ./cmd

# RUN mkdir -p assets/db
# RUN cd assets/db

# COPY ./assets/db *.go ./

# RUN cd ..
# RUN cd ..

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o app/docker-sqlc

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# EXPOSE 4000

# RUN ls -l

# # Run
# RUN chmod +x app/docker-sqlc
# CMD ["app/docker-sqlc"]


# syntax=docker/dockerfile:1

# Use build arguments for platform specification
ARG TARGETPLATFORM="linux/amd64"
FROM alpine:3.20

# Set up environment
WORKDIR /app

# Install dependencies (keep arch-specific components explicit)
RUN apk add --no-cache \
    ca-certificates \
    git \
    && wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz \
    && rm go1.23.3.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GOBIN="/go/bin"

# Copy Go modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source files with directory structure
COPY ./cmd ./cmd
COPY ./assets/db ./assets/db

# Explicitly build for AMD64 regardless of host architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /app/docker-sqlc ./cmd

# Verify assets/db copy
RUN ls -l ./assets/db

# Set executable permissions
RUN chmod +x /app/docker-sqlc

EXPOSE 4000

ENTRYPOINT ["/app/docker-sqlc"]