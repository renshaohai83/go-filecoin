package api

import (
	"context"
	"github.com/filecoin-project/go-filecoin/protocol/storage/deal"
	"io"

	uio "gx/ipfs/QmQXze9tG878pa4Euya4rrDpyTNX3kQe4dhCaBzBozGgpe/go-unixfs/io"
	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"

	"github.com/filecoin-project/go-filecoin/actor/builtin/paymentbroker"
	"github.com/filecoin-project/go-filecoin/address"
	"github.com/filecoin-project/go-filecoin/types"
)

// Ask is a result of querying for an ask, it may contain an error
type Ask struct {
	Miner  address.Address
	Price  *types.AttoFIL
	Expiry *types.BlockHeight
	ID     uint64

	Error error
}

// Client is the interface that defines methods to manage client operations.s
type Client interface {
	Cat(ctx context.Context, c cid.Cid) (uio.DagReader, error)
	ImportData(ctx context.Context, data io.Reader) (ipld.Node, error)
	ProposeStorageDeal(ctx context.Context, data cid.Cid, miner address.Address, ask uint64, duration uint64, allowDuplicates bool) (*deal.Response, error)
	QueryStorageDeal(ctx context.Context, prop cid.Cid) (*deal.Response, error)
	ListAsks(ctx context.Context) (<-chan Ask, error)
	Payments(ctx context.Context, dealCid cid.Cid) ([]*paymentbroker.PaymentVoucher, error)
}
