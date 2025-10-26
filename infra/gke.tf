# GKE Standard 클러스터 (노드풀 분리 생성)
resource "google_container_cluster" "cluster" {
  name     = var.cluster_name
  location = var.region

  network    = google_compute_network.vpc.self_link
  subnetwork = google_compute_subnetwork.subnet.self_link

  # VPC-native
  ip_allocation_policy {
    cluster_secondary_range_name  = "pods-range"
    services_secondary_range_name = "services-range"
  }

  networking_mode = "VPC_NATIVE"
  release_channel {
    channel = "REGULAR"
  }

  # 기본 노드풀 비활성화 (별도 node pool 만들기 위해)
  remove_default_node_pool = true
  initial_node_count       = 1

  workload_identity_config {
    workload_pool = "${var.project_id}.svc.id.goog"
  }

  depends_on = [google_project_service.services]
}

resource "google_container_node_pool" "nodes" {
  name       = "${var.cluster_name}-pool"
  location   = var.region
  cluster    = google_container_cluster.cluster.name

  node_count = var.node_count

  node_config {
    machine_type = var.machine_type
    disk_size_gb = 20
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
    # 컨테이너 최적화 이미지
    image_type = "COS_CONTAINERD"
    labels = {
      role = "app"
    }
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }

  autoscaling {
    min_node_count = 2
    max_node_count = 4
  }
}
