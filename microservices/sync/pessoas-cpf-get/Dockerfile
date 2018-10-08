FROM golang:1.11.1-stretch
LABEL maintainer="@douglaspands"
LABEL version="1.0.0"
COPY . /var/www
WORKDIR /var/www
RUN go get -d -v github.com/gorilla/mux
RUN go install -v github.com/gorilla/mux
RUN go build main.go
CMD ./main