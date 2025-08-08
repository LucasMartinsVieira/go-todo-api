@default:
  @just --list

create-migration:
  migrate create -ext=sql -dir=internal/database/migrations -seq init

migrate-up:
  migrate -path=internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
  migrate -path=internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

compose-down:
  docker compose -f development/compose.yml stop

compose-up:
  docker compose -f development/compose.yml up -d

sqlc-generate:
  cd internal/database && sqlc generate
