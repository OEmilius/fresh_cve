FROM golang:latest
RUN mkdir -p /go/src/fresh_cve
WORKDIR /go/src/fresh_cve
COPY . /go/src/fresh_cve
CMD ./fresh_cve
EXPOSE 8081
