provider "openstack" {
  application_credential_id     = var.application_credential_id
  application_credential_secret = var.application_credential_secret
  user_domain_name              = "NTNU"
  auth_url                      = "https://api.skyhigh.iik.ntnu.no:5000/v3/"
  region                        = "SkyHiGh"
}
