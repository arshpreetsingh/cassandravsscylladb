FROM cassandravsscylladb_base:latest as builder
Add . /etc/test/
WORKDIR /etc/test
#RUN tail -f /dev/null
# Build Service as Executable:
RUN go build -o testService -ldflags="-w -s"
#RUN tail -f /dev/null
ENTRYPOINT ["./testService"]
