//Description
//Tracks the full lifecycle of an employee from hiring to exit

//nodes

USE PY2NEO


People
* Hiring Manager
* Candidate
* Recruiter
* Compensation Manager
* Interviewing Team

Processes
1. Open new role
  * Hiring Manager
  * ATS system
2. Confirm job posting on external website
  * Recruiter
  * Company Website
  * subject to regulation
3. Candidate enters interview processes
  * Start of new processes
  * Lever
  * Recruiter
4. Interview process - PS, OS,
5. Round table / decision to hire
6. Offer created
7. Candidate accepts; if yes then spawn several other Processes
  * on-boarding
  * new hire IT provisioning
  * payroll (confirm direct deposit )


Technology
* Applicant Tracking System
* Compensation Tool
* Meeting


//relationships






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
