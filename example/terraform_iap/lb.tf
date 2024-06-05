module "lb-http" {
  source  = "terraform-google-modules/lb-http/google//modules/serverless_negs"
  version = "~> 10.0"

  name    = "lb"
  project = var.google_project_id

  ssl                             = true
  managed_ssl_certificate_domains = [var.hello_hostname]
  https_redirect                  = true

  backends = {
    default = {
      description = null
      groups = [
        {
          group = google_compute_region_network_endpoint_group.this.id
        }
      ]
      enable_cdn = false

      iap_config = {
        enable               = true
        oauth2_client_id     = google_iap_client.this.client_id
        oauth2_client_secret = google_iap_client.this.secret
      }
      log_config = {
        enable = true
      }
    }
  }

  depends_on = [
    google_project_service.this["compute.googleapis.com"],
  ]
}

resource "google_compute_region_network_endpoint_group" "this" {
  name                  = "hello-neg"
  network_endpoint_type = "SERVERLESS"
  region                = var.google_region
  cloud_run {
    service = module.hello-cloud-run.service_name
  }
}