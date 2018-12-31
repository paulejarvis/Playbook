CREATE

//nodes

  (
    o26:Process
     {
       state: "Action",
       description: "Create new entry in payroll system",
       policy: "Sarbanes-Oxley",
       time: 0.5,
       department: "Finance",
       workflow: "Payroll set up"
     }
  ),

  (
    o27:Process
     {
       state: "Action",
       description: "Deposit $1 in new hire's account to confirm direct deposit",
//       policy: "Finance",
       time: 0.25,
       department: "Finance",
       workflow: "Payroll set up"
     }
  )

//relationships

CREATE
  (o26)-[:THEN]->(o27)
