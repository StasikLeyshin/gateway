package api

import "context"

type serviceFunc[Req, Resp any] func(ctx context.Context, request Req) (Resp, error)

func CallService[
	GrpcRequest any,
	GrpcResponse any,
	ServiceRequestBase any,
	ServiceResponse interface {
		FromService() GrpcResponse
	},
	ServiceRequest interface {
		ToService(value GrpcRequest) ServiceRequest
		*ServiceRequestBase
	},
](
	ctx context.Context,
	request GrpcRequest,
	fn serviceFunc[ServiceRequest, ServiceResponse],
) (GrpcResponse, error) {
	serviceRequest := ServiceRequest(new(ServiceRequestBase)).ToService(request)

	serviceResponse, err := fn(ctx, serviceRequest)

	if err != nil {
		return *new(GrpcResponse), nil // TODO: Добавить обработку ошибок
	}

	resp := serviceResponse.FromService()

	return resp, nil
}
