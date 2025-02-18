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
cd Day2/ansible/install-nginx
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
![image](https://github.com/user-attachments/assets/83a27421-db1e-4e4a-a395-3d28c8472e97)
![image](https://github.com/user-attachments/assets/ec019f41-f83d-45ce-aad8-212dbd9b395d)

## Info - Ansible Role
<pre>
- helps us reuse code between many playbooks
- roles follows certain best practices and directory structures
- role can't be executed directory
- roles are like Dynamic Link Library(dll), which has some reusable code. Just like dll are loaded into application code, where they are invoke. Roles can invoked thru ansible playbook
- each roles normally focusses on configuration management activities required for a single software
- For example
  - we could an ansible role to install Mysql db server in Ubuntu, Rocky, RHEL, Windows ansible nodes, etc
  - we could an ansible role to install weblogic app server in Ubuntu, Rocky, RHEL and Windows ansible nodes
- there is a tool called ansible-galaxy which can used to download/install opensource roles from galaxy web portal, or we could the tool to develop our own custom ansible role
</pre>

## Lab - Creating a role to install nginx in Ubuntu and Rocky Linux ansible nodes
```
cd ~/terraform-17-21feb-2025
git pull
cd Day2/ansible
ansible-galaxy init nginx
tree nginx
```

Expected output
![image](https://github.com/user-attachments/assets/be7aa7ca-b495-4de6-8614-e49b25b4a1a0)
![image](https://github.com/user-attachments/assets/18a2e8c1-d53b-4299-a96a-0d95ed90affc)

Invoking the ansible role from an ansible playbook
```
cd ~/terraform-17-21feb-2025
git pull
cd Day2/ansible
cat playbook.yml
ansible-playbook -i hosts playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/7fcfdc81-5aa1-4b32-9d2c-74d61929cf2d)
![image](https://github.com/user-attachments/assets/04dc195a-f796-4751-aa79-9e3d46012560)
![image](https://github.com/user-attachments/assets/6dc063a1-46ac-4aad-b154-934ef2c8bcfd)
![image](https://github.com/user-attachments/assets/783a1ede-f562-4456-b692-f5e498e5e817)

## Info - Ansible vault
<pre>
- ansible-vault can help in encrypting/decrypting sensitive information like password credentials, etc
- the vault protected information can be retrieved safely and used within playbook
- using ansible vault one can
  - create a vault file storing sensitive data
  - view the password protected sensitive data providing the password
  - edit the vault protected file
  - encrypt any text file
  - decrypt already ansible vault encrypted file
  - change the password of ansible vault file
</pre>

## Lab - Ansible vault
```
ansible-vault create mysql-credentials.yml
ansible-vault view mysql-credentials.yml
ansible-vault edit mysql-credentials.yml
```
Expected output
![image](https://github.com/user-attachments/assets/5b63fc8c-8cb3-4997-ab9f-a830da8d8486)

## Lab - Using vault-protected data in playbook

When ansible prompts for password type 'root' without quotes
```
cd ~/terraform-17-21feb-2025
git pull
cd Day2/ansible/vault
cat mysql-credentials.yml
ansible-playbook playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/2aeec98d-7ea8-460a-91fa-fc9c5e9fd175)
![image](https://github.com/user-attachments/assets/6f538e8a-8fcf-447b-a757-cfc54b7529c1)

## Lab - Starting Ansible Tower (aka Ansible Automation Platform)
```
minikube status
minikube start
kubectl get pods -n ansible-awx
kubectl get svc -n ansible-awx
minikube service awx-demo-service --url -n ansible-awx
```

Expected output
![image](https://github.com/user-attachments/assets/38345bed-47d7-4570-8f7a-bd426efefd30)
![image](https://github.com/user-attachments/assets/18d8356c-0499-44c2-b2c2-59e35aee1b8c)


Ansible Tower Login Credentials
Retrieving Ansible Tower password
```
kubectl get secret -n ansible-awx | grep -i password
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

<pre>
username - admin
password - 
</pre>

Expected output
![image](https://github.com/user-attachments/assets/2047fc10-8373-4168-a630-f6d9631f9e41)
![image](https://github.com/user-attachments/assets/bd2aeeb4-73d3-422c-89a5-5515329fb12b)
