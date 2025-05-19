# Instance Module

This module provisions a virtual machine (VM) in OpenStack using the `openstack_compute_instance_v2` resource.
It allows you to specify a network port and associate an SSH key pair with the instance.

---

## Variables

| Name           | Description                              | Type   | Required |
|----------------|------------------------------------------|--------|----------|
| `name`         | Name of the compute instance             | string | Yes      |
| `image_id`     | ID of the image to use for the instance  | string | Yes      |
| `flavor_name`  | OpenStack flavor (e.g. `sx3.12c32r`)     | string | Yes      |
| `keypair_name` | Name of the OpenStack keypair for SSH    | string | Yes      |
| `port_id`      | ID of the networking port to attach to   | string | Yes      |

---

## Outputs

This module does NOT define any outputs by default.
You can extend it to output values like:

- `instance_id`
- `access_ip_v4`
- `name`

as needed for your use case.

---

## Example Usage

```hcl
module "instance" {
  source       = "../../modules/instance"
  name         = "my-app-instance"
  image_id     = "5bdb1498-831c-4de0-b7a0-8f63379c96ed"
  flavor_name  = "sx3.12c32r"
  keypair_name = var.keypair_name
  port_id      = module.port.networking_port_id

  depends_on = [
    module.port
  ]
}
```
