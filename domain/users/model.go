package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"grpc-go/framework/utils/bcrypt"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" jwt:"id"`
	Name      string             `bson:"name" jwt:"name"`
	Email     string             `bson:"email" jwt:"email"`
	Password  string             `bson:"password" jwt:"password"`
	CreatedAt time.Time          `bson:"created_at" jwt:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" jwt:"updated_at"`
}

func NewUser(name string, email string, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Prepare() error {
	password, err := bcrypt.GeneratePassword(user.Password)
	if err != nil {
		return err
	}

	user.ID = primitive.NewObjectID()
	user.Password = password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return nil
}
