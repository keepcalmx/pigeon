package storage

import (
	"context"
	"log"
	"time"

	C "github.com/keepcalmx/go-pigeon/common/constant"
	"github.com/keepcalmx/go-pigeon/ent/group"
	"github.com/keepcalmx/go-pigeon/storage/driver"
)

func Migrate() {
	// wait for mysql ready
	time.Sleep(2 * time.Second)

	client := driver.MySQL()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("mysql migrate succeed.")
	CreateDefaults()
}

func CreateDefaults() {
	client := driver.MySQL()
	defer client.Close()
	// 创建默认群组
	defaultGroup, _ := client.Group.Query().
		Where(group.UUIDEQ(C.SYSTEM_GROUP_UUID)).
		Only(context.Background())
	if defaultGroup == nil {
		_, err := client.Group.Create().
			SetUUID(C.SYSTEM_GROUP_UUID).
			SetName("默认群组").
			SetAvatar("https://findicons.com/files/icons/2779/simple_icons/1024/codepen_1024_black.png").
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			log.Println("create default group failed with err ", err)
		}
		log.Println("create default group succeed")
	}
}
