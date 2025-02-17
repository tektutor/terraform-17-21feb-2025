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
