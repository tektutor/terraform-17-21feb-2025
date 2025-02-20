resource docker_container "my_container" {
  image = "tektutor/ubuntu-ansible-node:1.0"
  name = "container_1"
  hostname  = "container_1"

  ports {
     internal = 22
     external = 2001
  }
}

resource time_sleep "wait_for_45_seconds" {
  create_duration  = "45s"
  destroy_duration = "5s"

  depends_on = [
    docker_container.my_container  
  ]
}

resource local_file "ssh_connection" {
  filename = "sample.txt"
  content  = "dummy"

  provisioner "local-exec" {
	command = "hostname" 
  }

  connection {
     type = "ssh"
     user = "root"
     port =  "2001"
     host = "localhost"
  }

  provisioner "remote-exec" {
    inline = [
        "hostname",
	"hostname -i",
    ]
  }
  depends_on = [
    time_sleep.wait_for_45_seconds  
  ]
}
