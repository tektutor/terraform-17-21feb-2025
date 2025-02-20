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
