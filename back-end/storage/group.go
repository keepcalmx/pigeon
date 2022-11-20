package storage

import (
	"context"
	"time"

	"github.com/keepcalmx/go-pigeon/common/utils"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/ent/group"
	"github.com/keepcalmx/go-pigeon/ent/user"
	"github.com/keepcalmx/go-pigeon/storage/driver"
)

func CreateGroup(name, avatar string) (*ent.Group, error) {
	client := driver.MySQL()
	defer client.Close()
	ctx := context.Background()

	group, err := client.Group.Create().
		SetUUID(utils.UUIDNoHyphen()).
		SetName(name).
		SetAvatar(avatar).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return group, err
}

func AddGroupUser(uuid string, userIDList []string) (*ent.Group, error) {
	client := driver.MySQL()
	defer client.Close()
	ctx := context.Background()

	usersIDs, err := client.User.Query().
		Where(user.UUIDIn(userIDList...)).
		Select(user.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}

	group, err := client.Group.Query().
		Where(group.UUIDEQ(uuid)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	group, err = group.Update().AddUserIDs(usersIDs...).Save(ctx)
	return group, err
}

func ListGroupUsers(uuid string) ([]*ent.User, error) {
	client := driver.MySQL()
	defer client.Close()
	ctx := context.Background()

	groupUsers, err := client.Group.Query().
		Where(group.UUIDEQ(uuid)).
		QueryUsers().
		All(ctx)
	return groupUsers, err
}

func ListGroups() ([]*ent.Group, error) {
	client := driver.MySQL()
	defer client.Close()
	ctx := context.Background()

	allGroups, err := client.Group.Query().All(ctx)
	return allGroups, err
}
