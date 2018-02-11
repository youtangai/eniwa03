eniwa: main.go
	GOOS=linux GOARCH=amd64 go build -o eniwa main.go

clean:eniwa
	rm -f eniwa

build:eniwa
	docker build -t yotanagai/eniwa03 .

push:
	docker push yotanagai/eniwa03

