database {
  host     = "127.0.0.1"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  database = "postgres"
  sslmode  = "disable"
}

app {
  use_grpc_reflect = true
  grpc_addr        = "127.0.0.1:5556"
  http_addr        = "127.0.0.1:5555"
  metrics_addr     = "/metrics"
}
