module "app" {
  source = "./deployments/application"

  external_network_name = "ntnu-internal"
  keypair_name          = "app-key"
  public_key            = var.public_key
}

module "frontend" {
  source = "./deployments/frontend"

  external_network_name = "ntnu-global"
  keypair_name          = "frontend-key"
  public_key            = var.public_key
}
