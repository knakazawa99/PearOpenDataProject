set-up:
	git config commit.template .github/.commit_template

up:
	docker compose up --build -d

down:
	docker compose down

create-gmail-token:
	cd gmail && go run quickstart.go
	cp gmail/credentials.json api/
	cp gmail/token.json api/