package transaction

import (
	"github.com/benpate/content"
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
)

type NilTransaction datatype.Map

func (txn NilTransaction) Execute(c *content.Content) (int, error) {
	return 0, derp.New(500, "content.NilTransaction", "Unrecognized Transaction Type", txn)
}

func (txn NilTransaction) Description() string {
	return "Nil Transaction"
}
