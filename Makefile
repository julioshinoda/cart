.EXPORT_ALL_VARIABLES:


run: 
	docker-compose up 

stop:
	docker-compose stop

test:
	docker build -t cart_test -f Dockerfile.test .
	docker run cart_test
