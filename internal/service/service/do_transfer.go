package service

import "context"

type TransferFunc[Request, Response any] func(ctx context.Context, request Request, client any) (Response, error)

func DoTransfer[
	Request any,
	Response any,
](
	ctx context.Context,
	client any,
	service *Service,
	request Request,
	transfer TransferFunc[Request, Response],
) (Response, error) {
	resp, err := transfer(ctx, request, client)

	return resp, err
}
