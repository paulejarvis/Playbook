package main

import (
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
	api "playbook/apis/procstore"
)

var toMarshall = api.Process{
	Name:        "process crush rocks",
	Description: "crushing rocks",
	StartsWith: &api.Step{
		ExecutedBy: &api.Object{
			Name: "rock crusher",
			InstantiatedBy: &api.Object_Employee_{
				Employee: &api.Object_Employee{
					RoleName: "rock crusher",
				},
			},
		},

		Name:        "get hammer",
		Description: "go get the hammer from the hammer room",
		Causes: []*api.State{
			{
				Subject: &api.Object{
					Name: "rock crusher",
					InstantiatedBy: &api.Object_Employee_{
						Employee: &api.Object_Employee{
							RoleName: "rock crusher",
						},
					}},
				Predicate: "has",
				Object: &api.State_StateObject{
					StateObject: &api.Object{
						Name:        "hammer",
						Description: "a rock crushing hammer, very heavy",
						InstantiatedBy: &api.Object_Tool_{
							Tool: &api.Object_Tool{},
						},
					},
				},
			},
		},

		FollowedBy: &api.Step_FollowingStep{
			FollowingStep: &api.Step{
				Name:        "crush rocks",
				Description: "start crushing rocks",
				Requires: []*api.State{
					{
						Subject: &api.Object{
							Name: "rock crusher",
							InstantiatedBy: &api.Object_Employee_{
								Employee: &api.Object_Employee{
									RoleName: "rock crusher",
								},
							},
						},
						Predicate: "has",
						Object: &api.State_StateObject{
							StateObject: &api.Object{
								Name:        "hammer",
								Description: "A rock crushing hammer, very heavy",
								InstantiatedBy: &api.Object_Tool_{
									Tool: &api.Object_Tool{},
								},
							},
						},
					},
				},
				Causes: []*api.State{
					{
						Subject: &api.Object{
							Name:        "rocks",
							Description: "the rocks",
							InstantiatedBy: &api.Object_Thing_{
								Thing: &api.Object_Thing{
									Props: map[string]string{
										"kind": "granite",
									},
								},
							},
						},
						Predicate: "are",
						Object:    &api.State_StateName{StateName: "crushed"},
					},
				},
			},
		},
	},
}

func main() {
	m := jsonpb.Marshaler{
		EnumsAsInts:  false,
		Indent:       " ",
	}

	if err := m.Marshal(os.Stdout, &toMarshall); err != nil {
		log.Fatal("failed to marshall out")
	}
}
