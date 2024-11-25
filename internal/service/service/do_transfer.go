package service

import "context"

type TransferFunc[Request, Response any] func(ctx context.Context, request Request, client any) (Response, error)

func DoTransfer[
	Request any,
	Response any,
	TransferRequest any,
	TransferResponse any,
](
	ctx context.Context,
	request Request,
	service *Service,
	transfer TransferFunc[Request, Response],
) (Response, error) {
	client := service.Connector.GetClientFromServerType(ctx)

	resp, err := transfer(ctx, request, client)

	return resp, err
}
