variable "google_project_id" {
  description = "The Google Cloud project ID"
  type        = string
}

variable "google_region" {
  description = "The Google Cloud region"
  type        = string
  default     = "us-west1"
}

variable "iap_support_email" {
  description = "The email address to use for IAP brand support"
  type        = string
}

variable "iap_members" {
  description = "The list of email addresses to allow access to the IAP-protected service"
  type        = list(string)
  default     = []
}

variable "hello_hostname" {
  description = "The hostname to use for the hello service"
  type        = string
}