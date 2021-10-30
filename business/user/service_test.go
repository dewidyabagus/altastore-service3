package user_test

import (
	"AltaStore/business"
	"AltaStore/business/user"
	userMock "AltaStore/business/user/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id        = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	email     = "email@test.com"
	firstname = "firstname"
	lastname  = "lastname"
	password  = "password"
	handphone = "handphone"
	address   = "address"
	createdby = "creator"
	updatedby = "modifier"
	deletedby = ""
)

var (
	userService    user.Service
	userRepository userMock.Repository

	userData       user.User
	insertUserData user.InsertUserSpec
	updateUserData user.UpdateUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
func setup() {
	userData = user.User{
		ID:        id,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
		HandPhone: handphone,
		Address:   address,
	}

	insertUserData = user.InsertUserSpec{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
	}

	updateUserData = user.UpdateUserSpec{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		HandPhone: handphone,
		Address:   address,
	}

	userService = user.NewService(&userRepository)
}

func TestFindUserByID(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByID(id)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, firstname, user.FirstName)
		assert.Equal(t, lastname, user.LastName)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, handphone, user.HandPhone)
		assert.Equal(t, address, user.Address)
	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		user, err := userService.FindUserByID(id)

		assert.NotNil(t, err)
		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByEmail", mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByEmail(email)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, firstname, user.FirstName)
		assert.Equal(t, lastname, user.LastName)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, handphone, user.HandPhone)
		assert.Equal(t, address, user.Address)
	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByEmail", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		user, err := userService.FindUserByEmail(email)

		assert.NotNil(t, err)
		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindUserByEmailAndPassword(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByEmailAndPassword(email, password)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, firstname, user.FirstName)
		assert.Equal(t, lastname, user.LastName)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, handphone, user.HandPhone)
		assert.Equal(t, address, user.Address)
	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		user, err := userService.FindUserByEmailAndPassword(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertUser(t *testing.T) {
	t.Run("Expect user email exist", func(t *testing.T) {
		userRepository.On("FindUserByEmail", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		err := userService.InsertUser(insertUserData)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrDataExists)
	})
	t.Run("Expect insert user success", func(t *testing.T) {
		userRepository.On("FindUserByEmail", mock.AnythingOfType("string")).Return(nil, nil).Once()
		userRepository.On("InsertUser", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.InsertUser(insertUserData)

		assert.Nil(t, err)
	})

	t.Run("Expect insert user failed", func(t *testing.T) {
		userRepository.On("FindUserByEmail", mock.AnythingOfType("string")).Return(nil, nil).Once()
		userRepository.On("InsertUser", mock.AnythingOfType("user.User")).Return(business.ErrInternalServer).Once()

		err := userService.InsertUser(insertUserData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		err := userService.UpdateUser(id, updateUserData, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect update user success", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.UpdateUser(id, updateUserData, id)

		assert.Nil(t, err)
	})
	t.Run("Expect update user failed", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User")).Return(business.ErrInternalServer).Once()

		err := userService.UpdateUser(id, updateUserData, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdateUserPassword(t *testing.T) {
	t.Run("Expect user not found by id", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		err := userService.UpdateUserPassword(id, password, password, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect user not found by email and password", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, business.ErrPasswordMisMatch).Once()
		err := userService.UpdateUserPassword(id, password, password, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrPasswordMisMatch)
	})
	t.Run("Expect update user password success", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUserPassword", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.UpdateUserPassword(id, password, password, id)

		assert.Nil(t, err)
	})
	t.Run("Expect update user password failed", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("FindUserByEmailAndPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("UpdateUserPassword", mock.AnythingOfType("user.User")).Return(business.ErrInternalServer).Once()

		err := userService.UpdateUserPassword(id, password, password, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		err := userService.DeleteUser(id, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect delete user success", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("DeleteUser", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.DeleteUser(id, id)

		assert.Nil(t, err)
	})
	t.Run("Expect delete user failed", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		userRepository.On("DeleteUser", mock.AnythingOfType("user.User")).Return(business.ErrInternalServer).Once()

		err := userService.DeleteUser(id, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}
