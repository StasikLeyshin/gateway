package app

import (
	"gateway/internal/transfer/transfer"
)

func (s *serviceProvider) Transfer() *transfer.Transfer {
	if s.transfer == nil {
		s.transfer = transfer.NewTransfer()
	}

	return s.transfer
}
