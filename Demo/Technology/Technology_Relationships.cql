//open_hiring_requirements

CREATE

  (o1)-[:CREATE {description:'Create new hiring requirement in ATS'}]->(t1), //create a new req in ATS
  (o3)-[:CREATE {description:'Create external hiring posting'}]->(t2) //create a new req on external page

//engineering_interviews

CREATE

  (o4)-[:CREATE {description:'Create candidate record in ATS'}]->(t1),
  (o5)-[:UPLOAD {description:'Upload feedback to ATS'}]->(t1), //upload feedback to ATS
  (o6)-[:UPLOAD {description:'Upload feedback to ATS'}]->(t1), //upload feedback to ATS
  (o8)-[:CREATE {description:'Create offer package in compensation tool'}]->(t3), //create offer in comp tool
  (o11)-[:DELETE {description:'Delete candidate record in ATS'}]->(t1), //delete candidate record in ATS
  (o12)-[:CREATE {description:'Create background check'}]->(t4), //create background check in Talentwise
  (o13)-[:UPDATE {description:'Update candidate to HIRED in ATS'}]->(t4), //Update candidate profile to "hired" in ATS
  (o15)-[:DELETE {description:'Delete hiring requirement in ATS'}]->(t1),
  (o15)-[:DELETE {description:'Delete external job posting'}]->(t2)

//new_hire_on-boarding

CREATE

  (o16)-[:UPLOAD {description:'Upload signed policies to ERP'}]->(t5),
  (o17)-[:UPLOAD {description:'Upload bank details to payroll system'}]->(t7)

//asset_provisioning

CREATE

  (o25)-[:UPLOAD {description:'Upload laptop details to asset tracking system'}]->(t6)

//payroll

CREATE

  (o26)-[:CREATE {description:'Create new entry in payroll system'}]->(t7)
