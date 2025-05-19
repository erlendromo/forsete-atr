# Credentials

variable "application_credential_id" {
  description = "OpenStack Application Credential ID"
  nullable    = false
  type        = string
}

variable "application_credential_secret" {
  description = "OpenStack Application Credential Secret"
  nullable    = false
  type        = string
  sensitive   = true
}

# Public key to OpenStack

variable "public_key" {
  description = "Public key for OpenStack keypair"
  nullable    = false
  type        = string
  sensitive   = true
}
