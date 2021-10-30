package user

//Service outgoing port for user
type Service interface {
	//InsertUser Insert new User into storage
	InsertUser(insertUserSpec InsertUserSpec) error

	//FindUserByUsernameAndPassword If data not found will return nil
	FindUserByEmailAndPassword(email string, password string) (*User, error)

	FindUserByEmail(email string) (*User, error)

	//FindUserByID If data not found will return nil without error
	FindUserByID(id string) (*User, error)

	//UpdateUserPaasword if data not found or old password wrong will return error
	UpdateUserPassword(id string, password, oldPassword string, modifier string) error

	//UpdateUserToken if data not found  will return error
	//UpdateUserToken(id string, token string) error

	//UpdateUser if data not found will return error
	UpdateUser(id string, updateUserSpec UpdateUserSpec, modifier string) error

	//Deleteuser if data not found will return error
	DeleteUser(id string, deleter string) error
}

//Repository ingoing port for user
type Repository interface {
	//InsertUser Insert new User into storage
	InsertUser(user User) error

	//FindUserByUsernameAndPassword If data not found will return nil
	FindUserByEmailAndPassword(email string, password string) (*User, error)

	FindUserByEmail(email string) (*User, error)

	//FindUserByID If data not found will return nil without error
	FindUserByID(id string) (*User, error)

	//UpdateUser if data not found will return error
	UpdateUser(user User) error

	//UpdateUser if data not found will return error
	UpdateUserPassword(user User) error

	//Deleteuser if data not found will return error
	DeleteUser(user User) error
}
