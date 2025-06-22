build: clean
	go build -o ./build/Executable ./app/*

test: build
	./build/Executable decode 10:HelloWorld
	./build/Executable decode i-055e 
	./build/Executable decode l5:helloi52ee
	./build/Executable decode d3:foo3:bar5:helloi52ee 
	./build/Executable info ./build/sample.torrent

clean:
	rm -f ./build/Executable
