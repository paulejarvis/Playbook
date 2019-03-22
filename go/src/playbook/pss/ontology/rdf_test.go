package ontology

import (
	"bytes"
	"fmt"
	"github.com/knakk/rdf"
	"testing"
)

func TestToRDF(t *testing.T) {
	n := &Process{
		Id: uid("PROCESS1"),
		Label: "Process Label",
		Description: "Process Description",
		Owns: []*Step{
			{Id:"STEP1"},
			{Id:"STEP2"},
		},
		Start: &Step{Id:"STEP1"},
		CanBeDoneBy: []*Role{{Id: "ROLE1"}},
	}

	out, err := ToRDF(n)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(out)

	buf := &bytes.Buffer{}
	enc := rdf.NewTripleEncoder(buf, rdf.NTriples)
	if err := enc.EncodeAll(out); err != nil {
		t.Fatal(err)
	}
	if err := enc.Close(); err != nil {
		t.Fatal(err)
	}

	t.Log(buf.String())
}
