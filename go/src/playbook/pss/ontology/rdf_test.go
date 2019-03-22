package ontology

import (
	"github.com/knakk/rdf"
	"playbook/testutils"
	"testing"
)

func TestToRDF(t *testing.T) {
	cases := []struct {
		n node
		e []rdf.Triple
	}{
		{n: &Process{
			Id:          uid("PROCESS1"),
			Label:       "Label",
			Description: "Description",
			Owns: []*Step{
				{Id: "STEP1"},
				{Id: "STEP2"},
			},
			Start:       &Step{Id: "STEP1"},
			CanBeDoneBy: []*Role{{Id: "ROLE1"}}},
			e: []rdf.Triple{
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(propClass)), Obj: mustLiteral(string(classProcess))},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(propLabel)), Obj: mustLiteral("Label")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(propDescription)), Obj: mustLiteral("Description")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(predOwns)), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(predOwns)), Obj: mustIRI("id:STEP2")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(predNext)), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:PROCESS1"), Pred: mustIRI(string(predCanBeDoneBy)), Obj: mustIRI("id:ROLE1")},
			}},

		{n: &Effect{
			Id:          uid("EFFECT1"),
			RequiredBy: []*Step{{Id:"STEP1"}},
			CausedBy: []*Step{{Id:"STEP2"}},
			HasSubject: &Thing{Id:"THING1"},
			StateThing: map[string]stateThing{
				"in": &State{Id:"STATE1"}},
		},
			e: []rdf.Triple{
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(propClass)), Obj: mustLiteral(string(classEffect))},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(predRequiredBy)), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(predCausedBy)), Obj: mustIRI("id:STEP2")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(predHasSubject)), Obj: mustIRI("id:THING1")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI("pred:in"), Obj: mustIRI("id:STATE1")},
			}},
		{n: &Media{
			Id:          uid("MEDIA1"),
			Label: "media",
			Url: "123.com",
			Previous: &Media{Id:"MEDIA1"},
		},
			e: []rdf.Triple{
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(propClass)), Obj: mustLiteral(string(classEffect))},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(predRequiredBy)), Obj: mustIRI("id:STEP1")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(predCausedBy)), Obj: mustIRI("id:STEP2")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI(string(predHasSubject)), Obj: mustIRI("id:THING1")},
				{Subj: mustIRI("id:EFFECT1"), Pred: mustIRI("pred:in"), Obj: mustIRI("id:STATE1")},
			}},
	}

	for i, c := range cases {
		actual, err := ToRDF(c.n)
		if err != nil {
			t.Fatal(err)
		}

		for j, a := range actual {
			e := c.e[j]
			t.Logf("checking %s", e.Serialize(rdf.NTriples))
			testutils.Equal(t, e.Subj, a.Subj, "subjects are not equal %d.%d", i, j)
			testutils.Equal(t, e.Pred, a.Pred, "predicates are not equal  %d.%d", i, j)
			testutils.Equal(t, e.Obj, a.Obj, "objects are not equal  %d.%d", i, j)
		}
	}
}