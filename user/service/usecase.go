package service

import (
	"errors"

	authMiddleware "github.com/Jiran03/gudtani/auth/middleware"
	"github.com/Jiran03/gudtani/user/domain"
)

type userService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// CreateToken implements domain.Service
func (us userService) CreateToken(email, password string) (token string, err error) {

	userObj, err := us.repository.GetByEmailPassword(email, password)

	if err != nil {
		return token, err
	}

	id := userObj.ID
	token = us.jwtAuth.GenerateToken(id)
	return token, nil
}

// InsertData implements domain.Service
func (us userService) InsertData(domain domain.User) (userObj domain.User, err error) {
	userObj, err = us.repository.Create(domain)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// UpdateData implements domain.Service
func (us userService) UpdateData(id int, domain domain.User) (userObj domain.User, err error) {
	user, errGetByID := us.repository.GetByID(id)

	if errGetByID != nil {
		return user, errGetByID
	}

	userId := user.ID
	userObj, err = us.repository.Update(userId, domain)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

func (us userService) GetByEmailPassword(email, password string) (id int, role string, err error) {
	userObj, errRepo := us.repository.GetByEmailPassword(email, password)
	id = userObj.ID
	role = userObj.Role
	err = errRepo

	if err != nil {
		return id, role, err
	}

	return id, role, nil
}

func (us userService) GetAllData() (userObj []domain.User, err error) {
	userObj, err = us.repository.Get()

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

func (us userService) GetByID(id int) (userObj domain.User, err error) {
	userObj, err = us.repository.GetByID(id)

	if err != nil {
		return userObj, errors.New("not found")
	}

	return userObj, nil
}

// DeleteData implements domain.Service
func (us userService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
	// return errConv.Conversion(errResp)
}

func NewUserService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return userService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
