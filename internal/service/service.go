package service

// var _ GlobalService = (*globalService)(nil)

//type ServerService interface {
//	//GetServer(ctx context.Context, uuid string) (*model.Server, error)
//}

//type RoleService interface {
//	Login(ctx context.Context, request *role_service.LoginRequest) (*role_service.LoginResponse, error)
//}

//type GlobalService struct {
//	transfer      transfer.Transfer
//	microServices MicroServices
//}
//
//func NewGlobalService(transfer transfer.Transfer) *GlobalService {
//	return &GlobalService{
//		transfer: transfer,
//	}
//}

//func (g *globalService) Login(ctx context.Context, request *role_service.LoginRequest) (*role_service.LoginResponse, error) {
//	response, err := request.Do(ctx, g)
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}

//type fromServiceToApi[T, R any] interface {
//	FromService(value T) R
//}
//
//type toServiceFromApi[T, R any] interface {
//	ToService(value T) R
//}
//
//func CallService[
//	// ServiceResponse any,
//	GrpcRequest any,
//	GrpcResponse any,
//	ServiceResponse fromServiceToApi[GrpcResponse, GrpcResponse],
//	ServiceRequest toServiceFromApi[GrpcRequest, ServiceRequest],
//	// GrpcResponse fromServiceToApi[ServiceResponse, GrpcResponse],
//](
//	ctx context.Context,
//	//g GlobalService,
//	request GrpcRequest,
//	fn func(context.Context, ServiceRequest) (ServiceResponse, error),
//	resp GrpcResponse,
//	re ServiceRequest,
//) (GrpcResponse, error) {
//
//	serviceRequest := re.ToService(request)
//
//	serviceResponse, err := fn(ctx, serviceRequest)
//
//	if err != nil {
//		return *new(GrpcResponse), nil //, HandleAppError(err)
//	}
//
//	resp = serviceResponse.FromService(resp)
//
//	//response = serviceResponse()
//	//
//	return resp, nil
//}
