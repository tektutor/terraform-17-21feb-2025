# Day 4

## Lab - Understanding remote-exec and connection block
```
cd ~/terraform-17-21feb-2025
git pull
cd Day4/terraform/remote-exec
terraform init
terraform apply --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/05be49e7-44c6-4eda-be36-e096e9109b52)
![image](https://github.com/user-attachments/assets/34faabd4-f51b-48f9-84b0-45a6ae8183c5)
![image](https://github.com/user-attachments/assets/bf2f1600-f173-48b8-8608-39b42a40ddc6)
![image](https://github.com/user-attachments/assets/11599fc1-3508-4d65-a58a-c41f22d2eb04)

Once you are done with this lab exercise, make sure you dispose the resources by invoking destroy.

```
cd ~/terraform-17-21feb-2025
git pull
cd Day4/terraform/remote-exec
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/881d930c-fdbf-4fe3-a3a2-018d03a3c80b)
![image](https://github.com/user-attachments/assets/777b8a65-8713-4e92-ac52-6a5f27f10c48)

## Lab - Using Docker SDK to manage docker resources in Golang
The go mod tidy command will look for go.mod and scans all the *.go files present in the folder to fetch all import statements.  I then downloads all the dependencies including the indirect dependencies.  In this process, the go.mod gets updated, golang also creates the go.sum file with the download package versions along with their checksums.

```
cd ~/terraform-17-21feb-2025
git pull
cd Day4/docker
go mod tidy
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/0e8d6890-1fe2-44f3-9d82-8642cc192e6a)

## Lab - Developing a custom Terraform Provider to manage text file
```
cd ~/terraform-17-21feb-2025
git pull
cd Day4/terraform-provider-file
```

We need to create a folder structure as shown below to create a custom terraform provider
![image](https://github.com/user-attachments/assets/38c991d8-1bf0-4325-875f-30d131621314)

To configure terraform download the provider from local system.
Create a folder 
```
mkdir -p /home/rps/go/bin
```
We need to create a file name ~/.terraformrc with the below content
<pre>
provider_installation {
  dev_overrides {
      "registry.terraform.io/tektutor/file" = "/home/rps/go/bin"
  }
  direct {}
}  
</pre>

Let's try to build the terraform file provider
```
cd ~/terraform-17-21feb-2025
git pull
cd Day4/terraform-provider-file
go mod tidy
go build -o terraform-provider-file
go install
ls -l /home/rps/go/bin
```

Expected output
![image](https://github.com/user-attachments/assets/17d07dc2-3708-47c7-8c7f-06a97f6aca36)


## Lab - Invoking the custom terraform file provider in a terraform script
```
cd ~/terraform-17-21feb-2025
git pull
cd Day4/test-custom-terraform-file-provider
terraform plan
terraform apply
terraform show
terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/8779008e-81a9-4f34-91f0-e7753c587ec1)
![image](https://github.com/user-attachments/assets/7c40ccf8-8df6-45c1-ae3f-90a84835e9e1)
![image](https://github.com/user-attachments/assets/2688e145-75b3-42b4-bd98-9791e84fb021)

