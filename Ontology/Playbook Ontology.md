# Playbook Ontology

The Playbook Ontology is a framework to describe business processes and operations.
It draws from industry-standard conventions and best practices such as UML,
Project Management Institute, and others, while remaining flexible to accommodate user-defined variance.

## Objects, Properties, and Relationships

The Playbook Ontology is composed of **Objects** with **Properties** and the **Relationships**
between Objects.

### Objects

There are 3 object types: *People*, *Process*, and *Technology*.

* People are the individuals within your organization who perform a function related to process
or technology. There are multiple possible roles that a people can play (see [Relationships](#relationships))
* Process refers to the actions in a workflow
* Technology encompasses the systems and tools used to execute the workflow. This
definition extends to non-technical systems such as meetings or phone calls

### Properties

Properties describe objects. Some properties are shared across object types while
others are unique to specific object types.

* People
  * Name (Label)
  * Title
  * Team
  * Qualifications (CPA, PMP, etc.)
  * Location
  * Email
  * Phone Number
  * Start Date
  * Employee ID (Primary Key)
* Process
  * Description (Label)
  * Type
    * Start / Input
    * End / Output
    * Action
    * Decision
  * Workflow
  * Stage
  * Process_ID (derived from Workflow + Stage) (Primary Key)
  * Control_Framework (Sarbanex-Oxley, Revenue 606, GDPR, etc.)
  * Time Required
  * Notes
* Technology
  * Application Name (Label)
  * Provider
  * Version
  * Overview
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
  * Decision Outputs (Boolean)
* Technology <-> Technology
  * Automated Connection (API or script)
  * Manual Connection
