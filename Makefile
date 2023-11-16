migrate:
	goose -dir tools/migrations postgres "host=localhost user=admin dbname=perfumeslctdb sslmode=disable password=pass" up
	
migrate-down:
	goose -dir tools/migrations postgres "host=localhost user=admin dbname=perfumeslctdb sslmode=disable password=pass" down

run:
	go run cmd/perf-bot/main.go