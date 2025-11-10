resource "google_service_account" "app_db_sa" {
  account_id   = "app-db-sa"
  display_name = "GKE App DB Access Service Account"
}

resource "google_project_iam_member" "cloudsql_client" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.app_db_sa.email}"
}

resource "google_project_iam_member" "cloudsql_instance_user" {
  project = var.project_id
  role    = "roles/cloudsql.instanceUser"
  member  = "serviceAccount:${google_service_account.app_db_sa.email}"
}

resource "google_service_account_iam_member" "gke_workload_identity_binding" {
  service_account_id = google_service_account.app_db_sa.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${var.project_id}.svc.id.goog[default/app-db]"
}
