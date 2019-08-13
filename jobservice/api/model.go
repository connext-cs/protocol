package api

import "time"

type (
	JobConfig struct {
		Id         int       `json:"id"`
		JobName    string    `json:"job_name"`
		JobConf    string    `json:"job_conf"`
		Desc       string    `json:"desc"`
		CreateTime time.Time `json:"create_time"`
	}
)
