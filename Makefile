set-up:
	git config commit.template .github/.commit_template

up:
	docker compose up --build -d

down:
	docker compose down