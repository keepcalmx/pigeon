package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	const_ "github.com/keepcalmx/go-pigeon/common/constant"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/ent/group"
	"gopkg.in/yaml.v3"
)

var config map[any]any

func init() {
	yf, err := os.ReadFile("./config/mysql.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yf, &config)
	if err != nil {
		log.Fatal(err)
	}
}

func MySQL() *ent.Client {
	client, err := ent.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&parseTime=True&loc=Local",
			config["user"], config["password"], config["host"], config["database"],
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return client
}

func Migrate() {
	// wait for mysql ready
	time.Sleep(3 * time.Second)

	client := MySQL()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("mysql migrate succeed.")
	CreateDefaults()
}

func CreateDefaults() {
	client := MySQL()
	defer client.Close()
	// 创建默认群组
	defaultGroup, _ := client.Group.Query().
		Where(group.UUIDEQ(const_.SYSTEM_GROUP_UUID)).
		Only(context.Background())
	if defaultGroup == nil {
		_, err := client.Group.Create().
			SetUUID(const_.SYSTEM_GROUP_UUID).
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
