# Keypair Module

This module provisions an OpenStack keypair using the `openstack_compute_keypair_v2` resource.
It is used to enable SSH access to instances via a provided public key.

---

## Variables

| Name         | Description                          | Type   | Required |
|--------------|--------------------------------------|--------|----------|
| `name`       | Name of the keypair                  | string | Yes      |
| `public_key` | The public SSH key to import         | string | Yes      |

---

## Outputs

This module does not define any outputs by default.
You may optionally add an output like the keypair name if needed.

---

## Example Usage

```hcl
module "keypair" {
  source = "../../modules/keypair"

  name       = var.keypair_name
  public_key = var.public_key
}
```
