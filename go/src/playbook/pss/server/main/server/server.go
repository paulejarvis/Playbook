package main

import (
	"flag"
	"fmt"
	"github.com/knakk/sparql"
	"log"
	"playbook/pss/ontology"
	"playbook/pss/server"
)

var repoLoc = flag.String("rl", "http://3.17.68.15/neptune/", "repo-location, the address of the rdf repo")

func main() {
	flag.Parse()
	repo, err := sparql.NewRepo(*repoLoc)
	if err != nil {
		log.Fatalf("failed to create repo %s", err)
	}

	n := &ontology.Process{
		Id:          "PROCESS1",
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

	if err := server.ReadResource(repo,n); err != nil {
		log.Fatalf("couldn't read! %s, err")
	}

	fmt.Println("SUCCESS")

}