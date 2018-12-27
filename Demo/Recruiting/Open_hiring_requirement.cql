CREATE
  (
    o1:Process
     {
       state: "Input / Output",
       description: "Open new role in ATS",
       time: 0.25,
       department: "Recruiting",
       workflow: "Open hiring requirement"
     }
  ),

  (
    o2:Process
     {
       state: "Decision",
       description: "Is the organization a Federal contractor?",
       policy: "VEVRAA",
//       time: 0
       department: "Recruiting",
       workflow: "Open hiring requirement"
     }
  ),

  (
    o3:Process
     {
       state: "Action",
       description: "Open job posting on organization's external careers page",
//       policy: "Finance",
       time: 0.25,
       department: "Recruiting",
       workflow: "Open hiring requirement"
     }
  )
