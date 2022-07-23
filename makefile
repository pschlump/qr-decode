
all:
	go build


test: test1 test2 test3
	@echo ""
	@echo PASS

test1:
	go build
	@mkdir -p out ref
	./qr-decode testdata/25b9045a71.png >out/test1.out
	@diff out/test1.out ref

test2:
	go build
	@mkdir -p out ref
	./qr-decode --raw testdata/25b9045a71.png >out/test2.out
	@diff out/test2.out ref

test3:
	go build
	@mkdir -p out ref
	./qr-decode --output out/test3.out testdata/25b9045a71.png
	@diff out/test3.out ref
