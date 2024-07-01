package converter

import (
	"crud-api/src/model"
	"crud-api/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age)
	domain.SetId(entity.Id.Hex())
	return domain
}
