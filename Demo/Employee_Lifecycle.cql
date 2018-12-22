//Description
//Tracks the full lifecycle of an employee from hiring to exit

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


//nodes

CREATE (

  io1:Input-Output
    {
      Workflow: "Engineering Interview",
      Label: "Candidate enters interview process"
    }
),

(
  p1:Process
    {
      Workflow: "Engineering Interview",
      Label: "Recruiter Screen\n 45-minute phone call"
    }
),

(
  p2:Process
    {
      Workflow: "Engineering Interview",
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
//list out nodes to return 
