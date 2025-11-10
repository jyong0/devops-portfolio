# --- Cloud SQL PostgreSQL
resource "google_sql_database_instance" "postgres" {
  name             = "devops-postgres"
  database_version = "POSTGRES_14"
  region           = var.region

  settings {
    tier = "db-f1-micro"
    backup_configuration {
      enabled = true
    }
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.vpc.self_link
    }
  }
}

resource "google_sql_user" "default" {
  name     = "app"
  instance = google_sql_database_instance.postgres.name
  password = var.db_password
}

resource "google_sql_database" "appdb" {
  name     = "appdb"
  instance = google_sql_database_instance.postgres.name
}
