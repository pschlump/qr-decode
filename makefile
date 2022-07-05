
all:
	go build


test:
	go build
	@mkdir -p out ref
	./qr-decode testdata/25b9045a71.png >out/test1.out
	@diff out/test1.out ref
	@echo PASS

