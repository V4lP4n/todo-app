FROM golang:latest
RUN mkdir /todoapp
# COPY ../backend /todoapp
# WORKDIR /todoapp/backend
#RUN go get github.com/mattn/go-sqlite3
# RUN go mod init /todoapp/backend/main.go
# RUN go run /todoapp/backend/main.go
# CMD ["/todoapp/main"]

COPY . /todoapp/
WORKDIR /todoapp/backend
ENV GOBIN /go/bin
RUN go get -u
RUN go build -o ../srv 
# RUN apt install sqlite3
WORKDIR /todoapp/infrasructure
RUN ls
