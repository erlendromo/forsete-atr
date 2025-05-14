# IP Module

This module provisions a floating IP in OpenStack and associates it with a specified networking port using the `openstack_networking_floatingip_v2` and `openstack_networking_floatingip_associate_v2` resources.

---

## Variables

| Name                    | Description                                        | Type   | Required |
|-------------------------|----------------------------------------------------|--------|----------|
| `external_network_name` | Name of the external network for the floating IP   | string | Yes      |
| `port_id`               | ID of the networking port to associate with the IP | string | Yes      |

---

## Outputs

| Name | Description                     |
|------|---------------------------------|
| `ip` | The provisioned floating IP     |

---

## Example Usage

```hcl
module "ip" {
  source = "../../modules/ip"

  external_network_name = var.external_network_name
  port_id               = module.port.networking_port_id

  depends_on = [
    module.port
  ]
}
```
