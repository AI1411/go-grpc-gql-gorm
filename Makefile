gen_article:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/article.proto

psql:
	docker compose exec db psql -h localhost -p 5432 -U root -d grpc_gql