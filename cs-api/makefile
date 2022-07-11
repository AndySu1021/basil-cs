server: sqlc
	go run main.go server

mock:
	mockgen -source pkg/interface/service.go -destination dist/mock/service_mock.go -package mock
	mockgen -source pkg/interface/repo.go -destination dist/mock/repo_mock.go -package mock
	mockgen -source pkg/interface/script.go -destination dist/mock/script_mock.go -package mock

sqlc:
	sqlc compile && sqlc generate

mc:
	migrate create -ext sql -seq -digits 3 -dir ./db/migrations $(n)

migrate.rollback:
	migrate -path ./db/migrations -database "mysql://root:abcd1234@tcp(localhost:3306)/cs-api?charset=utf8mb4&multiStatements=true&parseTime=true" -verbose down $(n)

test:
	go test ./pkg/module/...