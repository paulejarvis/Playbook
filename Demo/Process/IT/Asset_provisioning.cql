CREATE

  (
    o22:Process
     {
       state: "Decision",
       description: "Choose laptop model",
//       policy: "Finance",
//       time: 0
      department: "IT",
      workflow: "Asset provisioning"
     }
  ),

  (
    o23:Process
     {
       state: "Action",
       description: "Provision Windows laptop",
//       policy: "Finance",
       time: 1,
       department: "IT",
       workflow: "Asset provisioning"
     }
  ),

  (
    o24:Process
     {
       state: "Action",
       description: "Provision Mac laptop",
//       policy: "Finance",
       time: 1,
       department: "IT",
       workflow: "Asset provisioning"
     }
  ),

  (
    o25:Process
     {
       state: "Action",
       description: "Install security certificate and connect to domain",
//       policy: "Finance",
       time: 0.5,
       department: "IT",
       workflow: "Asset provisioning"
     }
  )

//Relationships

CREATE

  (o22)-[:THEN {decision:'Windows selected'}]->(o23),
  (o22)-[:THEN {decision:'Mac selected'}]->(o24),
  (o23)-[:THEN]->(o25),
  (o24)-[:THEN]->(o25)
