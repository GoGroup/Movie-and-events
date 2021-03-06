package repository

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/user"
	"github.com/jinzhu/gorm"
)

type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates a new object of UserGormRepo
func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

// Users return all users from the database
// func (userRepo *UserGormRepo) Users() ([]model.User, []error) {
// 	users := []model.User{}
// 	errs := userRepo.conn.Find(&users).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return users, errs
// }
func (userRepo *UserGormRepo) User(id uint) (*model.User, []error) {
	user := model.User{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

// UserByEmail retrieves a user by its email address from the database
func (userRepo *UserGormRepo) UserByEmail(email string) (*model.User, []error) {
	user := model.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

//UpdateUser updates a given user in the database
func (userRepo *UserGormRepo) UpdateUserAmount(user *model.User, Amount uint) *model.User {
	usr := user
	userRepo.conn.Model(&usr).UpdateColumn("amount", Amount)

	return usr
}

// // DeleteUser deletes a given user from the database
// func (userRepo *UserGormRepo) DeleteUser(id uint) (*model.User, []error) {
// 	usr, errs := userRepo.User(id)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	errs = userRepo.conn.Delete(usr, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return usr, errs
// }

// StoreUser stores a new user into the database
func (userRepo *UserGormRepo) StoreUser(user *model.User) (*model.User, []error) {
	usr := user
	errs := userRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (userRepo *UserGormRepo) EmailExists(email string) bool {
	user := model.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}
