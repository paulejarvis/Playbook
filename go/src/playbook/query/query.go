// Package query is the query producer for graphql+-
// It is extremely immature, and needs work
package query

//func FindEntitiesOfTypeWithName(ctx context.Context, dgraph *dgo.Dgraph, t ps.NodeType, name string) ([]string, error) {
//	// This whole unmarshalling thing is kinda gross at the moment, need
//	// to find something a lot more elegant.
//	query := fmt.Sprintf(`{
//		Query(func: eq(node_type, %d))
//		@filter(eq(name, "%s")) {
//			uid
//		}
//	}`, t, name)
//
//	txn := dgraph.NewTxn()
//	defer txn.Discard(ctx)
//
//	var found struct{Query []struct{Uid string `json:"uid"`}}
//
//	resp, err := txn.Query(ctx, query);
//	if err != nil {
//		return nil, errors.Wrapf(err, "failed to query")
//	}
//
//	if err := json.Unmarshal(resp.Json, &found); err != nil {
//		return nil, errors.Wrap(err, "failed to unmarshall")
//	}
//
//	uuids := make([]string, len(found.Query))
//	for i := range uuids {
//		uuids[i] = found.Query[i].Uid
//	}
//
//	return uuids, nil
//
//}
//
//
//func FindEntityStates(ctx context.Context, dgraph *dgo.Dgraph, es *ps.EntityState) ([]string, error) {
//	query := fmt.Sprintf(`{
//		query(func: eq(node_type, 3))
//		@filter(eq(entity_type, "%s") AND eq(state_name, "%s") AND eq(verb, %d)) {
//    		uid
//  		}
//    }`,es.EntityType, es.StateName, es.Verb)
//
//	txn := dgraph.NewTxn()
//	defer txn.Discard(ctx)
//
//	var found struct{Query []struct{Uid string `json:"uid"`}}
//
//	resp, err := txn.Query(ctx, query);
//	if err != nil {
//		return nil, errors.Wrapf(err, "failed to query")
//	}
//
//	if err := json.Unmarshal(resp.Json, &found); err != nil {
//		return nil, errors.Wrap(err, "failed to unmarshall")
//	}
//
//	uuids := make([]string, len(found.Query))
//	for i := range uuids {
//		uuids[i] = found.Query[i].Uid
//	}
//
//	return uuids, nil
//}