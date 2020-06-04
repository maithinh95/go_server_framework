# Build: $ docker build -t imagename:tag -f Dockerfile.tagfile .
# Run: $ docker run -it --rm imagename:tag
# OR   $ docker run -d -v share_dir:SHARE_DIR -v log_dir:LOG_DIR -e SHARE_DIR=share_dir -e LOG_DIR=log_dir --name dockername --restart on-failure imagename:tag 
# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.14 as builder

# Add Maintainer Info
LABEL maintainer="WeeDigital Company | admin@wee.vn"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/facepayserver

######## Start a new stage from scratch #######
FROM alpine:latest 

RUN apk --no-cache add ca-certificates
RUN apk update;\
    apk add tzdata

ARG TZ=Asia/Ho_Chi_Minh
RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime;\
    date

WORKDIR /root/

# Build Args
ARG SHARE_DIR=/etc/facepayserver
ARG LOG_DIR=/var/log/facepayserver

# Create Log Directory
RUN mkdir -p ${SHARE_DIR}
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV SHARE_DIR_LOCATION=${SHARE_DIR}
ENV LOG_DIR_LOCATION=${LOG_DIR}
ENV CONFIG_LOCATION=${SHARE_DIR}/config.d/setting.conf

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin ./bin
# Copy the config file from the previous stage
COPY --from=builder /app/config.d ${SHARE_DIR}/config.d


# RUN ls -al bin
# RUN ls -al /etc/facepayserver/config.d
# RUN ls -al /var/log/facepayserver

# This container exposes port 8080 to the outside world
EXPOSE 8000

# Declare volumes to mount
VOLUME [${SHARE_DIR}, ${LOG_DIR}}]

# Run the binary program produced by `go install`
CMD ["./bin/facepayserver"]