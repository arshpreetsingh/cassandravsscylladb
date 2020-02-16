FROM golang:1.13.5 as builder
ENV PKG_PATH /etc/relay
RUN unset GOPATH

#Add Dependencies to base image
ADD shared/ $PKG_PATH/shared/

# Add Maintainer Info
LABEL maintainer="Khagura https://medium.com/@arshpreetsingh"

# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apt-get update && apt-get install ca-certificates tzdata tar curl python && update-ca-certificates

#Adding all timezones and Setting UTC as Default
RUN echo "UTC" >  /etc/timezone


# Add all required modules into single image.
WORKDIR $PKG_PATH
ADD go.mod .
ADD go.sum .
ENV GO111MODULE=auto
RUN go mod download
RUN go mod verify
