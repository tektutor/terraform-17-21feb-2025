# Day 2

## Lab - Building a custom rocky linux ansible node container
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/CustomDockerAnsibleNodeImages/rockylinux-ansible
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/rocky-ansible-node:1.0 .
docker images
rm authorized_keys
```

Expected output
![image](https://github.com/user-attachments/assets/a0993d89-2de2-4d78-825e-dd431076f48f)
![image](https://github.com/user-attachments/assets/1914a548-6cb4-4443-973e-8a0a8fbb53e9)
![image](https://github.com/user-attachments/assets/72d434b9-b007-48f4-ac4c-31b59e337616)
![image](https://github.com/user-attachments/assets/6390a786-021c-44b7-9674-8968ca6b66d3)
