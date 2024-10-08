package transfer

import (
	"context"
	"gateway/internal/transfer/grpc/role/model"
)

type fromTransferToService[T, R any] interface {
	FromTransfer(value T) R
}

type toTransferFromService[T, R any] interface {
	ToTransfer(value T) R
}

type RoleTransfer interface {
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
}

type Transfer interface {
	RoleTransfer
}

func sendRequestToBSOrGate[
	ServiceRequest any,
	ServiceResponse any,
](
	ctx context.Context,
	client any,

) {
}

type serviceFunc[Req, Resp any] func(ctx context.Context, request Req) (Resp, error)

func CallService[
	GrpcRequest any,
	GrpcResponse any,
	ServiceRequestBase any,
	// ServiceResponse fromServiceToApi[GrpcResponse, GrpcResponse],
	// ServiceRequest toServiceFromApi[GrpcRequest, ServiceRequest],
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

	resp = serviceResponse.FromService()

	return resp, nil
}
