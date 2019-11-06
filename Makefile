all:
	-rm debug.log
	go build -o lime
	./lime
	rm lime
