# Day 1
## Provisioner Overview
<pre>
- is a Infrastructure as a code tool (Iaac)
- helps create a 
  - new machine with some OS installed in it
  - the machine can be created in on-prem, public/hybrid cloud
  - the machine can be provisioned in VMWare vsphere or similar hypervisors
- Examples
  - Docker
  - Vagrant
  - AWS Cloudformation
  - Terraform
</pre>  

## Configuration Management Overview
<pre>
- is one of the devops tools
- helps automating system admiminstrative tasks like
  - software installation
  - software uninstallation
  - software update/upgrade
  - network administration
  - patching OS
  - managing database servers, webservers, app servers
  - services
  - managing users
    - creating new users
    - modify user credentials
    - delete user
- the expectation is you have machine with some OS pre-installed
  - the OS can be Windows, Linux, Unix, Mac OSX
  - it can Raspberry Pi, it can be machine in onprem environment, it can be machine in public cloud
  - it can be a docker container
  - it can be a podman container
- you could run shell scripts, batches, powershell scripts, etc.,
- However, ansible can one be installed on Linux machine officially ( but it seems to work in Mac as well )
- Ansible can't be installed on Windows machins, but ansible can manage Windows machines
- Examples of Configuration Management Tools
  - Ansible
  - Puppet
  - Chef
</pre>

## Info - Puppet Overview
<pre>
- is a configuration management tool
- is an alternate to Ansible and Chef
- follows client/server architecture
- follows pull based architecture
- the servers that must be managed by Puppet should be installed with Puppet proprietary agent
</pre>

## Ansible Overview
<pre>
- is a configuration management tool developed by Michael Deehan in Python
- the DSL (Domain Specific Language - the language in which automation script is written) used by Ansible is YAML
- YAML is a superset of JSON(JavaScript Object Notation)
- Ansible playbooks can be written by anyone who knows YAML, Python knowledge is not mandatory
- Ansible is agentless, i.e we don't have to install any Ansible Proprietary tools on the Ansible Node
- Ansible Nodes are the servers that we would like to perform configuration management 
- Ansible Nodes cane be Windows Servers, Unix Server, Linux Servers, Mac machine
- Ansible come in 3 flavours
  - Ansible Core 
    - opensource, supports only command-line
    - doesn't support GUI
  - Ansible AWX
    - opensource, supports Web console GUI
    - developed on top of Ansible Core
  - Red Hat Enterprise product 
    - Ansible Tower or Ansible Automation Platform ( Licensed product from Red Hat )
    - developed on to of opensource AWX
    - supports webconsole (GUI)
</pre>

## Info - Ansible Modules
- Ansible modules comes out of the box when we install Ansible in our laptop/desktop
- Ansible modules for Unix/Linux/Mac they are written as Python scripts
- Ansible modules for Windows are written as Powershell scripts
- Each ansible modules does one job
  - Copy module helps copying a file from local machine to remote ansible node or vice versa
  - to install/uninstall softwares in Debain( Ubuntu like Linux distributions ) we have apt ansible module
  - to install/uninstall softwares in Red Hat Linux distributions, we could use yum ansible module
  - to manage services, we could use service ansible module

## Info - Ansible ad-hoc command
<pre>
- ansible ad-hoc command is just invoking a single ansible module without writing a playbook
- usually helps us understand what features a individual modules has, it is way we can experiment a module before using them in playbooks
- each command just invokes one Ansible module
</pre>

## Info - Ansible Playbook
<pre>
- is a YAML file, hence can be written using any plain text editor of your choice
- it follows a specific structure or format
- each Playbook has one or more Play
- each Play targets one to many hosts(ansible nodes to perform the automation)
- each Play runs zero to many Tasks
- each Task runs exactly one Ansible module
- each Play runs zero to many Roles
</pre>

## Info - Ansible Controller Machine ( ACM )
- the machine where Ansible is installed is called Ansible Controller Machine
- officially only a Linux machine can be an Ansible Controller Machine
- Windows machines can't be an Ansible Controller Machine
- the machine where we run the ansible playbook is called Ansible Controller Machine

## Info - Ansible Node ( these are the machines where Ansible will perform the automation )
- it can be a Windows machine
- it can be a Linux machine
- it can be an unix machine
- it can be a Mac machine
- it can be Network routers/switches

## Info - Ansible Inventory
<pre>
- it is a text file, who looks somewhat similar INI file
- it has the connection details of Ansible Nodes
- if the Ansible node happens to be a Window machine, then it will have WinRM connection details, login credentials, etc
- if the Ansible node happens to be a Unix/Linux/Mac/Switches/routers, then it have SSH connection details, login credentials, etc.,
</pre>

## Info - Installing ansible core in Ubuntu
```
sudo apt update
sudo apt install software-properties-common
udo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

## Info - What is Linux Distribution?
<pre>
For example
- Ubuntu - is a Linux distribution
- Fedora - is another Linux distribution
- RHEL - is yet another Linux distribution from Red Hat
- What is common among the above OS?
  - all are Linux OS
  - They all have Linux Kernel which is common
  - Linux Kernel is the core component of any Linux OS
  - The same set of Linux tools are supported by the above linux OS
  - Shells,bash shell, sh shell, etc.
- Different Linux distributions support different Package Managers

</pre>  

## Info - Package manager
<pre>
- is an utility that comes out of box with Linux OS
- depending on which linux family the OS distribution comes under, they may support different package managers
- Debian linux family ( package manager - apt or apt-get 0
  - ubuntu
  - kali
  - Raspberry Pi
- RedHat linux family ( package manager - yum or rpm or dnf 0 
- packager helps us 
  - install softwares
  - uninstall softwares
  - update/upgrade/downgrade softwares
  - downloads the compatible installer binary from Repository Servers
</pre>  

## Info - Linux Repository Server
<pre>
- Each Linux distribution company maintains a website( repository server )
- The repository server could be JFrog Artifactory, Sonatype Nexus, similar products
- Ubuntu 24.04, there can many variants
  - ARM processor
  - Apple silicon
  - Intel 32-bit Processor
  - AMD x64 bit Processor
  - 32 bit OS or 64 bit OS
- Each repository server will provide processor specific pages to download compatible softwares
- Each repository server will provide separate repository url for different version of ubuntu 18.04, 20,04, 22.04, 24,04
- separe page/url for 32-bit/64-bit OS 
</pre>  

## Info - Key pairs ( public and private key )
<pre>
- For any unix/linux/mac user, we can create a key pair(i.e pubilc and private key)
- The private key must be retained on the same machine where it was generated
- The private should never be shared with anyone
- The key pairs are always unique in nature
- The public key will only match with a particular private key
- The public key will be shared to trused unix/linux/mac machines to which one wish to perform ssh login
</pre>

## Lab - Cloning TekTutor training repository ( from linux terminal - one time activity )
```
cd ~
git clone https://github.com/tektutor/terraform-17-21feb-2025.git
cd terraform-17-21feb-2025
```
Each time I push new instructions and/or code samples, you just need to git pull to get the delta changes, hence clone must be done only the first time.

## Lab - Building a Custom Docker Image to provision ansible node containers
Let's generate key pair in the terminal for rps linux user
```
ssh-keygen
```

Expected ouput
![image](https://github.com/user-attachments/assets/9c1dc2a1-43a2-4c07-a45d-df94e505d899)

Let's build a custom build image
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/CustomDockerAnsibleNodeImages/ubuntu-ansible
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/ubuntu-ansible-node:1.0 .
docker images
```
Expected output
![image](https://github.com/user-attachments/assets/a448fd77-e30a-401a-83e9-c09371291cba)
![image](https://github.com/user-attachments/assets/9bd6ac2c-1ed2-49ff-bcfd-939cf263f64d)

## Lab - Let's create couple of containers using the custom image we build in the previous lab exercise
```
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:1.0
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:1.0
```

Listing the currently running containers
```
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/b99fff2a-d1f5-4133-8a63-49f8ffc153f7)

## Lab - Test if we are able to SSH into ubuntu1 and ubuntu2 ansible node containers

When if prompts with a question "Are you sure, do you want to connect to the machine (yes/no) ?" you type "yes".  It shouldn't be prompting for password as we configured public/private login authentication.

```
ssh -p 2001 root@localhost
exit

ssh -p 2002 root@localhost
exit
```

Expected output
![image](https://github.com/user-attachments/assets/a589cc06-eff5-4919-a697-4045fa5292a2)
![image](https://github.com/user-attachments/assets/c3dd46d4-4f41-4545-a166-2d255e4b1b6d)

## Lab - Running your first ansible ad-hoc command to ping ubuntu1 and ubuntu2 ansible container nodes
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible
cat inventory
ansible -i inventory all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/cbf9a1c8-420a-40c6-aa03-a9fea8229b04)


## Lab - Collecting facts about ansible nodes using setup module
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible/playbooks
ansible -i inventory ubuntu1 -m setup
```

Expected output
![image](https://github.com/user-attachments/assets/742969b2-75d5-4e3c-a9c3-63c51e39ee6d)
![image](https://github.com/user-attachments/assets/d3285511-525c-4803-ad85-e95a7a725cbb)

## Lab - Running your first ansible ping playbook
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible
ansible-playbook -i inventory ping-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/44bf3893-9e20-4888-973e-616609ff06c5)

## Lab - Getting help about specific ansible modules
```
ansible-doc -l | grep wc
ansible-doc -l | grep oracle
ansible-doc -l | grep mysql
ansible-doc -l | grep mssql
```

Expected output
![image](https://github.com/user-attachments/assets/f34b01c7-398b-4a00-860f-670295323a42)

## Lab - Multiple inventory
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible
ls
cat inventory
cat hosts
ansible -i inventory all -m ping
ansible -i hosts all -m setup
```

Expected output
![image](https://github.com/user-attachments/assets/7d9ccfb9-4831-4ca7-bdd6-37dcd7279e62)

In the hosts file, we [all:vars] captures all the group variables.  For the machines under all group, the variables listed under [all:vars] is applicable.  

In the hosts file, the ansible_port is a host variable, which has unique values for each hosts.

## Lab - Install nginx into ansible nodes, configure custom folder to pick web pages and deploy custom web page

Running ansible playbook
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible/install-nginx
cat playbook.yml
ansible-playbook -i hosts playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/b24b90f8-348e-44ef-8389-a4b88d2358d5)
![image](https://github.com/user-attachments/assets/3318ca68-8851-48fb-997a-1175bdfd4425)
![image](https://github.com/user-attachments/assets/a6f2f0c8-f430-4140-84a8-97a4bedf66ac)
![image](https://github.com/user-attachments/assets/744e008c-7adc-4b68-8e97-61cdc0e778b7)
![image](https://github.com/user-attachments/assets/db990468-0df5-43ed-8e9c-35935c71c1cd)


Finding the IP Address of ubuntu1 and ubuntu2 ansible node containers
```
docker inspect -f {{.NetworkSettings.IPAddress}} ubuntu1
docker inspect -f {{.NetworkSettings.IPAddress}} ubuntu2
```

Pinging the ansible nodes
```
ping 172.17.0.2
ping 172.17.0.3
```

Checking the nginx service status using ansible ad-hoc command
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible/install-nginx
ansible -i hosts ubuntu1 -m shell -a "service nginx status"
```

Accessing the web page from ubuntu1 and ubuntu2
```
curl http://172.17.0.2:80
curl http://localhost:8001

curl http://172.17.0.3:80
curl http://localhost:8002
```

Expected output
![image](https://github.com/user-attachments/assets/8b2dfdeb-b997-402d-b0a9-00f01837c374)
![image](https://github.com/user-attachments/assets/71446504-0a0f-4c3e-8751-ea09d59513dd)

## Lab - Using ftp, let's copy the default nginx config file to our local machine(Ansible Controller Machine)
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible/install-nginx
sftp root@172.17.0.2:/etc/nginx/sites-available/default .
ls -l
```

Expected output
![image](https://github.com/user-attachments/assets/5d255aef-ccde-43f0-b548-516a35662694)


## Lab - Run the refactored install nginx playbook
```
cd ~/terraform-17-21feb-2025
git pull
cd Day1/ansible/install-nginx
cat playbook.yml
echo "Nginx works!" > index.html
ls -l
cat index.html
ansible-playbook -i hosts playbook.yml

curl http://172.17.0.2
curl http://localhost:8001

curl http://172.17.0.3
curl http://localhost:8002
```

Expected output
![image](https://github.com/user-attachments/assets/fcf12ba5-4b4b-4669-8404-36b234b877b0)
![image](https://github.com/user-attachments/assets/294afc44-b227-4c7e-9d74-e8240278a514)
![image](https://github.com/user-attachments/assets/d3deeaca-72e4-4b5b-9e10-5ec30148a049)
![image](https://github.com/user-attachments/assets/6a82fed8-0355-4537-9a12-20fd98a03c0f)
