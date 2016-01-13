.PHONY: run curl

run: main.go
	go run $<

curl:
	curl -i "localhost:8081"
	curl -i "localhost:8081/foo"
	curl -i "localhost:8081/foo?a=hoge"
