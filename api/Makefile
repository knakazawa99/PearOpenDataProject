setup:
	go install github.com/golang/mock/mockgen@latest
	go install github.com/sanposhiho/gomockhandler@latest

mock:
	gomockhandler -config=gomockhandler.json mockgen
mock-check:
	gomockhandler -config=gomockhandler.json check

set-up:
	brew install ariga/tap/atlas

migrate-hash:
	atlas migrate hash --dir file://./db/migrations

migrate:
	make migrate-hash
	atlas migrate apply --url mysql://root:root@:3306/pear_system --dir file://./db/migrations

schema-inspect:
	atlas schema inspect -u mysql://root:root@:3306/pear_system

