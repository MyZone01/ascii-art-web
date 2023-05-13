run:
	docker build -t ascii-art-web .
	docker run --name ascii-container -p 8080:8080 -dt ascii-art-web
	@echo ""
	@echo "GO to: http://127.0.0.1:8080"

stop:
	docker stop ascii-container

remove:
	@docker rm $(docker ps -a -q)