package api

import "context"

//type fromServiceToApi[T, R any] interface {
//	FromService(value T) R
//}
//
//type toServiceFromApi[T, R any] interface {
//	ToService(value T) R
//}

type serviceFunc[Req, Resp any] func(ctx context.Context, request Req) (Resp, error)

func CallService[
	GrpcRequest any,
	GrpcResponse any,
	ServiceRequestBase any,
	// ServiceResponse fromServiceToApi[GrpcResponse, GrpcResponse],
	// ServiceRequest toServiceFromApi[GrpcRequest, ServiceRequest],
	ServiceResponse interface {
		FromService(value GrpcResponse) GrpcResponse
	},
	ServiceRequest interface {
		ToService(value GrpcRequest) ServiceRequest
		*ServiceRequestBase
	},
](
	ctx context.Context,
	request GrpcRequest,
	fn serviceFunc[ServiceRequest, ServiceResponse], //func(context.Context, ServiceRequest) (ServiceResponse, error),
	resp GrpcResponse,
	// re ServiceRequest,
) (GrpcResponse, error) {
	//serviceRequest := re.ToService(request)
	serviceRequest := ServiceRequest(new(ServiceRequestBase)).ToService(request)

	serviceResponse, err := fn(ctx, serviceRequest)

	if err != nil {
		return *new(GrpcResponse), nil //, HandleAppError(err)
	}

	resp = serviceResponse.FromService(resp)

	return resp, nil
}
