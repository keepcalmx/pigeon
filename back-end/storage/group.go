package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/keepcalmx/go-pigeon/db"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/ent/group"
	"github.com/keepcalmx/go-pigeon/ent/user"
)

func CreateGroup(name, avatar string) (*ent.Group, error) {
	client := db.MySQL()
	defer client.Close()
	ctx := context.Background()

	group, err := client.Group.Create().
		SetUUID(uuid.New().String()).
		SetName(name).
		SetAvatar(avatar).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return group, err
}

func AddGroupUser(uuid string, userIDList []string) (*ent.Group, error) {
	client := db.MySQL()
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
	client := db.MySQL()
	defer client.Close()
	ctx := context.Background()

	groupUsers, err := client.Group.Query().
		Where(group.UUIDEQ(uuid)).
		QueryUsers().
		All(ctx)
	return groupUsers, err
}

func ListAllGroups() ([]*ent.Group, error) {
	client := db.MySQL()
	defer client.Close()
	ctx := context.Background()

	allGroups, err := client.Group.Query().All(ctx)
	return allGroups, err
}
