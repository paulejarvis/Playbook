## Demo

We have developed a demo instance that demonstrates the value of Playbook's data-driven approach. The instance has integrated the following processes:
* Opening a new hiring requirement
* Engineering interview process
* New hire on-boarding
* Payroll set-up
* Laptop set-up

The [sample queries and results](https://github.com/paulejarvis/Playbook/blob/master/Demo/Queries%20and%20Unit%20Tests/Sample%20Queries%20and%20Screenshots.md) evidence how Playbook enables organizations to understand and manage their compliance needs and overall operations.

## Understanding risk and surfacing areas for further review

Playbook modules include expert guidance and pre-built tools (operations templates, controls, and queries) to help an organization understand their current risk / compliance footprint.

* *Operations templates* can be implemented to set-up a new process. The major steps and roles have already been encoded so an organization simply needs to customize the details specific to their organization. This can be used for quickly achieving a compliant operational design.
* *Controls* in Playbook are modeled on unit tests from software development methodology and return a TRUE / FALSE as to whether the current operational design satisfies the compliance requirement. An example might be confirming that there is a secondary review step performed by a manager whenever a new entry is created in the payroll system.
* *Queries* provide views of the data and are designed to surface risk or areas for further attention. An example might be showing all internal policies (such as non-discrimination or appropriate use of IT systems) and when these policies are being enforced within workflows.

## Addressing control gaps and reviewing & exporting results

Once an organization has used the tools and guidance in the modules to surface any control gaps or risks, they are ready to remediate the risks or share evidence with regulators or internal governing bodies (such as the Board of Directors or other leadership).

Playbook uses a version-control workflow to manage changes. For example, if a review of SOX 404 controls surfaces a gap in the Payroll processes, the Payroll specialist could "check-out" a new branch of the code (similar to a git workflow where a separate copy is created), make any necessary changes, verify that the new process is compliant by re-running the unit tests, and then submit the new process for acceptance by the Payroll Manager or CFO (who would then approve the "pull request" to merge the new branch into the master codebase)

There are several mechanisms to review operations:
* Within Playbook's "Review" interface
* Export to flowchart such as Visio or Graphviz
* Artifacts such as a risk register
