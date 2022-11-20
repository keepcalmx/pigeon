package driver

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/keepcalmx/go-pigeon/ent"
	"gopkg.in/yaml.v3"
)

var mysqlConf map[any]any
var mysqlDSN string

func init() {
	yf, err := os.ReadFile("./config/mysql.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yf, &mysqlConf)
	if err != nil {
		log.Fatal(err)
	}

	mysqlDSN = fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&parseTime=True&loc=Local",
		mysqlConf["user"], mysqlConf["password"], mysqlConf["host"], mysqlConf["database"],
	)
}

func MySQL() *ent.Client {
	client, err := ent.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	return client
}
