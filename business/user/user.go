package user

import (
	"time"
)

//User
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	HandPhone string
	Address   string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}

//NewUser create new User
func NewUser(
	id string,
	email string,
	firstname string,
	lastname string,
	password string,
	creator string,
	createdAt time.Time) User {

	return User{
		ID:        id,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
		CreatedAt: createdAt,
		CreatedBy: creator,
	}
}

//ModifyUser update existing UserData
func (oldData *User) ModifyUser(
	newFirstName,
	newLastName,
	newHandPhone,
	newAddress string,
	updatedAt time.Time,
	updater string) User {

	return User{
		ID:        oldData.ID,
		Email:     oldData.Email,
		FirstName: newFirstName,
		LastName:  newLastName,
		Password:  oldData.Password,
		HandPhone: newHandPhone,
		Address:   newAddress,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: updater,
	}
}

// //ModifyUserToken update existing UserData
// func (oldData *User) ModifyUserToken(
// 	newToken string,
// 	updatedAt time.Time) User {

// 	return User{
// 		ID:        oldData.ID,
// 		Email:     oldData.Email,
// 		FirstName: oldData.FirstName,
// 		LastName:  oldData.LastName,
// 		Password:  oldData.Password,
// 		HandPhone: oldData.HandPhone,
// 		Address:   oldData.Address,
// 		Token:     newToken,
// 		CreatedAt: oldData.CreatedAt,
// 		CreatedBy: oldData.CreatedBy,
// 		UpdatedAt: updatedAt,
// 		UpdatedBy: oldData.ID,
// 	}
// }

//ModifyUserPassword update existing UserData
func (oldData *User) ModifyUserPassword(
	newPassword string,
	updater string,
	updatedAt time.Time) User {

	return User{
		ID:        oldData.ID,
		Email:     oldData.Email,
		FirstName: oldData.FirstName,
		LastName:  oldData.LastName,
		Password:  newPassword,
		HandPhone: oldData.HandPhone,
		Address:   oldData.Address,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: updater,
	}
}

func (oldData *User) DeleteUser(
	deleteAt time.Time,
	deleter string) User {

	return User{
		ID:        oldData.ID,
		Email:     oldData.Email,
		FirstName: oldData.FirstName,
		LastName:  oldData.LastName,
		Password:  oldData.Password,
		HandPhone: oldData.HandPhone,
		Address:   oldData.Address,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: oldData.UpdatedAt,
		UpdatedBy: oldData.UpdatedBy,
		DeletedAt: deleteAt,
		DeletedBy: deleter,
	}
}
