run: bin/rest-gotrello
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/rest-gotrello: main.go
	go build -o bin/rest-gotrello main.go
clean:
	rm -rf bin
