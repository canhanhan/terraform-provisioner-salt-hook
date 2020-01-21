# terraform-provisioner-salt-hook

Experimental Terraform provisioner that triggers a hook request on SaltStack rest_cherrypy NetAPI module.

terraform-provisioner-salt-hook requires Go version 1.13 or greater.

## Usage

```terraform
resource "null_resource" "test" {
    provisioner "salt-hook" {
        address = "https://salt-master:8000"
        username = "test"
        password = "test"
        backend = "pam"
        id = "test"
        hebele = {
            sample_key = "sample_data"
        }
    }
}
```
