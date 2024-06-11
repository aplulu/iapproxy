terraform {
  required_version = "~> 1.7.0"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }

    random = {
      source = "hashicorp/random"
    }
  }
}

provider "google" {
  project = var.google_project_id
  region  = var.google_region

  default_labels = {
    "terraform" = "true"
  }
}