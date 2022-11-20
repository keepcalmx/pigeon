package cache

import (
	"fmt"
	"log"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/keepcalmx/go-pigeon/cache/driver"
	C "github.com/keepcalmx/go-pigeon/common/constant"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/storage"
)

var groupMsgEngine *redisearch.Client

// CacheGroupMsg 缓存存量群组消息
func CacheGroupMsg() {
	groupMsgEngine = driver.SearchEngine("pigeon_group_msg")

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("id")).
		AddField(redisearch.NewTextField("from")).
		AddField(redisearch.NewTextField("to")).
		AddField(redisearch.NewTextField("content")).
		AddField(redisearch.NewTextField("type"))

	groupMsgEngine.Drop()

	opt := redisearch.NewIndexDefinition().SetLanguage("chinese")
	if err := groupMsgEngine.CreateIndexWithIndexDefinition(sc, opt); err != nil {
		log.Fatal(err)
	}

	groups, err := storage.ListGroups()
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups {
		msgs, err := storage.GetGroupMsgs(
			group.UUID, 0, C.DEFAULT_GROUP_MSG_CACHE_LIMIT,
		)
		if err != nil {
			log.Fatal(err)
		}

		Documents := make([]redisearch.Document, len(msgs))
		for _, msg := range msgs {
			Documents = append(Documents, redisearch.NewDocument(strconv.Itoa(msg.ID), 1.0).
				Set("id", msg.ID).
				Set("from", msg.From).
				Set("to", msg.To).
				Set("content", msg.Content).
				Set("type", msg.Type))
		}

		groupMsgEngine.Index(Documents...)
		log.Println("cache group msg from mysql succeed.")
	}
}

func QueryGroupMsg(uuid, keyword string, offset, num int) (docs []redisearch.Document, total int, err error) {
	docs, total, err = groupMsgEngine.Search(
		redisearch.NewQuery(fmt.Sprintf("@to:%s @content:%s", uuid, keyword)).
			SetLanguage("chinese").
			Limit(offset, num),
	)

	return docs, total, err
}

func AddGroupMsg(msg *ent.GroupMsg) error {
	err := groupMsgEngine.Index(redisearch.NewDocument(strconv.Itoa(msg.ID), 1.0).
		Set("id", msg.ID).
		Set("from", msg.From).
		Set("to", msg.To).
		Set("content", msg.Content).
		Set("type", msg.Type))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
