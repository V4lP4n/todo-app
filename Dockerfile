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
# RUN set -x go get get github.com/mattn/go-sqlite3 dep ensure -v
RUN go get -u github.com/mattn/go-sqlite3
RUN go install -v ./...
RUN go build -o ../srv main.go 
RUN infrasructure/make_db.sh
CMD ["srv"]