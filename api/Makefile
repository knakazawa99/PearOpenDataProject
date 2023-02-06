setup:
	go install github.com/golang/mock/mockgen@latest
	go install github.com/sanposhiho/gomockhandler@latest

mock:
	gomockhandler -config=gomockhandler.json mockgen
mock-check:
	gomockhandler -config=gomockhandler.json check
