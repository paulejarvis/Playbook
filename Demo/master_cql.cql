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

//relationships

CREATE

  (o1)-[:THEN]->(o2),
  (o2)-[:YES]->(o3)

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

    //joining to the previous process

    CREATE
      (o3)-[:THEN]->(o4),
      (o2)-[:NO]->(o4)

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

//join to previous process

  CREATE

    (o13)-[:THEN]->(o16)

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

//join IT to eng hiring

CREATE

  (o13)-[:THEN]->(o22)

//adding payroll

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

//payroll relationships

CREATE
  (o26)-[:THEN]->(o27)

//join payroll to new hire

CREATE
  (o17)-[:THEN]->(o26)
