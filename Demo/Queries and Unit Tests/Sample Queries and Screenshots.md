# Overview

The following examples are drawn from the [Demo of Playbook](https://github.com/paulejarvis/Playbook/tree/master/Demo) and demonstrate how customers could derive value from querying the integrated process data to better understand the dataset or check that the processes fulfill compliance requirements ("Unit Tests"). These are just examples - please see the [Playbook Ontology](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Playbook%20Ontology.md) and [Neo4j Cheatsheet](https://neo4j.com/docs/cypher-refcard/current/) for more details on the data structure and Cypher query language.

## Compliance Queries


* **Scenario** The Chief Compliance officer (CCO) wants to know who is accountable for compliance with all policies (internal and external) across the organization
* **Query** Return all processes that are regulated, how much time each step takes, and who owns the steps in the workflow. This query will ensure clear ownership and accountability for achieving compliance

```Cypher
MATCH (People)-[:RESPONSIBLE]->(Process)
WHERE Process.policy IS NOT NULL
RETURN People.role, Process.workflow, Process.description, Process.policy, Process.time
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/All_regulated_processes.PNG)

* **Scenario** The CFO wants to know how much time is being spent on SOX compliance across the company so that he / she can decide whther to hire a consultant to help streamline these processes
* **Query** Return all processes labeled as SOX

```Cypher
MATCH (People)-[:RESPONSIBLE]->(Process)
WHERE Process.policy = "Sarbanes-Oxley"
RETURN People.role,Process.workflow, Process.description, Process.policy, Process.time
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/SOX_Query.PNG)

* **Scenario** The CCO wants to know who is accountable and responsible for enforcing policies internal to the company and which workflows these policies apply to. NOTE: the differences between "Accountable" and "Responsible" are covered in the [Playbook Ontology](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Playbook%20Ontology.md)
* **Query** Search the relationships between people and processes, filter to just processes that are governed by internal policies, and return the people connected as Responsible or Accountable

```Cypher
MATCH
	(People)-[r]->(Process)
WHERE
	Process.policy =~ 'Internal policy*' AND
  type(r) in ['RESPONSIBLE','ACCOUNTABLE']
RETURN People.role, Process.workflow, Process.description, Process.policy, type(r)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/Internal%20policy%20owners.PNG)

## IT Planning

* **Scenario** The CIO wants to understand the current use of systems across the organization (broken out by department and workflow) to contextualize IT investment and understand the return on investment and what might happen if a system is replaced
* **Query** Search for technologies linked to processes and return the department, application name, and workflows

```Cypher
MATCH (Process)-[r]->(Technology)
WHERE Technology.application is not null
Return Technology.application, Process.department, collect(Distinct Process.workflow)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/Tech%20systems%20by%20department%20and%20workflow.PNG)

* **Scenario** A systems administrator needs to understand what level of access each user requires to do their job so that permissions are limited appropriately. This report can also be used to verify appropriate separation of duties in an audit
* **Query** Query the path from process to technology and from people to process to understand the actions that each user needs to take in each system

```Cypher
MATCH
	(Process)-[r]->(Technology),
  (People)-[s]->(Process)
WHERE
	type(r) in ["CREATE", "DELETE", "UPLOAD"] AND
  People.role is not null
RETURN People.role ,Process.department, Process.workflow, type(r), Technology.application
ORDER BY People.role
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/Tech%20permissions%20by%20role.PNG)

## Operational Efficiency

* **Scenario** The Director of Engineering is worried that the team is not hiring candidates fast enough (and that it's taking too much of the interviewer's time) and wants to understand the time requirements of hiring a single candidate
* **Query** Search the workflow "Engineering interviews" and return time per process along with the owner of that process

```Cypher
MATCH (People)-[r]->(Process)
WHERE Process.workflow = "Engineering interviews" AND People.role IS NOT NULL
RETURN People.role, collect(Process.time)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/Time%20by%20role%20for%20Engineering%20Interviews.PNG)

* **Scenario** A manager wants to understand what might happen if someone on their team leaves - what actions are they currently owning and what would need to be backfilled? What would the replacement resource need to know how to do? This report could also serve as the blueprint for training new hires in their role
* **Query** Search across all workflows and categorize each People --> Process relationship by the type of relationship (Accountable, Responsible, Participant, Reviewer, or Consulted) and group by role

```Cypher
MATCH
	(People)-[r]->(Process)
WHERE
	type(r) in ["ACCOUNTABLE", "RESPONSIBLE", "PARTICIPANT", "REVIEWER", "CONSULTED"]
RETURN People.role, Process.workflow, Process.description, type(r)
ORDER BY People.role
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/Responsibilities%20by%20role.PNG)

## Integrating Workflows Across Departments

* **Scenario** The Director of HR believes that the on-boarding process could be improved but knows that there are actions outside of HR (such as provisioning a laptop or setting up payroll) so wants to search the entire workflow to understand all steps across all departments
* **Query** Query multiple workflows (across departments) and return the end-to-end integrated workflow

```Cypher
MATCH (Process)
WHERE Process.workflow IN
	['Payroll set up', 'New hire on-boarding', 'Asset provisioning' ] OR
   Process.description = "Candidate hired"
RETURN Process
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/New%20hire%20processes%20across%20departments.PNG)
![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Screenshots/New%20hire%20processes%20across%20departments%20(1).PNG)
