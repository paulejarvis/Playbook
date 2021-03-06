# Playbook Ontology

The Playbook Ontology is a framework to describe business processes and operations.
It draws from industry-standard conventions and best practices within the fields of Human Resources, Project Management / Operations, and Information Technology, while remaining flexible to accommodate user-defined variance.

![](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Schema.PNG)

## Objects, Properties, and Relationships

The Playbook Ontology is composed of **Objects** with **Properties** and the **Relationships**
between Objects.

### Objects

There are 3 base object types: *People*, *Process*, and *Technology*. A *Workflow* is the end-to-end process steps. As an example, "New hire payroll set-up" would be a workflow and "confirm employee's direct deposit details" would be a process step.

* People are the individuals within your organization who perform a function related to process
or technology. There are multiple possible roles that a people can play (see [Relationships](#relationships))
* Process refers to the actions in a workflow
* Technology encompasses the systems and tools used to execute the process. This
definition extends to non-technical systems such as meetings or phone calls

### Properties

Properties describe objects. Some properties are shared across object types while others are unique to specific object types.

* People
  * Name
  * Company
  * Title (business title )
  * Team
  * Role (this is process-dependent and can differ from Title; Label)
  * Qualifications (CPA, PMP, etc.)
  * Location
  * Email
  * Phone Number
  * Start Date
  * Employee ID (Primary Key)
* Process
  * Description (Label)
  * Process Stage
    * Input / Output (either the end of one process, the start of another process, or both)
    * Action
    * Decision
  * Department (inherited from parent folder)
  * Workflow (inherited from file name)
  * Stage
  * Control_Framework (Sarbanes-Oxley, Revenue 606, GDPR, etc.)
  * Time Required (Actions)
  * Recurring (Y/N)
  * Frequency
  * Start Date
  * Notes
* Technology
  * Application (Label)
  * Provider
  * Version
  * Overview
  * Function (i.e. Applicant Tracking System, CRM, Payroll, etc.)
  * Implementation Date
  * Authentication
  * Hosting

### Relationships

* People <-> Process
  * Responsible (Actually completes the task; can be shared across multiple individuals)
  * Accountable (Answerable for the activity or decision; only one person)
  * Reviewer (Double-checks the work or provides a second set of eyes; can be multiple people)
  * Consulted (Provides input prior to a final decision; can be multiple people)
  * Sign-Off (Final confirmation or approval; often an executive or Board Member)
  * Participant (Attends an event but not responsible for any action)
* People <-> Technology
  * User (this could be expanded to more specific permissions profiles)
  * Workflow Administrator
  * Business Owner
  * System Administrator
* Process <-> Technology
  * Upload / Import or Download / Export
  * Create / Delete entry or record
  * Update
* People <-> People
  * Reports to / Manages
  * Supervises / Supervised by
  * Substituting for / Covered by (for leaves of absence)
* Process <-> Process
  * Then (default connector between "Action" process nodes; assumes that the previous action had completed successfully)
  * Decision Outputs (conditional; only one option can be true)
  * Blocked on / Permits
* Technology <-> Technology
  * Automated Connection (API or script)
  * Manual Connection
