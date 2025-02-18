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

## Lab - Let's create couple of rocky ansible node containers using our custom rocky linux docker image
In case rocky containers are already there, you may delete as shown below
```
docker rm -f rocky1 rocky2
```

Now you should be able to recreate rocky1 and rocky2
```
docker run -d --name rocky1 --hostname rocky1 -p 2003:22 -p 8003:80 tektutor/rocky-ansible-node:1.0
docker run -d --name rocky2 --hostname rocky2 -p 2004:22 -p 8004:80 tektutor/rocky-ansible-node:1.0
```

Expected output
![image](https://github.com/user-attachments/assets/23245ea3-83e8-4200-bc6f-6278bd9ffd56)

## Lab - Let's try to do SSH into the rocky ansible node containers
```
ssh -p 2003 root@localhost
exit

ssh -p 2004 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/c0f919dc-f5d5-4fc2-95fe-68546790af42)

## Lab - Let's add rocky ansible node containers to ansible static inventory
```
cd ~/terraform-17-21feb-2025
git pull
cd Day2/ansible
cat hosts
```

Expected output
![image](https://github.com/user-attachments/assets/50ecbe2c-fb70-4aa0-bd34-1d4ba68bbb52)
![image](https://github.com/user-attachments/assets/843e3e7c-34c3-4fa1-ba1c-cfca911c977b)

## Lab - Ansible static inventory vs dynamic inventory
```
cd ~/terraform-17-21feb-2025
git pull
cd Day2/ansible
cat hosts
cat dynamic-inventory.py
```

Expected output
![image](https://github.com/user-attachments/assets/e261dac4-302b-4df7-8983-b82cd904dbb2)
![image](https://github.com/user-attachments/assets/39aa7f6c-4543-4428-bf21-53e4d1428d73)
![image](https://github.com/user-attachments/assets/1b70ef4e-caa6-4a7e-8cf0-4423cea3913e)


We can run ansible ad-hoc command using static or dynamic inventory
```
ansible -i hosts all -m ping
ansible -i ./dynamic-inventory.py all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/7c854ea3-7528-47ce-a673-73e0ea286eae)
![image](https://github.com/user-attachments/assets/a68bacd7-df58-42ac-be9c-b49ed8285368)

## Lab - Running the playbook on different Linux distributions
```
cd ~/terraform-17-21feb-2025
git pull
cd Day2/ansible
cat playbook.yml
ansible-playbook -i hosts playbook.yml
curl http://localhost:8001
curl http://localhost:8002
curl http://localhost:8003
curl http://localhost:8004
```

Expected output
![image](https://github.com/user-attachments/assets/473b6aee-1484-4b08-aab7-90de543f07c5)
![image](https://github.com/user-attachments/assets/67e4cc17-82be-47e5-9b36-ec114be7fd06)
![image](https://github.com/user-attachments/assets/eddaa3b0-19ac-4a03-b95e-8834da6d6986)
![image](https://github.com/user-attachments/assets/4656a5e5-45b5-423f-bcf4-3170ce43d161)
![image](https://github.com/user-attachments/assets/803bc3e7-01cc-41a8-b269-bc6cfcc36726)
![image](https://github.com/user-attachments/assets/fe071165-14cb-4cf5-b3bc-acbd597d29dd)
![image](https://github.com/user-attachments/assets/eee97125-d443-4863-baa5-4ab2c2a00963)
