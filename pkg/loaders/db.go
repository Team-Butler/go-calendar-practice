package loaders

import (
	"encoding/json"
	"fmt"
	"go-calendar-practice/db/ent"
	"io/ioutil"

	_ "github.com/lib/pq"
)

type pgConfigType struct {
	Host string
	Port string
	User string
	Dbname string
	Password string
}

func PostgresSQLConnet(DEBUG_MODE string) (*ent.Client, error) {
	file, readErr := ioutil.ReadFile("config/pg.json")
	if readErr != nil {
		return nil, readErr
	}
	
	var pgConfig pgConfigType
	json.Unmarshal([]byte(file), &pgConfig)
	
	client, connectErr := ent.Open(
		"postgres", 
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
			pgConfig.Host, 
			pgConfig.Port, 
			pgConfig.User, 
			pgConfig.Dbname, 
			pgConfig.Password,
		),
	)
	
	if connectErr != nil {
		return nil, connectErr
	}
	
	return client, nil
}
 