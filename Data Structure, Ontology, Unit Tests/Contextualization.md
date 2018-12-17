# What is "Contextualization"

**Contextualization** refers to connecting the processes captured in Playbook to the
people and technologies in-place at the organization. These are often represented by the
organization hierarchy / reporting structure (codified in the directory) and the
network architecture / systems map.

## Why is Contextualization important?

Contextualization ensures the process workflows don't exist in a vacuum - they are connected
to the actual systems and people at the organization. As an example, if a Process node titled
"review payroll journal for current period" is joined to a People node ("Payroll Specialist") via the
relationship "Reviewer" (meaning that the Payroll Specialist is the Reviewer of the process),
Playbook searches for people in the Directory with the Title or Role "Payroll Specialist"
to identify who would actually perform this function.

Not all people or technologies need to be contextualized - in some cases a placeholder or "dummy"
object is used to represent all instances of this object. An example would be using a "Candidate"
object in all hiring or interviewing workflows without contextualizing the candidate by tying the
"Candidate" object to an actual person. This also allows for capturing workflows without
infringing on PII regulations. 
