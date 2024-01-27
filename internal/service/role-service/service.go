package role_service

//import def "gateway/internal/service"
//
//var _ def.RoleService = (*service)(nil)
//
//type service struct {
//	//producer kafka.KafkaProducer
//}
//
//func NewService(
//// producer kafka.KafkaProducer,
//) *service {
//	return &service{
//		//producer: producer,
//	}
//}

//var _ service.RoleService = (*roleService)(nil)
//
//type roleService struct {
//	//producer kafka.KafkaProducer
//	service service.GlobalService
//}
//
//func NewRoleService(service service.GlobalService) *roleService {
//	return &roleService{
//		service: service,
//	}
//}
//
//func (r *roleService) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
//	response, err := request.Do(ctx, r.service)
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}

type RoleService interface {
	LoginService
}
