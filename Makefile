.PHONY: clean

build:
	docker build -t sqlx/test:0.0.1 --no-cache .
test:
	docker run -d --name sqlxtest -p 3306:3306 sqlx/test:0.0.1
	docker logs -f sqlxtest

clean:
	- docker rm -f sqlxtest
	- docker rmi sqlx/test:0.0.1