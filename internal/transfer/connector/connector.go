package connector

import "gateway/internal/transfer"

type (
	connector struct {
		transfer transfer.Transfer
	}
)

func NewConnector() *connector {
	return &connector{}
}
