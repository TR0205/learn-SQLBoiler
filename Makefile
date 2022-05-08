up:
	docker compose up -d
build1:
	docker compose build --no-cache --force-rm
ps:
	docker compose ps
app:
	docker compose exec app bash
db1:
	docker compose exec db bash
log-app:
	docker compose logs app