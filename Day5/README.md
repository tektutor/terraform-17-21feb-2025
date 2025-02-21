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
