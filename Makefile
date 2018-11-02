run: bin/heroku-go-example
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/heroku-go-example: main.go
	go get -u github.com/gorilla/mux || go build -o bin/heroku-go-example main.go

clean:
	rm -rf bin
