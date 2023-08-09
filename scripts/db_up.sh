DB_URL=$1
migrate -source file://migrations -database $DB_URL up
