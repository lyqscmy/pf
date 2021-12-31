.PHONY: build
build:
	go build ./cmd/pf
install: build
	cp pf ~/.local/bin/
.PHONY: clean
clean:
	rm pf
