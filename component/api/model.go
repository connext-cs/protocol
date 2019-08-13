package api

import "time"

type (
	Component struct {
		Id              int       `json:"id" xorm:"autoincr"`
		ComponentConfId int       `json:"component_conf_id"`
		Url             string    `json:"url"`
		Desc            string    `json:"desc"`
		CreateTime      time.Time `json:"create_time"`
	}
	ComponentConf struct {
		Id         int       `json:"id" xorm:"autoincr"`
		Name       string    `json:"name"`
		CreateTime time.Time `json:"create_time"`
	}
)
