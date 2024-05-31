package orm

import (
	"my_module/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	res := u.DB.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (u *UserRepo) GetByIdUser(id uint) (*models.User, error) {
	var user models.User

	res := u.DB.First(&user, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (u *UserRepo) Create(user models.User) error {
	res := u.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *UserRepo) Update(user models.User) error {

	res := u.DB.Model(&user).Where("id = ?", user.ID).Updates(models.User{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		Age:        user.Age,
		Field:      user.Field,
		IsEmployee: user.IsEmployee,
	})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *UserRepo) Delete(id uint) error {
	res := u.DB.Where("id = ?", id).Delete(&models.User{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
