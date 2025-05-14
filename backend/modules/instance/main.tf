resource "openstack_compute_instance_v2" "main" {
  name        = var.name
  image_id    = var.image_id
  flavor_name = var.flavor_name
  key_pair    = var.keypair_name

  network {
    port = var.port_id
  }
}
