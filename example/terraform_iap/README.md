# Infrastructure for Testing IAP Proxy

This directory contains Terraform code to set up the infrastructure required for testing the IAP (Identity-Aware Proxy) Proxy, also known as `iapproxy`. 

## Prerequisites

- Terraform v1.7.0 or later
- Google Cloud SDK
- A Google Cloud Platform (GCP) project with billing enabled

## Setup

1. Authenticate your GCP account:

```bash
gcloud auth login
gcloud auth application-default login
```

2. Navigate to the `terraform_iap`:

```bash
cd example/terraform_iap
```

3. Copy the `terraform.tfvars.example` file to `terraform.tfvars` and update the variables:

```bash
cp terraform.tfvars.example terraform.tfvars
vim terraform.tfvars
```

4. Initialize Terraform:

```shell
terraform init
```

5. Plan the infrastructure changes:

```shell
terraform plan
```

6. Apply the infrastructure changes:

```shell
terraform apply
```

## Usage

Once the infrastructure is set up, you can start testing the `iapproxy`.

## Cleanup

To destroy the created infrastructure:

```bash
terraform destroy
```
