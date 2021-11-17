FROM golang:latest
RUN mkdir /todoapp
COPY . /todoapp/
WORKDIR /todoapp/backend
ENV GOBIN /go/bin
RUN go get -u
RUN go build -o ../srv 
WORKDIR /todoapp/infrasructure
RUN ls
