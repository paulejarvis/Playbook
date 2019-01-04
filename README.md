# Playbook

## What is the problem?

Organizations lack a single source of truth for business processes. Executives, managers, and end users within the organization, as well as auditors, advisors, consultants, and others outside of the organization, do not have a shared view of how the organization operates.

## Why is this important now?

Businesses are getting more complex, competition is increasing and margins are decreasing, and the global regulatory landscape is becoming more complicated. Organizations that aspire to be competitive require a unified operating manual, or Playbook, to ensure that the they are operating as effectively as safely possible, optimally using resources, keeping overhead and cost to a minimum, complying with applicable laws and regulations, and reducing unnecessary friction and delays.

## Why hasn't this been solved before?

Existing tools focus on visualization - this is an oversimplification of the complexity of modern regulated processes and does not scale. Process documentation, if it exists, is distributed across the organization and stored in various formats (flowcharts, checklists, wikis, etc). The lack of a universal format or interface results in disconnected and siloed process documentation that does not capture how workstreams are interdependent and interconnected and overlooks the reality that outputs from one process may for the inputs to a second process or how resources may be shared across processes. This one-dimensional approach produces documentation that does not accurately capture current state (and is burdensome to produce) and does not help the organization achieve its goals of operating safely and legally, reducing cost and friction, and optimizing resource usage.

## What is Playbook?

Playbook is a platform for integrated business processes. Playbook *treats process like code* and uses a data model that represents business processes as People, Process, and Technology objects with relationships between these objects. The [Playbook Ontology](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Playbook%20Ontology.md) describes the data model in further detail.

## Why is this valuable?

Playbook establishes a new standard and interface for organizations and their external partners to share a single, reliable view of how the organization operates. This is valuable for:

* Analyzing operations across and identifying friction, powering informed decisions on where to invest in resources and technology (i.e. cutting an IT system that is only used in one workflow or shifting hiring to an understaffed team)
* Establishing a platform and shared standard for organizations to review their operations with advisory firms and auditors, reducing the time and cost of understanding their compliance needs and risks
* Updating processes with nuanced version control, eliminating the need for time-consuming and disruptive re-orgs
* Empowering every employee to keep documentation up-to-date and actionable
* Optimizing employee productivity as job requirements and workflow steps are clearly documented for all roles

## Who are the customers?

Playbook seeks to partner with major advisory and audit firms to implement Playbook's data model and interface as the standard in their client engagements. Clients would also have access to pre-built process templates (to implement new processes), queries (to understand current processes), and tests (to surface control gaps and risks).

#~~~~Under Development~~~~#

## How does Playbook work?

### Initial set-up

* Organizations first complete an questionnaire designed to identify their compliance needs. The result of the questionnaire is a list of suggested "modules" which each correspond to a regulatory framework and contain:
  * Expert guidance and analysis
  * Pre-built controls, unit tests, and queries, to assess current compliance / control status ()
  * Operational templates which can be adapted for the organization and used to design compliant operations
* Organizations import their current operations and processes into Playbook and then use the tools contained in the module to understand if they are currently compliant and if not, how they must adapt to reach compliance.
  * This import can be accomplished manually via the Playbook interface, starting from a pre-built template, or through structured import of existing process docs (Visio, Excel, etc.)
  * The initial configuration is typically done by operational resources within the organization (such as the PMO or Internal Audit) or in conjunction with Playbook partners
  * Additional details can be imported automatically from the organization's directory or existing IT infrastructure  

### Understanding risk and surfacing areas for further review

Playbook modules include expert guidance and pre-built tools (operations templates, controls, and queries) to help an organization understand their current risk / compliance footprint.

* *Operations templates* can be implemented to set-up a new process. The major steps and roles have already been encoded so an organization simply needs to customize the details specific to their organization. This can be used for quickly achieving a compliant operational design.
* *Controls* in Playbook are modeled on unit tests from software development methodology and return a TRUE / FALSE as to whether the current operational design satisfies the compliance requirement. An example might be confirming that there is a secondary review step performed by a manager whenever a new entry is created in the payroll system.
* *Queries* provide views of the data and are designed to surface risk or areas for further attention. An example might be showing all internal policies (such as non-discrimination or appropriate use of IT systems) and when these policies are being enforced within workflows.

### Addressing control gaps and reviewing & exporting results

Once an organization has used the tools and guidance in the modules to surface any control gaps or risks, they are ready to remediate the risks or share evidence with regulators or internal governing bodies (such as the Board of Directors or other leadership).

Playbook uses a version-control workflow to manage changes. For example, if a review of SOX 404 controls surfaces a gap in the Payroll processes, the Payroll specialist could "check-out" a new branch of the code (similar to a git workflow where a separate copy is created), make any necessary changes, verify that the new process is compliant by re-running the unit tests, and then submit the new process for acceptance by the Payroll Manager or CFO (who would then approve the "pull request" to merge the new branch into the master codebase)

There are several mechanisms to review operations:
* Within Playbook's "Review" interface
* Export to flowchart such as Visio or Graphviz
* Artifacts such as a risk register

## Playbook pricing

Playbook is priced by module and functionality. Modules vary according to complexity and the level of detail required to produce expert guidance, templates, controls, and queries. There is no license fee for Playbook and modules require an annual subscription and are updated regularly.

Once data is ingested in Playbook it's free for organizations to explore the data or write custom controls or queries.

## Demo

We have developed a demo instance that demonstrates the value of Playbook's data-driven approach. The instance has integrated the following processes:
* Opening a new hiring requirement
* Engineering interview process
* New hire on-boarding
* Payroll set-up
* Laptop set-up

The [sample queries and results](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Sample%20Queries%20and%20Screenshots.md) evidence how Playbook enables organizations to understand and manage their compliance needs and overall operations.

### Disclaimer

While Playbook is developed in partnership with subject-matter experts and the content is updated regularly, Playbook cannot guarantee that an organization will pass an audit / compliance review - that determination can only be made by the appropriate governing body.
