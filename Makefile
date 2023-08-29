build:
	docker build -t ascii-art-web .
run:
	docker run --name ascii-art-web -p 8080:8080 ascii-art-web 