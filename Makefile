# run days
run:
	go run main.go

test:
	go test ./cmd/...

bench:
	go test -bench=. ./...
