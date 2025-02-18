![image](https://github.com/user-attachments/assets/06a409a3-d122-4e53-8575-7c412d65f443)# Day 2

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

## Lab - Creating a project in Ansible Tower
Open your Ansible Tower Dashboard
![image](https://github.com/user-attachments/assets/d8b7b733-149f-4bcd-a8c3-daad24ecc48c)

Navigate to Resources --> Projects
![image](https://github.com/user-attachments/assets/df7ae704-407b-4a1b-b52f-d9ae44fbb862)
Click "Add"
![image](https://github.com/user-attachments/assets/d129342b-94b0-4bc7-aafc-d6a9f33b52ed)
![image](https://github.com/user-attachments/assets/3f3c8ee1-6fba-4e3c-9529-7fd933289ccb)
![image](https://github.com/user-attachments/assets/e69530da-5f18-418a-921c-ec5b0cd426e5)
![image](https://github.com/user-attachments/assets/404714b8-602e-4163-ba97-a3160c082f9c)
![image](https://github.com/user-attachments/assets/6e198a0a-1208-4bd7-99ac-0f9795904635)
Click "Save"
![image](https://github.com/user-attachments/assets/bca2d424-a33a-4de5-b193-39f9819b3015)
![image](https://github.com/user-attachments/assets/f21840f7-9d84-499f-9977-2f841111533f)
![image](https://github.com/user-attachments/assets/c9a2f9ef-87ee-44c3-b050-6c6933652316)

## Lab - Create a credential to store Private Key within Ansible Tower
Open your Ansible Tower Dashboard
![image](https://github.com/user-attachments/assets/d8b7b733-149f-4bcd-a8c3-daad24ecc48c)

Navigate to Resources --> Credentials
![image](https://github.com/user-attachments/assets/f538e21b-6e0c-4208-a684-049e72b1f9ad)
Click "Add"
![image](https://github.com/user-attachments/assets/0360f6d3-5e47-4452-b715-6f210cf51f67)
![image](https://github.com/user-attachments/assets/2d8ca5ce-bdd3-4bda-b970-24569d71781e)
```
cat ~/.ssh/id_ed25519
```
![image](https://github.com/user-attachments/assets/94af92dd-675c-4266-80d7-ec25ed546d83)
![image](https://github.com/user-attachments/assets/ecd566b0-dd7a-4403-a6db-7cb8189fb142)
![image](https://github.com/user-attachments/assets/96a1959e-69e4-4f4a-a273-9a1fae432312)
Click "Save"
![image](https://github.com/user-attachments/assets/b3565331-9430-4494-8e44-146f4a7d126e)

## Lab - Let's create an Inventory in Ansible Tower
Open your Ansible Tower Dashboard
![image](https://github.com/user-attachments/assets/d8b7b733-149f-4bcd-a8c3-daad24ecc48c)

Navigate to Resources --> Inventory
![image](https://github.com/user-attachments/assets/0fe420e8-aafd-4d8e-8dc3-5129f85e08eb)
Click "Add"
![image](https://github.com/user-attachments/assets/e4dbb2a2-238a-4908-ae39-9a7920a967f4)
Select "Add inventory"
![image](https://github.com/user-attachments/assets/387cb834-b345-48db-972f-8274bb425941)
![image](https://github.com/user-attachments/assets/8270acb9-70cf-4d73-8840-ec4d980e42a4)
Click "Save"
![image](https://github.com/user-attachments/assets/09d377c9-76c9-483a-9a91-bee1bc6a0da9)

In order add "Hosts" into the Inventory, click on "Hosts" Tab
![image](https://github.com/user-attachments/assets/2aaf4bdc-3cf7-4ff3-a7b8-321fa45fea78)
![image](https://github.com/user-attachments/assets/12b80d14-04d1-4330-867c-1eeb3452ecfa)
Click "Add"
![image](https://github.com/user-attachments/assets/8cbc6360-5533-4c4e-a865-7bb9a97ffb99)
Find the IP address of your lab machine
```
sudo apt install -y net-tools
ifconfig ens160
```
![image](https://github.com/user-attachments/assets/f65585e9-523d-4221-9a18-9efe1da8da77)
![image](https://github.com/user-attachments/assets/a2c981cf-2b35-4a8b-9dee-64dde29c6dc7)
```
ansible_port: 2001
ansible_user: root
ansible_host: 192.168.2.176
```
![image](https://github.com/user-attachments/assets/f6ddfcfd-0b0b-4d5c-906e-2905bb05d0c6)
Click "Save"
![image](https://github.com/user-attachments/assets/2b611532-d4a3-4a31-82f6-f8ea649e1280)

Let's add Ubuntu2
```
ansible_port: 2002
ansible_user: root
ansible_host: 192.168.2.176
```
![image](https://github.com/user-attachments/assets/fb65728f-871d-468f-b47d-7a8333cb3d04)
Click "Add"
![image](https://github.com/user-attachments/assets/6e760a12-7fb1-438a-9368-4d5cd42b5479)
Click "Save"
![image](https://github.com/user-attachments/assets/1653d574-5f27-4a8f-93cb-1411cecbcd2e)

Let's add Rocky1
```
ansible_port: 2003
ansible_user: root
ansible_host: 192.168.2.176
```
![image](https://github.com/user-attachments/assets/dbaaf328-f421-45b7-8b5b-ea8e725567df)
Click "Save"

Let's add Rocky2
![image](https://github.com/user-attachments/assets/e7637451-0164-48d3-a0cf-5692e6e1f329)
![image](https://github.com/user-attachments/assets/9ac0e090-9045-44a4-8eaf-1ec29720d7fe)
Click "Save"
![image](https://github.com/user-attachments/assets/0bfc5abb-83f7-46d8-ae0e-8c84510dd6ff)

![image](https://github.com/user-attachments/assets/185eadf9-90f3-45e7-86e2-f58406a8cb40)

## Lab - Create a Job Template to invoke Ansible Playbook from Ansible Tower
Open your Ansible Tower Dashboard
![image](https://github.com/user-attachments/assets/d8b7b733-149f-4bcd-a8c3-daad24ecc48c)

Navigate to Resources --> Templates
![image](https://github.com/user-attachments/assets/2ad8e99a-7228-450d-ae26-1f780e152c1a)
Click "Add"
![image](https://github.com/user-attachments/assets/1894b68f-930e-4fc4-9acf-a2a055a7221a)
Select "Add Job Template"
![image](https://github.com/user-attachments/assets/a3c22fec-5fce-4359-aac4-2e80e2282cbd)
![image](https://github.com/user-attachments/assets/3574dc41-bd76-4ce0-9b69-9b124c07ece2)
![image](https://github.com/user-attachments/assets/0f451e49-9d78-4040-b393-e512cb203d0d)
![image](https://github.com/user-attachments/assets/03b5fc67-d46a-45cf-8a52-f608cde3e250)
![image](https://github.com/user-attachments/assets/a731c9ab-da0e-432c-85fa-1c4cd969bdb8)
![image](https://github.com/user-attachments/assets/de680afc-cc1d-4175-a81a-179936d1d3a6)
![image](https://github.com/user-attachments/assets/0efcfcf6-020c-4313-8c6f-c5ac80462092)
Click "Save"
![image](https://github.com/user-attachments/assets/55a599e2-09d7-4ff7-a17a-7bd7976c99d2)
![image](https://github.com/user-attachments/assets/5908d44b-be86-42f0-be27-356802cd3fd9)
![image](https://github.com/user-attachments/assets/b3ba31b4-b7de-4901-9e49-815a77efd4d1)
![image](https://github.com/user-attachments/assets/594bac11-901a-49e0-8746-f34f9b5c2cd3)

## Lab - Installing tower cli
```
sudo apt install -y python3-pip
pip install ansible-tower-cli --break-system-packages
```

Expected output
![image](https://github.com/user-attachments/assets/a5800d97-d9f0-4229-8476-cc4d15eed5ed)
![image](https://github.com/user-attachments/assets/95b1f71d-0dbe-42fb-a6a8-cae788fbafb9)

## Lab - Using tower-cli
```
export PATH=$PATH:/home/rps/.local/bin
tower-cli config host https://192.168.49.2/31225
tower-cli config username admin
tower-cli config password your-admin-password
tower-cli config verify_ssl false
tower-cli project list
```

## Info - Golang Overview
<pre>
- golang is developed by Google in C Programming language
- it is faster than any interpretter based languages like Ruby, Python, etc.,
- it is also faster than most compiler based programming languages
- it is a compiler
- golang syntax resembles like C programming language
- golang remove all confusing features, keeping only useful, non-confusing features
- golang avoid memory management issues we have in C/C++
- golang support about 25 keywords
</pre>
