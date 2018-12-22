## Overview

Playbook is built on the neo4j graph database. See the [Playbook Architecture](https://github.com/paulejarvis/Playbook/blob/master/Tech/Playbook%20Architecture.png) for more details.

### Inputs

* People
  * Organizational directory (i.e. Active Directory)
  * https://labs.mwrinfosecurity.com/blog/visualising-organisational-charts-from-active-directory/
* Process
  * Ingest existing process documentation (Visio, Excel, etc)
  * User-defined process input via the UI ([See User Input](https://github.com/paulejarvis/Playbook/tree/master/Tech/User%20Input))
* Technology
  * Systems map (such as nmap)

### Drivers

* Py2Neo for connecting to neo4j
* Pyviz and viz.js for Graphviz visualization
* https://github.com/neo4j/neo4j/tree/2.3/community/graphviz

### Git

* Git is used to back up all the data in neo4j
* [Git Statistics](https://lukasmestan.com/git-quick-stats/) provides metadata on git activity
