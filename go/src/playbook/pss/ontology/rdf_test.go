package ontology

import (
	"fmt"
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
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(propTag, string(propClass), "")), Obj: mustLiteral(string(classProcess))},
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(propTag, string(propLabel), "")), Obj: mustLiteral("Label")},
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(propTag, string(propDescription), "")), Obj: mustLiteral("Description")},
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predOwns), "")), Obj: mustIRI("dl:id/STEP1")},
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predOwns), "")), Obj: mustIRI("dl:id/STEP2")},
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predNext), "")), Obj: mustIRI("dl:id/STEP1")},
				{Subj: mustIRI("dl:id/PROCESS1"), Pred: mustIRI(toiriString(predTag, string(predCanBeDoneBy),"")), Obj: mustIRI("dl:id/ROLE1")},
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
				{Subj: mustIRI("dl:id/EFFECT1"), Pred: mustIRI(toiriString(propTag, string(propClass), "")), Obj: mustLiteral(string(classEffect))},
				{Subj: mustIRI("dl:id/EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predRequiredBy), "")), Obj: mustIRI("dl:id/STEP1")},
				{Subj: mustIRI("dl:id/EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predCausedBy), "")), Obj: mustIRI("dl:id/STEP2")},
				{Subj: mustIRI("dl:id/EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predHasSubject), "")), Obj: mustIRI("dl:id/THING1")},
				{Subj: mustIRI("dl:id/EFFECT1"), Pred: mustIRI(toiriString(predTag, string(predState), "in")), Obj: mustIRI("dl:id/STATE1")},
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


func TestBindTermToField(t *testing.T) {
	cases := []struct{
		n Resource
		pred rdf.Term
		obj rdf.Term
		e Resource
		getter func(resource Resource) interface{}
	}{
		{n: &Process{},
			pred: mustIRI(toiriString(propTag, string(propDescription), "")),
			obj: mustLiteral("Description"),
			e: &Process{Description:"Description"},
			getter: func(r Resource) interface{} {
				return r.(*Process).Description
			},
		},
		{n: &Step{},
			pred: mustIRI(toiriString(predTag, string(predNext), "")),
			obj: mustIRI("dl:id/foooo"),
			e: &Step{NextSteps:[]*Step{{Id:"foooo"}}},
			getter: func(r Resource) interface{} {
				return r.(*Step).NextSteps[0].Id
			},
		},
		{n: &Step{},
			pred: mustIRI(toiriString(predMapTag, string(predNext), "yes")),
			obj: mustIRI("dl:id/foooo"),
			e: &Step{Decisions:map[string]*Step{"yes":{Id:"foooo"}}},
			getter: func(r Resource) interface{} {
				fmt.Println(r)
				return r.(*Step).Decisions["yes"].Id
			},
		},
	}

	for i, c := range cases {
		if err := BindTermToField(c.n, c.pred, c.obj); err != nil {
			t.Fatal(err)
		}
		testutils.Equal(t,c.getter(c.e), c.getter(c.n), "mismatch...%d", i)
	}

}
