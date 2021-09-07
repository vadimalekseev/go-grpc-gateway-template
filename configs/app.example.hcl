database {
  host = "127.0.0.1"
  port = 5432
  user = "postgres"
  password = "postgres"
  database = "postgres"
  sslmode = "disable"
}

app {
  domain = "example.com"

  httpAddr = ":5555"
  grpcAddr = ":5556"
}
