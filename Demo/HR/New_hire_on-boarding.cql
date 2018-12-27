CREATE

  (
    o16:Process
     {
       state: "Action",
       description: "Sign internal policies (insider trading, appropriate use of systems, code of conduct, etc)",
       policy: "Internal policy",
       time: 1,
       department: "HR",
       workflow: "New hire on-boarding"
     }
  ),

  (
    o17:Process
     {
       state: "Input / Output",
       description: "Provide bank account details for payroll direct deposit",
//       policy: "Finance",
       time: 0.25,
       department: "HR",
       workflow: "New hire on-boarding"
     }
  ),

  (
    o18:Process
     {
       state: "Action",
       description: "Attend presentation on business strategy",
//       policy: "Finance",
       time: 1,
       department: "HR",
       workflow: "New hire on-boarding"
     }
  ),

  (
    o19:Process
     {
      state: "Decision",
      description: "Approved for corporate card?",
      policy: "Internal Policy - Corporate credit card policy",
//       time: 0
      department: "HR",
      workflow: "New hire on-boarding"
     }
  ),

  (
    o20:Process
     {
       state: "Action",
       description: "Receive corporate card",
//       policy: "Finance",
       time: 0.5,
       department: "HR",
       workflow: "New hire on-boarding"
     }
  ),

  (
    o21:Process
     {
       state: "Input / Output",
       description: "New hire ready for staffing assignment",
//       policy: "Finance",
//       time: 0.5,
       department: "HR",
       workflow: "New hire on-boarding"
     }
  )

//relationships

CREATE

  (o16)-[:THEN]->(o17),
  (o17)-[:THEN]->(o18),
  (o18)-[:THEN]->(o19),
  (o19)-[:YES]->(o20),
  (o19)-[:NO]->(o21),
  (o20)-[:THEN]->(o21)
