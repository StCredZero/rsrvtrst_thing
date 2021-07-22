install-db:
	script/init_db.sh
	go run cmd/main.go -init

run:
	go run cmd/main.go

