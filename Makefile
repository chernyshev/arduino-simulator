.DEFAULT_GOAL := disassemble

disassemble:
	go run *.go disasm --in sketches/bin/Blink.bin > out/blink-disasm.txt
