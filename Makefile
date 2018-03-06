all:

clean:
	docker stop rqf-mongo
	docker stop rqf-test
	docker rm rqf-mongo
	docker rm rqf-test

test:
	docker build -t rqf-test .
	docker run --name rqf-mongo -d mongo
	docker run --name rqf-test --link rqf-mongo:mongo rqf-test
	@make clean