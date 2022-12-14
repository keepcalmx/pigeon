package storage

import (
	"context"

	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/ent/groupmsg"
	"github.com/keepcalmx/go-pigeon/storage/driver"
)

func CreateGroupMsg(msg *ent.GroupMsg) (*ent.GroupMsg, error) {
	client := driver.MySQL()
	defer client.Close()

	groupMsg, err := client.GroupMsg.Create().
		SetFrom(msg.From).
		SetTo(msg.To).
		SetType(msg.Type).
		SetContent(msg.Content).
		SetCreatedAt(msg.CreatedAt).
		Save(context.Background())
	return groupMsg, err
}

func CreateBulkGroupMsg(msgs []*ent.GroupMsg) ([]*ent.GroupMsg, error) {
	client := driver.MySQL()
	defer client.Close()

	bulk := make([]*ent.GroupMsgCreate, len(msgs))
	for i, msg := range msgs {
		bulk[i] = client.GroupMsg.Create().
			SetFrom(msg.From).
			SetTo(msg.To).
			SetType(msg.Type).
			SetContent(msg.Content).
			SetCreatedAt(msg.CreatedAt)
	}
	groupMsgs, err := client.GroupMsg.CreateBulk(bulk...).Save(context.Background())
	return groupMsgs, err
}

func GetGroupMsgs(uuid string, offset, limit int) ([]*ent.GroupMsg, error) {
	client := driver.MySQL()
	defer client.Close()

	groupMsgs, err := client.GroupMsg.Query().
		Where(groupmsg.ToEQ(uuid)).
		Offset(offset).Limit(limit).
		All(context.Background())
	return groupMsgs, err
}
