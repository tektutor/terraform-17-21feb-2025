module "create-docker-containers" {
   source="./child_module/"
   container_count = var.container_count
}
