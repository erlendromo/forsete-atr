# Main Entry Point Terraform Configuration

This configuration is the entry point for provisioning the infrastructure for both the application and frontend components using Terraform. It initializes two modules: one for the application and another for the frontend. It also handles the OpenStack credentials and public key management.

---

## Architecture

This configuration provisions:

- **Application Infrastructure**: Deploys the backend application in an OpenStack environment with network, security, compute, and floating IP configuration.
- **Frontend Infrastructure**: Deploys the frontend application in a separate OpenStack environment with similar networking, security, and compute configurations.

Each module (application and frontend) is configured with the appropriate external network name and key pair, allowing secure SSH access and the necessary network connectivity.

---

## Modules

1. **App Module**
   - Provisions the backend application infrastructure with a private network, security groups for SSH and application ports, and a virtual machine (VM).
   - Associates a floating IP to make the application publicly accessible (only accessible through the NTNU network or using VPN).

2. **Frontend Module**
   - Provisions the frontend application infrastructure with a similar setup to the backend, but on a separate network, allowing independent scaling and access (publically accessible for everyone).

---

## Variables

| Name                            | Description                                                   | Type   | Required |
|---------------------------------|---------------------------------------------------------------|--------|----------|
| `application_credential_id`     | The OpenStack Application Credential ID                       | string | Yes      |
| `application_credential_secret` | The OpenStack Application Credential Secret                   | string | Yes      |
| `public_key`                    | The public SSH key used to create key pairs for the instances | string | Yes      |

---

## Outputs

- **`app_ip`**: The floating IP address allocated for the backend application.
- **`frontend_ip`**: The floating IP address allocated for the frontend application.

---

## Example Usage

```hcl
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
```
