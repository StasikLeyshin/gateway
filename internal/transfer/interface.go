package transfer

import "context"

type fromTransferToService[T, R any] interface {
	FromTransfer(value T) R
}

type toTransferFromService[T, R any] interface {
	ToTransfer(value T) R
}

func CallService[
	GrpcRequest any,
	GrpcResponse any,
	ServiceRequest toServiceFromApi[GrpcRequest, ServiceRequest],
	ServiceResponse fromServiceToApi[GrpcResponse, GrpcResponse],
](
	ctx context.Context,
	fn func(context.Context, ServiceRequest) (ServiceResponse, error),
	resp GrpcResponse,
	re ServiceRequest,
) (GrpcResponse, error) {
	request :=


	serviceRequest := re.ToService(request)

	serviceResponse, err := fn(ctx, serviceRequest)

	if err != nil {
		return *new(GrpcResponse), nil //, HandleAppError(err)
	}

	resp = serviceResponse.FromService(resp)

	return resp, nil
}
