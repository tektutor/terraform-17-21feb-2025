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

