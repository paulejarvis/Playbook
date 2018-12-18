## Overview

Playbook is built on the neo4j graph database. See the [Playbook Architecture](#https://github.com/paulejarvis/Playbook/blob/master/Tech/Playbook%20Architecture.png)) for more details.

### Inputs

* People
  * Organizational directory (i.e. Active Directory)
* Process
  * Ingest existing process documentation (Visio, Excel, etc)
  * User-defined process input via the UI [See User Input](#/User%Input)
* Technology
  * Systems map (such as nmap)

inputs --> py2neo --> neo4j

from neo4j --> browser / UI; can then be exported to py2neo --> pyviz --> graphviz or queried directly

git backs up all data in neo4j

git statistics for data on most recently updated repos etc (https://lukasmestan.com/git-quick-stats/)
