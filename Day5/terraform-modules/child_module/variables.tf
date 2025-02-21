variable docker_image_name {
  description = "name of the docker image"
  default     = "tektutor/ubuntu-ansible-node:1.0"
}

variable "container_count" {
   description = "Number of containers you wish to create"
   type = number

   validation {
     condition = var.container_count > 1 && var.container_count <= 5
     error_message = "The container_count should be greater than 0 and must be less than or equal to 5"
   }
   
}
