variable "project_id" {
  description = "GCP Project ID"
  type        = string
  default     = "storied-reserve-473208-t1"
}

variable "region" {
  description = "Region for resources (e.g., asia-northeast3)"
  type        = string
  default     = "asia-northeast3"
}

variable "network_name" {
  description = "VPC network name"
  type        = string
  default     = "devops-portfolio-net"
}

variable "subnet_name" {
  description = "Subnet name"
  type        = string
  default     = "devops-portfolio-subnet"
}

variable "subnet_cidr" {
  description = "Subnet CIDR"
  type        = string
  default     = "10.10.0.0/20"
}

variable "pods_secondary_cidr" {
  description = "Secondary range for Pods (VPC-native)"
  type        = string
  default     = "10.20.0.0/16"
}

variable "services_secondary_cidr" {
  description = "Secondary range for Services (VPC-native)"
  type        = string
  default     = "10.30.0.0/20"
}

variable "cluster_name" {
  description = "GKE cluster name"
  type        = string
  default     = "devops-gke"
}

variable "node_count" {
  description = "Initial node count"
  type        = number
  default     = 2
}

variable "machine_type" {
  description = "Node machine type"
  type        = string
  default     = "e2-medium"
}

variable "db_password" {
  type        = string
  description = "Password for PostgreSQL user"
  sensitive   = true
}
