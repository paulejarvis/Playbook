# Playbook

## What is Playbook?

Playbook is the new standard for process documentation, providing a shared interface and universal data model for organizations and professional services firms (such as lawyers, auditors, and advisors) to review current processes and operations, iteratively update processes with nuanced version-control, and access expert guidance, pre-built process templates, and queries and tests designed to surface risk and controls gaps.

Playbook partners with subject-matter experts to develop content across major regulatory and compliance frameworks. Playbook's data model and software platform enable a data-driven approach to compliance; reducing risk, cost, and time required.  

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

### What is the Playbook data model?

Playbook models the people, processes, and technologies, powering the organization as objects in a graph database. Objects have properties and are connected to each other via relationships (which also have properties describing the relationship). The [Playbook Ontology](https://github.com/paulejarvis/Playbook/blob/master/Data%20Structure%20and%20Ontology/Playbook%20Ontology.md) describes the data model in further detail.

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

You can connect to the Playbook demo instance using these details:
* URL:
* Username:
* Password:

### Disclaimer

While Playbook is developed in partnership with subject-matter experts and the content is updated regularly, Playbook cannot guarantee that an organization will pass an audit / compliance review - that determination can only be made by the appropriate governing body.
