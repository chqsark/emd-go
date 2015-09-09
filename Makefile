all:
	gofmt -e -s -w .
	go vet .
	$(CC) -g -fPIC -c -o lib/lib.o lib/emd.c
	$(CC) -g -fPIC -shared -o liblib.so lib/lib.o
	LD_LIBRARY_PATH=lib/ go build demo_emd.go
