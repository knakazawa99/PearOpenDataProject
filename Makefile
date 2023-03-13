set-up:
	git config commit.template .github/.commit_template

up:
	docker compose up --build -d

down:
	docker compose down

create-gmail-token:
	cd gmail
	go run quickstart.go
	cp credentials.json ../api/
	cp token.json ../api/