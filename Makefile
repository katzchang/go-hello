.PHONY: run curl

run: main.go
	go run $<

curl: index foo bar

index:
	curl -i "localhost:8081"

foo:
	curl -i "localhost:8081/foo"
	curl -i "localhost:8081/foo?a=hoge"

bar:
	curl -i "localhost:8081/bar?a=%7B%22n%22%3A%22Alice%22%2C%22b%22%3A%22hige-%22%7D"
	curl -i "localhost:8081/bar?a="
	curl -i "localhost:8081/bar"
