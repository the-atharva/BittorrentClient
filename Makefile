build: clean
	go build -o ./build/Executable ./app/*

clean:
	rm -f ./build/Executable
