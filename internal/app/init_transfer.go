package app

import (
	transferInter "gateway/internal/transfer"
	"gateway/internal/transfer/connector"
	"gateway/internal/transfer/transfer"
)

func (s *serviceProvider) Transfer() transferInter.Transfer {
	if s.transfer == nil {
		s.transfer = transfer.NewTransfer()
	}

	return s.transfer
}

func (s *serviceProvider) Connector() *connector.Connector {
	if s.connector == nil {
		s.connector = connector.NewConnector(s.Transfer())
	}

	return s.connector
}
