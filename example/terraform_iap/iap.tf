resource "google_iap_brand" "this" {
  project           = var.google_project_id
  application_title = "Example Application"
  support_email     = var.iap_support_email

  depends_on = [
    google_project_service.this["iap.googleapis.com"],
  ]
}

resource "google_iap_client" "this" {
  brand        = google_iap_brand.this.name
  display_name = "Example Client"
}

# IAPアクセス用SA
resource "google_service_account" "iap_access" {
  account_id = "iap-access"
}

# トークン作成権限を追加
resource "google_project_iam_member" "iap_access" {
  project = var.google_project_id
  member  = "serviceAccount:${google_service_account.iap_access.email}"
  role    = "roles/iam.serviceAccountOpenIdTokenCreator"
}

resource "google_iap_web_backend_service_iam_binding" "this" {
  project             = var.google_project_id
  web_backend_service = module.lb-http.backend_services["default"].name
  role                = "roles/iap.httpsResourceAccessor"
  members             = concat(["serviceAccount:${google_service_account.iap_access.email}"], var.iap_members)
}

resource "google_project_service_identity" "this" {
  provider = google-beta
  project  = var.google_project_id
  service  = "iap.googleapis.com"
}

resource "google_project_iam_member" "iap" {
  project = var.google_project_id
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_project_service_identity.this.email}"
}