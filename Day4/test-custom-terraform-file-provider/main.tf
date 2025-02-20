terraform {
   required_providers {
	localfile = {
	   source = "tektutor/file"
	}
   }
}

resource "localfile" "myfile" {
  file_name = "./somefile1.txt"
  file_content = "Test file - some content"
}
