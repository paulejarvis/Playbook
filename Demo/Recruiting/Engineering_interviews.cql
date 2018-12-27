CREATE
  (
    o4:Process
     {
       state: "Input / Output",
       description: "Candidate identified",
//       policy: "Finance",
//       time: 0
      department: "Recruiting",
      workflow: "Engineering interviews"
     }
  ),

  (
    o5:Process
     {
       state: "Action",
       description: "Phone screen",
//       policy: "Finance",
       time: 0.5,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o6:Process
     {
       state: "Action",
       description: "Onsite interview",
//       policy: "Finance",
       time: 3,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o7:Process
     {
       state: "Decision",
       description: "Decide whether to extend offer",
//       policy: "Finance",
//       time: 0
      department: "Recruiting",
      workflow: "Engineering interviews"
     }
  ),

  (
    o8:Process
     {
       state: "Action",
       description: "Create offer",
//       policy: "Finance",
       time: 0.25,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o9:Process
     {
       state: "Action",
       description: "Send conditional offer to candidate",
//       policy: "Finance",
       time: 0.25,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o10:Process
     {
       state: "Action",
       description: "Follow-up call to understand why candidate declined offer",
//       policy: "Finance",
       time: 0.5,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o11:Process
     {
       state: "Action",
       description: "Close candidate",
//       policy: "Finance",
       time: 0.5,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o12:Process
     {
       state: "Action",
       description: "Perform background check",
       policy: "Internal policy",
       time: 0.5,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  ),

  (
    o13:Process
     {
       state: "Input / Output",
       description: "Candidate hired",
//       policy: "Finance",
//       time: 0
      department: "Recruiting",
      workflow: "Engineering interviews"
     }
  ),

  (
    o14:Process
     {
       state: "Decision",
       description: "Close hiring requirement?",
//       policy: "Finance",
//       time: 0
      department: "Recruiting",
      workflow: "Engineering interviews"
     }
  ),

  (
    o15:Process
     {
       state: "Input / Output",
       description: "Remove hiring requirement from ATS and external careers page",
//       policy: "Finance",
       time: 0.25,
       department: "Recruiting",
       workflow: "Engineering interviews"
     }
  )

//Relationships

CREATE

  (o4)-[:THEN]->(o5),
  (o5)-[:THEN]->(o6),
  (o6)-[:THEN]->(o7),
  (o7)-[:THEN {decision:'Offer approved'}]->(o8),
  (o8)-[:THEN]->(o9),
  (o9)-[:THEN {external: 'Candidate accepts offer'}]->(o12),
  (o9)-[:THEN {external: 'Candidate declines offer'}]->(o10),
  (o12)-[:THEN]->(o13),
  (o13)-[:THEN]->(o14),
  (o14)-[:YES {decision:'No new hires needed'}]->(o15),
  (o14)-[:NO {decision: 'Continue interviews'}]->(o4),
  (o5)-[:EXIT {exit: 'Candidate rejected'}]->(o11), //PS reject
  (o6)-[:EXIT {exit: 'Candidate rejected'}]->(o11), //OS reject
  (o7)-[:EXIT {exit: 'Candidate rejected'}]->(o11), //RT reject
  (o12)-[:EXIT {exit: 'Candidate rejected'}]->(o11) //background check fail reject
