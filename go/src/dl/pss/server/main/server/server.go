package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/knakk/sparql"
	"log"
	"dl/pss/ontology"
	"dl/pss/server"
)

var repoLoc = flag.String("rl", "http://3.17.68.15/neptune/", "repo-location, the address of the rdf repo")

func main() {
	flag.Parse()
	repo, err := sparql.NewRepo(*repoLoc)
	if err != nil {
		log.Fatalf("failed to create repo %s", err)
	}

	n := &ontology.Process{
		Id:          "PROCESS2",
		Label:       "Label",
		Description: "Description",
		Owns: []*ontology.Step{
			{Id: "STEP1"},
			{Id: "STEP2"},
		},
		Start:       &ontology.Step{Id: "STEP1"},
		CanBeDoneBy: []*ontology.Role{{Id: "ROLE1"}}}

	if _, err := server.WriteRDF(repo, n); err != nil {
		log.Fatalf("failed to write to repo %s", err)
	}

	p := &ontology.Process{Id:"PROCESS2"}
	if err := server.ReadResource(repo, p); err != nil {
		log.Fatalf("couldn't read! %s, err")
	}

	b, _ := json.MarshalIndent(p, "", " ")
	fmt.Println(string(b))

	fmt.Println("SUCCESS")

}