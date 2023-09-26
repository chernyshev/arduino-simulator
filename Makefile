.DEFAULT_GOAL := disassemble
SRC := $(shell find . -name "*.go")

disassemble:
	go run *.go disasm --in sketches/bin/Blink.bin > out/blink-disasm.txt

# go install github.com/daixiang0/gci@latest
fmt: ## Format and fix import order
	@goimports -w -format-only $(SRC)
	@gci write --skip-generated -s standard -s default -s "prefix(cr46)" .
