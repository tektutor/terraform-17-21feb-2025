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
