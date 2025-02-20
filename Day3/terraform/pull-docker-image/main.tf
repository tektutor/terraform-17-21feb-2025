terraform {
  required_providers {
    docker = {
       source = "kreuzwerker/docker"
       version = "3.0.2"
    }
  }
}


# Creates an instance of the docker provider
provider docker {

}

data "docker_image" "nginx" {
   name = "nginx:latest"
}

resource "docker_container" "nginx_container1" {
  image = data.docker_image.nginx.name
  name  = "nginx1"
  hostname = "nginx1"
}
