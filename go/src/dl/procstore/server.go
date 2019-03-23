package procstore

import (
	"context"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/pkg/errors"
	"log"
	ps "dl/apis/procstore"
	"dl/dbtranslation"
)

func NewProcStoreServer(dgraph *dgo.Dgraph) ps.ProcStoreServer {
	return &procStoreServerImpl{
		dGo: dgraph,
		translator: dbtranslation.NewTranslator(),
	}
}

type procStoreServerImpl struct {
	dGo *dgo.Dgraph
	translator dbtranslation.Translator
}

func (svr *procStoreServerImpl) CreateProcess(ctx context.Context, process *ps.Process) (*ps.Assigned, error) {
	txn := svr.dGo.NewTxn()
	defer txn.Discard(ctx)

	pb, err := svr.translator.ProcessToDbJson(process)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall process")
	}

	mu := &api.Mutation{
		SetJson:pb,
	}

	log.Println(string(pb))

	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		return nil, errors.Wrap(err, "failed to run mutation")
	}

	return &ps.Assigned{Uids:assigned.Uids}, nil
}