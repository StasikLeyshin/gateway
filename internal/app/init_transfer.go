package app

import (
	transferInter "gateway/internal/transfer"
	"gateway/internal/transfer/transfer"
)

func (s *serviceProvider) Transfer() transferInter.Transfer {
	if s.transfer == nil {
		s.transfer = transfer.NewTransfer()
	}

	return s.transfer
}
