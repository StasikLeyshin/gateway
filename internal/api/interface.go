package api

import "context"

type fromServiceToApi[T, R any] interface {
	FromService(value T) R
}

type toServiceFromApi[T, R any] interface {
	ToService(value T) R
}

func CallService[
	GrpcRequest any,
	GrpcResponse any,
	ServiceResponse fromServiceToApi[GrpcResponse, GrpcResponse],
	ServiceRequest toServiceFromApi[GrpcRequest, ServiceRequest],
](
	ctx context.Context,
	request GrpcRequest,
	fn func(context.Context, ServiceRequest) (ServiceResponse, error),
	resp GrpcResponse,
	re ServiceRequest,
) (GrpcResponse, error) {
	serviceRequest := re.ToService(request)

	serviceResponse, err := fn(ctx, serviceRequest)

	if err != nil {
		return *new(GrpcResponse), nil //, HandleAppError(err)
	}

	resp = serviceResponse.FromService(resp)

	return resp, nil
}
