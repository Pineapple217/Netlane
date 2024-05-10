codegen:
	templ generate
	go generate ./pkg/ent

build:
	@make --no-print-directory codegen
	go build -o ./tmp/main.exe ./cmd/main

start:
	@./tmp/main.exe

run: 
	@make --no-print-directory build
	@make --no-print-directory start