tech requirements per roles - this informs separation of duties & IT controls

what are the process requirements per role?

are my internal policies being applied at the key moments

what additional processes are connected to my on-boarding


## Compliance

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

## IT Planning

* Which technology systems are being used across the business? Break it out by department and workflow

```Cypher
MATCH (Process)-[r]->(Technology)
WHERE Technology.application is not null
Return Technology.application, Process.department, collect(Distinct Process.workflow)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Tech%20systems%20by%20department%20and%20workflow.PNG)

## Operational Efficiency

* How much time does it take to move an engineering candidate all the way through the interview process? How much time is required from each role?

```Cypher
MATCH (People)-[r]->(Process)
WHERE Process.workflow = "Engineering interviews" AND People.role IS NOT NULL
RETURN People.role, collect(Process.time)
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/Time%20by%20role%20for%20Engineering%20Interviews.PNG)

## Cross-functional Processes

* What steps are required for a new hire to complete on-boarding (spanning across departments)?

```Cypher
MATCH (Process)
WHERE Process.workflow IN
	['Payroll set up', 'New hire on-boarding', 'Asset provisioning' ] OR
   Process.description = "Candidate hired"
RETURN Process
```

![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/New%20hire%20processes%20across%20departments.PNG)
![](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries/Screenshots/New%20hire%20processes%20across%20departments%20(1).PNG)
