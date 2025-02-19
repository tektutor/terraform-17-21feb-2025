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

resource "docker_image" "ubuntu" {
   name = "ubuntu:20.04"
}
