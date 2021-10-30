package user

import (
	"AltaStore/business/user"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Implementasi repositori user
type Repository struct {
	DB *gorm.DB
}

type User struct {
	ID        string    `gorm:"type:uuid;primary_key"`
	Email     string    `gorm:"email;index:idx_email,unique;type:varchar(50)"`
	FirstName string    `gorm:"firstname;type:varchar(50)"`
	LastName  string    `gorm:"lastname;type:varchar(50)"`
	Password  string    `gorm:"password;type:varchar(100)"`
	HandPhone string    `gorm:"handphone;type:varchar(50)"`
	Address   string    `gorm:"address;type:varchar(100)"`
	CreatedAt time.Time `gorm:"created_at"`
	CreatedBy string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt time.Time `gorm:"updated_at"`
	UpdatedBy string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt time.Time `gorm:"deleted_at"`
	DeletedBy string    `gorm:"deleted_by;type:varchar(50)"`
}

func newUserTable(user user.User) *User {

	return &User{
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Password,
		user.HandPhone,
		user.Address,
		user.CreatedAt,
		user.CreatedBy,
		user.UpdatedAt,
		user.UpdatedBy,
		user.DeletedAt,
		user.DeletedBy,
	}
}

func (col *User) ToUser() user.User {
	var user user.User

	user.ID = col.ID
	user.Email = col.Email
	user.FirstName = col.FirstName
	user.LastName = col.LastName
	user.HandPhone = col.HandPhone
	user.Address = col.Address
	user.CreatedAt = col.CreatedAt
	user.CreatedBy = col.CreatedBy
	user.UpdatedAt = col.UpdatedAt
	user.UpdatedBy = col.UpdatedBy
	user.DeletedAt = col.DeletedAt
	user.DeletedBy = col.DeletedBy

	return user
}

// Menghasilkan ORM DB untuk user repository
func NewDBRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

//InsertUser Insert new User into storage
func (repo *Repository) InsertUser(user user.User) error {

	userData := newUserTable(user)

	err := repo.DB.Create(userData).Error
	if err != nil {
		return err
	}

	return nil
}

//FindUserByEmailAndPassword If data not found will return nil
func (repo *Repository) FindUserByEmailAndPassword(email string, password string) (*user.User, error) {

	var userData User

	err := repo.DB.Where("email = ?", email).Where("deleted_by = ''").First(&userData).Error
	if err != nil {
		return nil, err
	}

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

func (repo *Repository) FindUserByEmail(email string) (*user.User, error) {

	var userData User

	err := repo.DB.Where("email = ?", email).Where("deleted_by = ''").First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindUserByID If data not found will return nil without error
func (repo *Repository) FindUserByID(id string) (*user.User, error) {

	var userData User

	err := repo.DB.Where("id = ?", id).Where("deleted_by = ''").First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//UpdateUserPassword if data not found or old password wrong will return error
func (repo *Repository) UpdateUserPassword(user user.User) error {
	userData := newUserTable(user)

	err := repo.DB.Model(&userData).Updates(User{
		Password: userData.Password,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser Update existing user in database
func (repo *Repository) UpdateUser(user user.User) error {
	userData := newUserTable(user)

	err := repo.DB.Model(&userData).Updates(User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Address:   userData.Address,
		Password:  userData.Password,
		HandPhone: userData.HandPhone,
		UpdatedAt: userData.UpdatedAt,
		UpdatedBy: userData.UpdatedBy,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

//DeleteUser Set IsDelete true in database
func (repo *Repository) DeleteUser(user user.User) error {
	userData := newUserTable(user)

	err := repo.DB.Model(&userData).Updates(User{
		DeletedBy: userData.DeletedBy,
		DeletedAt: userData.DeletedAt,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
