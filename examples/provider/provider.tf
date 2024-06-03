terraform {
  required_providers {
    sendgrid = {
      version = "1.0.0"
      source  = "registry.terraform.io/bid-ops-development/sendgrid"
    }
  }
}

provider "sendgrid" {
}
