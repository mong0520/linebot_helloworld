build:
	@docker-compose build

up:
	@docker-compose up -d

down:
	@docker-compose down

restart:
	@docker-compose restart

tunnel:
	@ngrok http 2000

debug_%:
	@docker exec -it $* /bin/bash