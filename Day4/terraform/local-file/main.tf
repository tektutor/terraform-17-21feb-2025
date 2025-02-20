resource "local_file" "myfile" {
  content = "Sample text in file"
  filename = "./sample.txt"
  file_permission = "0644"
}
