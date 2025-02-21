# Day 5

## Info - Terraform Modules Overview
<pre>
- collection of many terraform configuration files(.tf) files in a dedicated directory
- modules encapsulate group of terraform resources related to a single infrastructure/task
- modules allows us to reuse code
- the terraform scripts kept in the top-level folder is referred as root module
- the terraform scripts kept in the top-level sub-folder is referred as child module
- a root module may have zero to many child modules
</pre>

## Lab - Terraform Root and Child Modules
```
cd ~/terraform-17-21feb-2025
git pull
cd Day5/terraform-modules
tree
terraform init
terraform plan
terraform apply --auto-approve
docker ps
terraform show

terraform destroy --auto-approve
```

Expected output
![image](https://github.com/user-attachments/assets/3bfe367c-3e5e-4a2a-a24e-cb821c70f8ec)
![image](https://github.com/user-attachments/assets/04636db8-6024-4dad-be85-141a8e7820a9)

![image](https://github.com/user-attachments/assets/63850d97-5964-4a13-aa55-2cdb7355b181)
![image](https://github.com/user-attachments/assets/550ed964-2115-42e4-858c-d7a1351ff6c8)
![image](https://github.com/user-attachments/assets/290fdd0b-6e47-4a27-930d-ce6502b1e680)
![image](https://github.com/user-attachments/assets/5daf9085-e563-41b3-b118-3b14b0414fa0)

![image](https://github.com/user-attachments/assets/2bb9a9b8-b705-4781-b0bb-7855ca3e0ae0)
![image](https://github.com/user-attachments/assets/49729673-d5cc-4624-b454-c0d90f554ae1)

![image](https://github.com/user-attachments/assets/b15db064-3d9f-4039-80ba-cbc961c0382a)
![image](https://github.com/user-attachments/assets/3eba987e-affd-439c-ac42-b8745fd0b04b)
![image](https://github.com/user-attachments/assets/735cc554-1b52-4f8e-83f4-38c621d0e960)
![image](https://github.com/user-attachments/assets/a867f354-8720-4547-b823-95017a23cef4)
![image](https://github.com/user-attachments/assets/93b91c00-af5c-4e40-8972-f3b21bed5152)
![image](https://github.com/user-attachments/assets/77b5f3c5-2191-4053-83d2-7af87cacd286)


## Info - Configuration Drift
<pre>
- system configurations gradually deviate from their desired or documented state 
- happens when changes are made to software or infrastructure settings over time without a proper change management process
- occurs when system configuration is updated manually, often without governance 
- this affects individual machines, software configurations, clusters, or entire IT systems. 
- can have serious consequences, such as inconsistent configurations that cause unpredictable system behavior and increased difficulty in troubleshooting issues
</pre>

## Info - Commons causes of Configuration Drift
<pre>
- manual changes
- inconsistent and manual deployments
- external dependencies
- difference in environments
- lack of version control
- poor or lack of documentation
</pre>

## Info - Risks associated with Configuration Drifts
<pre>
- Security Vulnerabilities
- Performance issues
- Compliance issues
- Makes troubleshooting very difficult
- increased down-time
- decreased reliability
- poor user-experience
</pre>

## Info - Solution to Configuration Drifts
<pre>
- Using Version Control
- Continuous Integration
- Use Configuration Management Tools to override manual changes in continuous fashion
- Use Infrastructure as Code Tools to override manual changes
- Regularly monitor and audit infrastructure
</pre>

## Info - Monitoring
<pre>
- is the process of collecting and analyzing data from IT infrastructure, system and processes
- using the data to improve business outcomes and drive value to the organization
- collects data to help keep your infrastructure up and running without any downtime
- Tools to
  - Data Collection (Logs)
  - Alerting
  - Remediation
  - Agent based monitoring
  - Agentless monitoring
  - Collecting Metrics
- Examples
  - Dynatrace
  - DataDog
  - Splunk
  - Prometheus & Grafana
</pre>

## What is Jenkins?
- is a build automation server
- used mainly for CI/CD Builds
- this was developed in Java by Kohsuke Kawaguchi, former employee of Sun Microsystems
- Initially it was developed as Hudson is an opensource project
- Then Oracle acquired Sun Microsystems, then part of Hudson including Kohsuke Kawaguchi had quit Oracle
  created a new branch from Hudson called Jenkins
- The other part of the Hudson team they continue to develop the product as Hudson
- There is lot common code between Hudson and Jenkins
- More than 10000 active contributors are there for Jenkins
- Many other software vendors got inspired by Jenkins similar products came out in market like Bamboo, Team City, Microsoft TFS, etc.,
- Jenkins supports CI/CD build for products built in any software stack
  
## What is Cloudbees?
- Cloudbees is the enterprise paid variant of Jenkins
- Feature wise Jenkins and Cloudbees pretty much they are same
- We get support for Cloudbees while we only get community support for Jenkins
- Cloudbees is more stable as it is a paid version
  
## Jenkins Alternatives
- Bamboo
- Team City
- Cloudbees
- Microsoft Team Foundation Server (TFS)
