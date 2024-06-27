package txn

import (
	"time"

	"github.com/oreo-dtx-lab/oreo/pkg/config"
)

type ItemOptions struct {
	Key       string
	Value     string
	TxnId     string
	TxnState  config.State
	TValid    time.Time
	TLease    time.Time
	Prev      string
	LinkedLen int
	IsDeleted bool
	Version   string
}

type DataItemFactory interface {
	NewDataItem(ItemOptions) DataItem
}
