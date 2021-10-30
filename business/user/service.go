package user

import (
	"AltaStore/business"
	"AltaStore/util/validator"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//InsertUserSpec create user spec
type InsertUserSpec struct {
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Password  string `validate:"required"`
}

type UpdateUserSpec struct {
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	HandPhone string
	Address   string
}

type UpdateUserPasswordSpec struct {
	NewPassword string `validate:"required"`
	OldPassword string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	userdata, _ := s.repository.FindUserByEmail(insertUserSpec.Email)
	if userdata != nil {
		return business.ErrDataExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(insertUserSpec.Password), bcrypt.DefaultCost)
	if err != nil {
		return business.ErrInvalidSpec
	}
	var newuuid = uuid.New().String()
	user := NewUser(
		newuuid,
		insertUserSpec.Email,
		insertUserSpec.FirstName,
		insertUserSpec.LastName,
		string(hashedPassword),
		newuuid,
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

//FindUserByUsernameAndPassword If data not found will return nil
func (s *service) FindUserByEmailAndPassword(email string, password string) (*User, error) {
	return s.repository.FindUserByEmailAndPassword(email, password)
}

//FindUserByUsername If data not found will return nil
func (s *service) FindUserByEmail(email string) (*User, error) {
	return s.repository.FindUserByEmail(email)
}

//FindUserByID If data not found will return nil without error
func (s *service) FindUserByID(id string) (*User, error) {
	return s.repository.FindUserByID(id)
}

//UpdateUserPaasword if data not found or old password wrong will return error
func (s *service) UpdateUserPassword(id string, newpassword, oldPassword string, modifier string) error {

	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.DeletedBy != "" {
		return business.ErrNotFound
	} else {
		_, err := s.repository.FindUserByEmailAndPassword(user.Email, oldPassword)
		if err != nil {
			return business.ErrPasswordMisMatch
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
	if err != nil {
		return business.ErrInvalidSpec
	}
	modifiedUser := user.ModifyUserPassword(
		string(hashedPassword),
		modifier,
		time.Now(),
	)

	return s.repository.UpdateUserPassword(modifiedUser)
}

//UpdateUser if data not found will return error
func (s *service) UpdateUser(id string, updateUserSpec UpdateUserSpec, modifier string) error {
	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.DeletedBy != "" {
		return business.ErrNotFound
	}

	modifiedUser := user.ModifyUser(
		updateUserSpec.FirstName,
		updateUserSpec.LastName,
		updateUserSpec.HandPhone,
		updateUserSpec.Address,
		time.Now(),
		modifier,
	)

	return s.repository.UpdateUser(modifiedUser)
}

//Deleteuser if data not found will return error
func (s *service) DeleteUser(id string, modifier string) error {
	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	}

	deleteUser := user.DeleteUser(
		time.Now(),
		modifier,
	)

	return s.repository.DeleteUser(deleteUser)
}
