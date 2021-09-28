bufupdate:
	buf mod update
bufgenerate:
	buf generate
migrateup:
	migrate -path db/migration -database "postgresql://surendrakandel:postgres@localhost:5432/market?sslmode=disable" -verbose up	
migratedown:
	migrate -path db/migration -database "postgresql://surendrakandel:postgres@localhost:5432/market?sslmode=disable" -verbose down