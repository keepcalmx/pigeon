package storage

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/keepcalmx/go-pigeon/common/data"
	"github.com/keepcalmx/go-pigeon/common/utils"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/ent/user"
	"github.com/keepcalmx/go-pigeon/model/rest"
	"github.com/keepcalmx/go-pigeon/storage/driver"
)

func CreateUser(userForm rest.CreateUserForm) (*ent.User, error) {
	client := driver.MySQL()
	defer client.Close()

	rand.Seed(time.Now().UnixNano())
	randomAvatar := data.UserAvatars[rand.Intn(len(data.UserAvatars))]
	user, err := client.User.Create().
		SetUUID(utils.UUIDNoHyphen()).
		SetUsername(userForm.Username).
		SetPassword(userForm.Password).
		SetEmail(userForm.Email).
		SetNickname(userForm.Nickname).
		SetAvatar(randomAvatar).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	log.Println("create user with form: ", userForm)
	return user, err
}

func GetUserGroups(userID string) ([]*ent.Group, error) {
	client := driver.MySQL()
	defer client.Close()

	user_, err := client.User.Query().
		Where(user.UUIDEQ(userID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	groups, err := user_.QueryGroups().All(context.Background())
	return groups, err
}

func GetUserByUUID(userID string) (*ent.User, error) {
	client := driver.MySQL()
	defer client.Close()

	user_, err := client.User.Query().
		Where(user.UUIDEQ(userID)).
		Only(context.Background())
	return user_, err
}

func GetUserByUsername(username string) (*ent.User, error) {
	client := driver.MySQL()
	defer client.Close()

	user_, err := client.User.Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	return user_, err
}

func GetAllUsers() ([]*ent.User, error) {
	client := driver.MySQL()
	defer client.Close()

	allUsers, err := client.User.Query().All(context.Background())
	return allUsers, err
}
