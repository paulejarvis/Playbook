tech requirements per roles - this informs separation of duties & IT controls

what are the process requirements per role?

What processes is each system involved in?

what regulatory requirements apply to my hiring process?
* search all nodes in recruiting / Engineering, to see where control framework is not null

are my internal policies being applied at the key moments



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
