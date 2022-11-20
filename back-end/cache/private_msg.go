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

var privateMsgEngine *redisearch.Client

// CachePrivateMsg 缓存存量私聊消息
func CachePrivateMsg() {
	privateMsgEngine = driver.SearchEngine("pigeon_private_msg")

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("id")).
		AddField(redisearch.NewTextField("from")).
		AddField(redisearch.NewTextField("to")).
		AddField(redisearch.NewTextField("content")).
		AddField(redisearch.NewTextField("type"))

	privateMsgEngine.Drop()

	opt := redisearch.NewIndexDefinition().SetLanguage("chinese")
	if err := privateMsgEngine.CreateIndexWithIndexDefinition(sc, opt); err != nil {
		log.Fatal(err)
	}

	users, err := storage.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		msgs, err := storage.GetUserPrivateMsgs(
			user.UUID, 0, C.DEFAULT_PRIVATE_MSG_CACHE_LIMIT,
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

		privateMsgEngine.Index(Documents...)
		log.Println("cache private msg from mysql succeed.")
	}
}

func QueryPrivateMsg(from, to, keyword string, offset, num int) (docs []redisearch.Document, total int, err error) {
	docs, total, err = privateMsgEngine.Search(
		redisearch.NewQuery(fmt.Sprintf("@from: %s @to:%s @content:%s", from, to, keyword)).
			SetLanguage("chinese").
			Limit(offset, num),
	)

	return docs, total, err
}

func AddPrivateMsg(msg *ent.PrivateMsg) error {
	err := privateMsgEngine.Index(redisearch.NewDocument(strconv.Itoa(msg.ID), 1.0).
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
