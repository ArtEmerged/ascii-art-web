build:
	docker build -t ascii-art-web .
run:
	docker run -d --name ascii-art-web -p 8080:8080 ascii-art-web 
	echo "Listening on: http://localhost:8080/"