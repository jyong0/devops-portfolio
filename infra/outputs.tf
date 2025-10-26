output "cluster_name" {
  value = google_container_cluster.cluster.name
}

output "region" {
  value = var.region
}

output "endpoint" {
  value = google_container_cluster.cluster.endpoint
}

output "workload_identity_pool" {
  value = google_container_cluster.cluster.workload_identity_config[0].workload_pool
}
