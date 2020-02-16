FROM cassandravsscylla_base:latest as builder

Add . /etc/test/
WORKDIR /etc/test
#RUN tail -f /dev/null
# Build Service as Executable:

RUN GOOS=linux GOARCH=amd64 go build -x -o main -ldflags="-w -s"

############################
#STEP 2 build a small image
############################
FROM alpine

ENV PKG_PATH /etc/relay

WORKDIR /

RUN apk update

# Import CA certs and TZ-Data from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/timezone /etc/timezone
WORKDIR $PKG_PATH

#Copying Go App Binary
COPY --from=builder $PKG_PATH/main $PKG_PATH/


#Starting Go Application
CMD $PKG_PATH/main
