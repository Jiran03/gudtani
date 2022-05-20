package repository

import (
	"github.com/Jiran03/gudtani/user/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (ur userRepo) Delete(id int) (err error) {
	var record User
	return ur.DB.Delete(&record, id).Error
}

func (ur userRepo) GetByEmailPassword(email, password string) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.Where("email = ?", email).First(&newRecord).Error

	if err != nil {
		return domain, err
	}

	domain = toDomain(newRecord)
	return domain, nil
}

func (ur userRepo) Get() (userObj []domain.User, err error) {
	var newRecord []User
	err = ur.DB.Find(&newRecord).Error

	if err != nil {
		return userObj, err
	}

	for _, value := range newRecord {
		userObj = append(userObj, toDomain(value))
	}

	return userObj, nil
}

// GetByID implements domain.Repository
func (ur userRepo) GetByID(id int) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.Where("id = ?", id).First(&newRecord).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Save implements domain.Repository
func (ur userRepo) Create(domain domain.User) (userObj domain.User, err error) {
	record := fromDomain(domain)
	err = ur.DB.Create(&record).Error

	if err != nil {
		return userObj, err
	}

	return toDomain(record), nil
}

func (ur userRepo) Update(id int, domain domain.User) (userObj domain.User, err error) {
	var newRecord User
	record := fromDomain(domain)
	err = ur.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":       id,
		"name":     record.Name,
		"email":    record.Email,
		"password": record.Password,
		"address":  record.Address,
		"gender":   record.Gender,
		"role":     record.Role,
	}).Error

	if err != nil {
		return userObj, err
	}

	userObj = toDomain(newRecord)
	return userObj, nil
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepo{
		DB: db,
	}
}
