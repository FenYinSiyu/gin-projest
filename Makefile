build-linux-amd64:
	docker build --platform=linux/amd64 -t yinsiyu/go-project .

launch-app:
	docker-compose up -d

image-push:
	docker push yinsiyu/go-project