resource "openstack_compute_keypair_v2" "main" {
  name       = var.name
  public_key = var.public_key
}
