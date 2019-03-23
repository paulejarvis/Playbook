package ontology

import (
	"github.com/knakk/rdf"
	"playbook/testutils"
	"testing"
)

func TestToRDF(t *testing.T) {
	cases := []struct {
		n Resource
		e []rdf.Triple
	}{
		{n: &Process{
			Id:          "PROCESS1",
			Label:       "Label",
			Description: "Description",
			Owns: []*Step{
				{Id: "STEP1"},
				{Id: "STEP2"},
			},
			Start:       &Step{Id: "STEP1"},
			CanBeDoneBy: []*Role{{Id: "ROLE1"}}},
			e: []rdf.Triple{
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(classTag)), Obj: mustLiteral(string(classProcess))},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(propTag, string(propLabel))), Obj: mustLiteral("Label")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(propTag, string(propDescription))), Obj: mustLiteral("Description")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predOwns))), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predOwns))), Obj: mustIRI("id:STEP2")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predNext))), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predCanBeDoneBy))), Obj: mustIRI("id:ROLE1")},
			}},

		{n: &Effect{
			Id:         "EFFECT1",
			RequiredBy: []*Step{{Id: "STEP1"}},
			CausedBy:   []*Step{{Id: "STEP2"}},
			HasSubject: &Thing{Id: "THING1"},
			StateThing: map[string]stateThing{
				"in": &State{Id: "STATE1"}},
		},
			e: []rdf.Triple{
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(toiriString(classTag)), Obj: mustLiteral(string(classEffect))},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predRequiredBy))), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predCausedBy))), Obj: mustIRI("id:STEP2")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predHasSubject))), Obj: mustIRI("id:THING1")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predState), "in")), Obj: mustIRI("id:STATE1")},
			}},
	}

	for i, c := range cases {
		actual, err := ToRDF(c.n)
		if err != nil {
			t.Fatal(err)
		}

		for j, a := range actual {
			e := c.e[j]
			t.Logf("checking \n%s vs \n%s", e.Serialize(rdf.NTriples), a.Serialize(rdf.NTriples))
			testutils.Equal(t, e.Subj, a.Subj, "subjects are not equal %d.%d", i, j)
			testutils.Equal(t, e.Pred, a.Pred, "predicates are not equal  %d.%d", i, j)
			testutils.Equal(t, e.Obj, a.Obj, "objects are not equal  %d.%d", i, j)
		}
	}
}
