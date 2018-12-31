# Overview

The following examples are drawn from the [Demo of Playbook](https://github.com/paulejarvis/Playbook/tree/master/Demo) and demonstrate how customers could derive value from querying the integrated process data to better understand the dataset or check that the processes fulfill compliance requirements ("Unit Tests"). These are just examples - please see the [Playbook Ontology](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Playbook%20Ontology.md) and [Neo4j Cheatsheet](https://neo4j.com/docs/cypher-refcard/current/) for more details on the data structure and Cypher query language.

## Compliance Queries

* Return all processes that are regulated, how much time each step takes, and who owns the steps in the workflow

```Cypher
MATCH (People)-[:RESPONSIBLE]->(Process)
WHERE Process.policy IS NOT NULL
RETURN People.role, Process.workflow, Process.description, Process.policy, Process.time
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/All_regulated_processes.PNG)

* How much time is being spent on SOX compliance? Who's time is it? This might inform a decision to hire a consultant to streamline the processes

```Cypher
MATCH (People)-[:RESPONSIBLE]->(Process)
WHERE Process.policy = "Sarbanes-Oxley"
RETURN People.role,Process.workflow, Process.description, Process.policy, Process.time
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/SOX_Query.PNG)

* Who is accountable or responsible for enforcing internal policies? NOTE: the differences between "Accountable" and "Responsible" are covered in the [Playbook Ontology](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Playbook%20Ontology.md)

```Cypher
MATCH
	(People)-[r]->(Process)
WHERE
	Process.policy =~ 'Internal policy*' AND
  type(r) in ['RESPONSIBLE','ACCOUNTABLE']
RETURN People.role, Process.workflow, Process.description, Process.policy, type(r)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Internal%20policy%20owners.PNG)

## IT Planning

* Which technology systems are being used across the business? Break it out by department and workflow - this contextualizes IT investments across the org and helps to justify the value of these investments as well as provide information on what might happen if a system was taken down or replaced

```Cypher
MATCH (Process)-[r]->(Technology)
WHERE Technology.application is not null
Return Technology.application, Process.department, collect(Distinct Process.workflow)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Tech%20systems%20by%20department%20and%20workflow.PNG)

* What permissions / roles does each person need to accomplish their job? This view enables administrators to properly limit permissions and systems access and complete separation of duties audits

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

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Tech%20permissions%20by%20role.PNG)

## Operational Efficiency

* How much time does it take to move an engineering candidate all the way through the interview process? How much time is required from each role? Where can we make the process more efficient?

```Cypher
MATCH (People)-[r]->(Process)
WHERE Process.workflow = "Engineering interviews" AND People.role IS NOT NULL
RETURN People.role, collect(Process.time)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Time%20by%20role%20for%20Engineering%20Interviews.PNG)

* What process steps are owned by each role? If someone currently in that role leaves or needs to be backfilled, what does the replacement need to be capable of doing? This can also serve as the blueprint for training new hires in their roles

```Cypher
MATCH
	(People)-[r]->(Process)
WHERE
	type(r) in ["ACCOUNTABLE", "RESPONSIBLE", "PARTICIPANT", "REVIEWER", "CONSULTED"]
RETURN People.role, Process.workflow, Process.description, type(r)
ORDER BY People.role
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Responsibilities%20by%20role.PNG)

## Integrating Processes Across Departments

* What steps are required for a new hire to complete on-boarding (spanning across departments)? Are we preparing new hires adequately and front-loading any administrative tasks?

```Cypher
MATCH (Process)
WHERE Process.workflow IN
	['Payroll set up', 'New hire on-boarding', 'Asset provisioning' ] OR
   Process.description = "Candidate hired"
RETURN Process
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/New%20hire%20processes%20across%20departments.PNG)
![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/New%20hire%20processes%20across%20departments%20(1).PNG)
