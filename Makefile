build:
	GOOS=linux GOARCH=amd64 go build  -o linux_amd64  .
	docker build --tag 'tempaicontext' .


run:
	docker run -p 8080:8080 'tempaicontext'
