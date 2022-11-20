package storage

import (
	"context"

	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/ent/privatemsg"
	"github.com/keepcalmx/go-pigeon/storage/driver"
)

func CreatePrivateMsg(msg *ent.PrivateMsg) (*ent.PrivateMsg, error) {
	client := driver.MySQL()
	defer client.Close()

	privateMsg, err := client.PrivateMsg.Create().
		SetFrom(msg.From).
		SetTo(msg.To).
		SetType(msg.Type).
		SetRead(msg.Read).
		SetContent(msg.Content).
		SetCreatedAt(msg.CreatedAt).
		Save(context.Background())
	return privateMsg, err
}

func ListPrivateMsgs(from, to string) ([]*ent.PrivateMsg, error) {
	client := driver.MySQL()
	defer client.Close()

	privateMsgs, err := client.PrivateMsg.Query().
		Where(privatemsg.And(privatemsg.FromIn(from, to), (privatemsg.ToIn(from, to)))).
		All(context.Background())
	return privateMsgs, err
}

func GetUserPrivateMsgs(from string, offset, limit int) ([]*ent.PrivateMsg, error) {
	client := driver.MySQL()
	defer client.Close()

	privateMsgs, err := client.PrivateMsg.Query().
		Where(privatemsg.FromEQ(from)).
		Offset(offset).Limit(limit).
		All(context.Background())
	return privateMsgs, err
}
