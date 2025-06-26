BIN=app

run:build
	./$(BIN)

build:clean
	go build -o $(BIN) ./cmd

clean:
	rm -rf  $(BIN)


	
