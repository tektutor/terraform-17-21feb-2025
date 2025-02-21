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

## Continuous Integration (CI) Overview
<pre>
- it is a fail-fast engineering process to find issues early 
- when bugs are detected early during development phase, they are easy to fix
- cost of fixing the bugs is also relatively cheaper
- it is similar SCRUM daily stand-up meeting, which is an inspect and adapt meeting
- the team that follows CI/CD, the engineers would be pushing code to version control several times a day
- code would be integrated many times a day
- Jenkins or similar CI/CD server will grab the latest code, they trigger a build, as part of the build, automated test cases would be executed to verify if the new code is as expected, if the new code is breaking any existing functionality.
- the build that was triggered due to new code integration succeeds, it means no functionality is broken, everything works as expected
- continous frequent feedback is given by CI/CD builds, eventually improving the code quaility and functional quality
- CI helps confidently deliver releases in a short amount of time
- Unit and Integration is the scope of CI
</pre>

## Continuous Deployment (CD) Overview
<pre>
- Once the dev release suceeds all the automated test cases added by dev team, it is automatically promoted for QA testing
- the dev certified release binaries are deployed automatically to QA environment for further automated QA testing
- the QA would automate, end to end functionality test, regression test, smoke test, performance test, stress test, component/API test, etc
</pre>

## Continuous Delivery (CD) Overview
<pre>
- the QA certified build(release) is automatically deployed into production or pre-prod environment
- in the pre-prod environment the customer or the Operations team would verify if the new release is working as expected
- especially fintech companies, after manual approval the binaries could go live in production environment
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
- Cloudbees ( Enterprise Jenkins )
- Microsoft Team Foundation Server (TFS)

## Lab - Download Jenkins war file
Download the Generic Java Package (war) file from the left side (LTS)
<pre>
cd ~/Downloads
wget https://get.jenkins.io/war-stable/2.492.1/jenkins.war
</pre>

Expected output
![image](https://github.com/user-attachments/assets/066f4a54-d900-40fd-990d-584c96d0c1d9)
![image](https://github.com/user-attachments/assets/f37f88a3-2082-4d5a-8150-caaeb3e3b0aa)

## Lab - Launching Jenkins from terminal
```
cd ~/Downloads
java -jar ./jenkins.war
```

Expected output
![image](https://github.com/user-attachments/assets/f50e515b-7158-4c30-b576-32378d4522bc)
![image](https://github.com/user-attachments/assets/68a70a4d-568f-45dd-8d32-740320fa519c)

## Lab - Accessing Jenkins Dashboard from your RPS Cloud machine chrome web browser
<pre>
http://localhost:8080  
</pre>

Expected output
![image](https://github.com/user-attachments/assets/543fea33-0ac8-4e3c-817d-40bed8d816d3)

You can copy the initial admin password from the web browser itself as shown below
![image](https://github.com/user-attachments/assets/c3f6675f-2e1e-4e4a-abe4-76ead3154de9)
![image](https://github.com/user-attachments/assets/30b159db-900c-4790-83b8-734b19ed1ddb)
![image](https://github.com/user-attachments/assets/6b12456e-d842-4318-8e8e-729cd4b0035e)
![image](https://github.com/user-attachments/assets/55d0559a-2536-4e2d-8ad3-24771edd5d18)
![image](https://github.com/user-attachments/assets/ad95a77f-377e-4ee3-8b08-641611a943ad)
Select "Install suggested plugins"
![image](https://github.com/user-attachments/assets/7d786bdd-0082-412f-86f7-f228b33fd985)
![image](https://github.com/user-attachments/assets/d8773a01-d2f3-44b8-984c-8961dda0e131)
![image](https://github.com/user-attachments/assets/44b84d06-c58e-4c64-9f78-80117102199f)
![image](https://github.com/user-attachments/assets/115f9217-043b-41db-b824-24b9af753d66)
![image](https://github.com/user-attachments/assets/08430ce7-bf31-4a5a-8c69-511414e49eec)
![image](https://github.com/user-attachments/assets/025f0fb4-d48b-4829-aac6-723bf2171c3d)
Click "Save and Continue"
![image](https://github.com/user-attachments/assets/78b2f63e-98a2-4e55-b043-6921efc60571)
Click "Save and Finish"
![image](https://github.com/user-attachments/assets/b976ee8d-672a-47e3-9f30-873456c757e1)
Click "Start using Jenkins"
![image](https://github.com/user-attachments/assets/55d7757d-aead-45ca-8d2c-ef98e582b263)
![image](https://github.com/user-attachments/assets/78cd113d-9ecb-44a0-8e4b-67b86178ca37)
                                                            
## Lab - Creating a CI/CD DevOps Pipeline

Open Jenkins Dashboard in RPS Lab machine, chrome web browser
```
http://localhost:8080
```
Expected output
![image](https://github.com/user-attachments/assets/6fc5913b-3d17-480d-8f4a-741042319019)

Click "Create job"
![image](https://github.com/user-attachments/assets/d100ba90-8c24-4809-8883-374e52d496b7)
Select "Pipeline", under the "Enter an item name" type "DevOps CICD Pipeline"
![image](https://github.com/user-attachments/assets/21d58d32-5964-43a8-87a7-33f1b6ea601c)
Click "Ok"
![image](https://github.com/user-attachments/assets/939b442c-b793-47a7-b42e-f174e8f9c861)

General Section
![image](https://github.com/user-attachments/assets/54d633c0-8722-4b8b-9be5-ae99d6a4b015)

Trigger Section
![image](https://github.com/user-attachments/assets/088ebc3d-4ad5-4675-9e26-41a1bdeaded0)
Select "Poll SCM", under Schedule type "H/02 * * * *" to configure polling every 2 minutes once.
<pre>
H/02 * * * *  
</pre>
![image](https://github.com/user-attachments/assets/15cd722e-b8a8-4073-92a2-8bf794543d62)


Pipeline section
![image](https://github.com/user-attachments/assets/7b81f0fe-ee3d-43af-8635-9713c77502ba)
Select "Pipeline script from SCM"
![image](https://github.com/user-attachments/assets/08b749dc-2754-4291-adbe-ffcbd45c2ca1)
Under "SCM" select option "Git"
![image](https://github.com/user-attachments/assets/9d13a5c8-a734-4a9b-a171-6c36211ed17f)
![image](https://github.com/user-attachments/assets/3d33cfc4-d3c9-4978-ad60-dd3c68074ccd)
You need to type the Repository URL
<pre>
https://github.com/tektutor/terraform-17-21feb-2025.git  
</pre>

You need to replace Branch specifier from "master" to "main"
![image](https://github.com/user-attachments/assets/4b43799c-c660-4578-9269-a55b9ad1e4cb)

Script Path
![image](https://github.com/user-attachments/assets/edcc63ef-c609-4910-b5ab-be928c2afb71)
Script Path has to modified to 
<pre>
Day5/DevOpsCICDPipeline/Jenkinsfile  
</pre>
![image](https://github.com/user-attachments/assets/63414dfb-aab5-4c3a-ab52-66fb71fa79e1)
Click "Save" at the bottom
![image](https://github.com/user-attachments/assets/d45fa93b-ed20-48b0-90de-716667bdc783)

Status
![image](https://github.com/user-attachments/assets/854db6de-05a7-4d73-8e33-1327f7a87d18)
![image](https://github.com/user-attachments/assets/afbf186e-fa8e-4cab-a933-f6c97893621b)
![image](https://github.com/user-attachments/assets/30c359d8-0516-427e-82e4-66e7d3ad6a3e)
![image](https://github.com/user-attachments/assets/5f1ce0ec-4800-417e-9628-6eff098f2e3f)
![image](https://github.com/user-attachments/assets/b91e1d36-51cc-4765-9e8b-5bf8d4622a47)
![image](https://github.com/user-attachments/assets/ec43469d-c6a8-4b08-a265-c78ce11fca74)
![image](https://github.com/user-attachments/assets/7f790928-1c40-4695-b3f3-bc95c139f1b7)
