tech requirements per roles - this informs separation of duties & IT controls

what are the process requirements per role?

What processes is each system involved in?

what regulatory requirements apply to my hiring process?
* search all nodes in recruiting / Engineering, to see where control framework is not null

how much time am I spending on each candidate who makes it to the offer stage?
  whos time is involved? how much time for each role?


are my internal policies being applied at the key moments

can we have a time trigger to revisit proesses every year? should that be a proprety for timing?

how mcuh time am I spending on SOX compliance? Should we hire a consultant to help streamline?



//Show all the processes that are regulated and how much time each step takes and who owns the steps in the workflow
```Cypher
MATCH (People)-[:RESPONSIBLE]->(Process)
WHERE Process.policy IS NOT NULL
RETURN People.role,Process.workflow, Process.description, Process.policy, Process.time
```
