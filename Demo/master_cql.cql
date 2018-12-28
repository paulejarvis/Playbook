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
             description: "Issue corporate card",
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

//END OF PROCESS
//START OF PEOPLE

CREATE

//people

  (
    p1:People
     {
//     name: "real name",
       title: "Senior VP",
       role: "Hiring Manager",
       team: "Engineering"
     }
  ),

  (
    p2:People
     {
//     name: "real name",
       title: "Recruiter",
       role: "Recruiter",
       team: "Recruiting"
     }
  ),

  (
    p4:People
     {
//       name: "real name",
//       title: "",
       role: "Interview Team",
       team: "Engineering"
     }
  ),

  (
    p5:People
     {
//       name: "real name",
       title: "HR Specialist",
       role: "HR Specialist",
       team: "HR"
     }
  ),

  (
    p6:People
     {
//       name: "real name",
//       title: "text",
       role: "New Hire"
//       team: "text"
     }
  ),

  (
    p7:People
     {
//     name: "real name",
       title: "Director of Business Development",
       role: "Business RP",
       team: "Business Development"
     }
  ),

  (
    p8:People
     {
//     name: "real name",
       title: "Corporate Card Specialist",
       role: "Corporate Card Specialist",
       team: "Finance"
     }
  ),

  (
    p9:People
     {
//       name: "real name",
//       title: "",
       role: "Manager"
//       team: "text"
     }
  ),

  (
    p10:People
     {
//       name: "real name",
       title: "IT Analyst",
       role: "IT Analyst",
       team: "IT"
     }
  ),

  (
    p11:People
     {
//       name: "real name",
       title: "Payroll Specialist",
       role: "Payroll Specialist",
       team: "Finance"
     }
  ),

  (
    p12:People
     {
//       name: "real name",
       title: "Payroll Manager",
       role: "Payroll Manager",
       team: "Finance"
     }
  )

//PEOPLE --> PROCESS relationships

//open_hiring_requirements

CREATE

  (p2)-[:RESPONSIBLE]->(o1),
  (p2)-[:RESPONSIBLE]->(o2),
  (p2)-[:RESPONSIBLE]->(o3),
  (p1)-[:ACCOUNTABLE]->(o1)

//engineering_interviews

CREATE

    (p2)-[:RESPONSIBLE]->(o4), //recruiter responsible for candidate identification
    (p2)-[:RESPONSIBLE]->(o5), //recruiter responsible for PS
    (p4)-[:RESPONSIBLE]->(o6), //interview team responsible for OS
    (p4)-[:RESPONSIBLE]->(o7), //interview team responsible for offer decision
    (p1)-[:ACCOUNTABLE]->(o7), //HM accountable for offer decision
    (p2)-[:CONSULTED]->(o7), //recruiter CONSULTED for offer decision
    (p1)-[:ACCOUNTABLE]->(o8), //HM accountable for offer creation
    (p2)-[:RESPONSIBLE]->(o7), //recruiter RESPONSIBLE for offer creation
    (p2)-[:RESPONSIBLE]->(o9), //recruiter RESPONSIBLE for sending offer
    (p1)-[:RESPONSIBLE]->(o10), //HM RESPONSIBLE for FLUP call if offer declined
    (p2)-[:RESPONSIBLE]->(o11), //recruiter RESPONSIBLE closing candidate
    (p2)-[:RESPONSIBLE]->(o12), //recruiter RESPONSIBLE for background check
    (p1)-[:ACCOUNTABLE]->(o14), //HM accountable for decision on keeping hiring open
    (p2)-[:CONSULTED]->(o14), //recruiter consulted on keeping hiring open
    (p2)-[:RESPONSIBLE]->(o15) //recruiter responsible for taking down hiring req

//new_hire_on-boarding

CREATE

  (p6)-[:REPORTS_TO]->(p9), //new hire reports to manager
  (p6)-[:RESPONSIBLE]->(o16), //new hire responsible for signing policies
  (p5)-[:ACCOUNTABLE]->(o16), //HR ACCOUNTABLE for signing policies
  (p5)-[:RESPONSIBLE]->(o17), //new hire responsible for providing bank details
  (p11)-[:CONSULTED]->(o17), //payroll specialist is consulted on bank details
  (p6)-[:PARTICIPANT]->(o18), //new hire PARTICIPANT in biz presentation
  (p7)-[:ACCOUNTABLE]->(o18), //director of BD accountable for presentation
  (p9)-[:RESPONSIBLE]->(o19), //manager responsible for decision on corp card
  (p6)-[:CONSULTED]->(o19), //new hire CONSULTED for decision on corp card
  (p8)-[:RESPONSIBLE]->(o20) //payroll RESPONSIBLE for issuing corp card

//asset_provisioning

CREATE

  (p6)-[:RESPONSIBLE]->(o22),
  (p10)-[:RESPONSIBLE]->(o23),
  (p10)-[:RESPONSIBLE]->(o24),
  (p10)-[:RESPONSIBLE]->(o25)

//payroll

CREATE

  (p11)-[:RESPONSIBLE]->(o26), //payroll specialist is RESPONSIBLE for creating new payroll entry
  (p12)-[:REVIEWER]->(o26), //payroll manager REVIEWS creation of new payroll entity
  (p11)-[:RESPONSIBLE]->(o27) //payroll specialist is RESPONSIBLE for $1 test of direct deposity

//END OF People
//START OF TECHNOLOGIES

CREATE

//people

  (
    t1:Technology
     {
       application: "Lever",
       function: "Applicant tracking system",
       provider: "Lever"
     }
  ),

  (
    t2:Technology
     {
       application: "External Careers Page",
       function: "External Careers Page",
       provider: "Internal"
     }
  ),

  (
    t3:Technology
     {
       application: "Compensation Tool",
       function: "Compensation details and guidance for employees",
       provider: "Internal"
     }
  ),

  (
    t4:Technology
     {
       application: "Talentwise",
       function: "Background check",
       provider: "Talentwise"
     }
  ),

  (
    t5:Technology
     {
       application: "Workday",
       function: "ERP",
       provider: "Workday"
     }
  ),

  (
    t6:Technology
     {
       application: "AssetExplorer",
       function: "IT asset tracking",
       provider: "ManageEngine"
     }
  ),

  (
    t7:Technology
     {
       application: "Workforce Now",
       function: "Payroll system",
       provider: "ADP"
     }
  )

//TECH RELATIONSHIPS

//open_hiring_requirements

CREATE

  (o1)-[:CREATE {description:'Create new hiring requirement in ATS'}]->(t1), //create a new req in ATS
  (o3)-[:CREATE {description:'Create external hiring posting'}]->(t2) //create a new req on external page

//engineering_interviews

CREATE

  (o4)-[:CREATE {description:'Create candidate record in ATS'}]->(t1),
  (o5)-[:UPLOAD {description:'Upload feedback to ATS'}]->(t1), //upload feedback to ATS
  (o6)-[:UPLOAD {description:'Upload feedback to ATS'}]->(t1), //upload feedback to ATS
  (o8)-[:CREATE {description:'Create offer package in compensation tool'}]->(t3), //create offer in comp tool
  (o11)-[:DELETE {description:'Delete candidate record in ATS'}]->(t1), //delete candidate record in ATS
  (o12)-[:CREATE {description:'Create background check'}]->(t4), //create background check in Talentwise
  (o13)-[:UPDATE {description:'Update candidate to HIRED in ATS'}]->(t4), //Update candidate profile to "hired" in ATS
  (o15)-[:DELETE {description:'Delete hiring requirement in ATS'}]->(t1),
  (o15)-[:DELETE {description:'Delete external job posting'}]->(t2)

//new_hire_on-boarding

CREATE

  (o16)-[:UPLOAD {description:'Upload signed policies to ERP'}]->(t5),
  (o17)-[:UPLOAD {description:'Upload bank details to payroll system'}]->(t7)

//asset_provisioning

CREATE

  (o25)-[:UPLOAD {description:'Upload laptop details to asset tracking system'}]->(t6)

//payroll

CREATE

  (o26)-[:CREATE {description:'Create new entry in payroll system'}]->(t7)
