FROM golang:1.17
WORKDIR /app
COPY ./ /app
RUN go mod download
EXPOSE 8080
CMD ["/server"]
#RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build ./cmd/anyData/main.go" --command=./runserver
#RUN cd ./cmd/anyData && go run main.go