package server

import (
	"bytes"
	"fmt"
	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
	"github.com/pkg/errors"
	"playbook/pss/ontology"
)

func WriteRDF(repo *sparql.Repo, r ontology.Resource) (*sparql.Results, error) {
	triples, err := ontology.ToRDF(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert resource to RDF")
	}

	rw := &bytes.Buffer{}
	encoder :=  rdf.NewTripleEncoder(rw, rdf.NTriples)
	if err := encoder.EncodeAll(triples); err != nil {
		return nil, errors.Wrap(err, "failed to encode triples")
	}
	if err := encoder.Close(); err != nil {
		return nil, errors.Wrap(err, "failed to close encoder")
	}


	if err := repo.Update(fmt.Sprintf("INSERT DATA {%s}", rw.String())); err != nil {
		return nil, errors.Wrap(err, "failed to update")
	}

	return nil, nil
}


func ReadResource(repo *sparql.Repo, r ontology.Resource) error {
	// TODO(henry): pagination needs to happen sooner rather than later...
	fmt.Println(r.GetIri())
	res, err := repo.Query(fmt.Sprintf("select ?p ?o where {<%s> ?p ?o} limit 100", r.GetIri()))
	if err != nil {
		return errors.Wrap(err, "failed to query")
	}

	var pairs []ontology.PredObjPair
	sols := res.Solutions()
	for _, pair := range sols {
		pairs = append(pairs, ontology.PredObjPair{P: pair["p"], O: pair["o"] })
	}

	if err := ontology.DecodeTerms(r, pairs); err != nil {
		return errors.Wrap(err, "failed to decode")
	}

	return nil

}