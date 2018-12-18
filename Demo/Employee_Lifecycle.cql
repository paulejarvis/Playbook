//Description
//Tracks the full lifecycle of an employee from hiring to exit

//initializing

from py2neo.data import Node, Relationship

//nodes
//people objects

ppl1 = Node("People", Role="Hiring Manager")
ppl2 = Node("People", Role="Candidate")
ppl3 = Node("People", Role="Recruiter")
ppl4 = Node("People", Role="Compensation Manager")
ppl5 = Node("People", Role="Interviewing Team")

//process objects

Processes
1. Open new role
  * Hiring Manager
  * ATS system
2. Confirm job posting on external careers page
  * Recruiter
  * Company Website
  * subject to regulation
3. Candidate enters interview processes
  * Start of new processes
  * Lever
  * Recruiter
4. Interview process - PS, OS,
5. Round table / decision to hire (Y/N)
6. Offer created
7. Candidate accepts; if yes then spawn several other Processes
  * on-boarding
  * new hire IT provisioning
  * payroll (confirm direct deposit )

//technology objects

Technology
* Applicant Tracking System
* Compensation Tool
* Meeting
* External Careers Page


//relationships

1. Hiring manager "CREATES" new role in ATS systems
2. Recruiter "CREATES" new job posting on External Careers Page
3. 

//what unit tests are applicable

is the position posted on the external website before an offer is made
is there secondary review of the comp package by comp manager before hiring manager sends an offer?










//nodes

CREATE (

  s:Start
    {
      Process: "Engineering Interview",
      Label: "Candidate enters interview process"
    }
),

(
  p1:Process
    {
      Process: "Engineering Interview",
      Label: "Recruiter Screen\n 45-minute phone call"
    }
),

(
  p2:Process
    {
      Process: "Engineering Interview",
      Label: "Team Screen\n 60 minute video call"
    }
),

(
  p3:Process
    {
      Process: "Engineering Interview",
      Label: "Candidate Debrief Meeting"
    }
)

//relationships

CREATE
  (s)-[sch:Scheduled]->(p1),
  (p1)-[pass1:Pass]->(p2),
  (p2)-[pass2:Pass]->(p3)



//return

Return
  s,
  p1,
  p2,
  p3,
  sch,
  pass1,
  pass2
