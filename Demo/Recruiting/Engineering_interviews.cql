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
       description: "Reject candidate",
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
