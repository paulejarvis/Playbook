## Unit Tests

Unit Tests are run against the code base to assess if the current process design satisfies active control requirements.

### Example

Payroll journal entries must be reviewed by a different Payroll Specialist than the original Payroll Specialist. In this case, the unit tests checks to see if the action "Review payroll journal entry" is performed by a different People object than the prior action "Prepare payroll journal entry" and returns TRUE or FALSE
