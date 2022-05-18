package repository

import (
	"fmt"

	"github.com/Jiran03/gudhani/user/domain"
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

// // GetByEmailPassword implements domain.Repository
// func (ur userRepo) GetByEmailPassword(email string, password string) (id int, role string, err error) {
// 	newRecord := User{}
// 	err = ur.DB.Where("email = ? AND password = ?", email, password).Find(&newRecord).Error
// 	fmt.Println("ini Repo", newRecord)
// 	newRes := toDomain(newRecord)
// 	fmt.Println("ini RepoRes", newRes)
// 	id = newRes.Id
// 	role = newRes.Role
// 	return id, role, err
// }
// GetByEmailPassword implements domain.Repository
func (ur userRepo) GetByEmailPassword(email, password string) (domain domain.User, err error) {
	var newRecord User
	fmt.Println("repo", email, password)
	err = ur.DB.Where("email = ?", email).First(&newRecord).Error
	domain = toDomain(newRecord)
	fmt.Println(domain)
	return domain, err
}

func (ur userRepo) Get() (userObj []domain.User, err error) {
	var newRecord []User
	err = ur.DB.Find(&newRecord).Error
	fmt.Println(newRecord)
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
	return toDomain(record), err
}

func (ur userRepo) Update(id int, domain domain.User) (userObj domain.User, err error) {
	var newRecord User
	fmt.Println("id repo", id)
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
	userObj = toDomain(newRecord)
	return
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepo{
		DB: db,
	}
}
