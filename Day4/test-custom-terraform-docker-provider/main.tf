terraform {
   required_providers {
       docker = {
          source = "tektutor/docker"
	  version = "1.0.0"
       }
   }
}

resource "docker_container" "hello-ms-container1" {
	image_name = "tektutor/spring-ms:1.0"
	container_name = "hello-microservice-container"
}
