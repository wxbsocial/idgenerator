export $(grep -v '^#' .env.test | xargs)
go test -v -cover 