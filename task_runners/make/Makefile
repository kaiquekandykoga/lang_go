build:
	go build -o make make.go
	chmod +x make

exec: build
	./make

clean:
	rm -f make

all: build exec clean

run:
	go run make.go

.PHONY: run build exec all clean