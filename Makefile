include ./backend/.env
export

all: docker-run

docker-run:
	docker-compose --env-file ./backend/.env up --build
.PHONY: docker-run

clean:
	docker system prune