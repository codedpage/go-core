package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Config struct {
		Name     string `json:"SiteName"`
		URL      string `json:"SiteUrl"`
		Database struct {
			Name     string
			Host     string
			Port     int
			Username string
			Password string
		}
	}
	conf := Config{}

	jsonString := `{ "SiteName": "My Cat Blog", "SiteUrl": "www.mycatblog.com", "Database": { "Name": "cats", "Host": "localhost", "Port": 3306, "Username": "user1", "Password": "Password1" } } `

	err := json.Unmarshal([]byte(jsonString), &conf)
	//err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	//convert in interface : readable data
	fmt.Println(conf)

	//formatting
	fmt.Printf("Site: %s (%s)\n", conf.Name, conf.URL)
	db := conf.Database
	fmt.Printf(
		"DB: mysql://%s:%s@%s:%d/%s\n",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

}
