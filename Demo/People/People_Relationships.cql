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

  (p11)-[:RESPONSIBLE]->(o26),
  (p12)-[:REVIEWER]->(o26),
  (p11)-[:RESPONSIBLE]->(o27)
