
.PHONY: clean example 

test:
	cd openligadb && go test -v -race ./... 



example:
	cd example/simple && go build && ./simple

all:
	cd openligadb && go build

clean:
	cd example/simple && rm simple 
