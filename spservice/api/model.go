package api

import "time"

type (
	HostInfo struct {
		Id          int32     `json:"id" xorm:"autoincr"`
		ManageJobId int32     `json:"manage_job_id"`
		HostName    string    `json:"host_name"`
		HostType    int32     `json:"host_type"`
		HostIp      string    `json:"host_ip"`
		HostPort    int32     `json:"host_port"`
		LoginType   int32     `json:"login_type"`
		UserName    string    `json:"user_name"`
		UserPasswd  string    `json:"user_passwd"`
		SecretKey   string    `json:"secret_key"`
		Desc        string    `json:"desc"`
		CreateTime  time.Time `json:"create_time"`
		DelFlag     int32     `json:"del_flag"`
	}
	SoftwareConfig struct {
		Id               int32     `json:"id" xorm:"autoincr"`
		SoftwareType     int32     `json:"software_type"`
		SoftwareName     string    `json:"software_name"`
		SoftwareVersions string    `json:"software_versions"`
		Desc             string    `json:"desc"`
		CreateTime       time.Time `json:"create_time"`
		DelFlag          int32     `json:"del_flag"`
	}
	SoftwareInfo struct {
		Id                int32     `json:"id" xorm:"autoincr"`
		HostId            int32     `json:"host_id"`
		SoftwareId        int32     `json:"software_id"`
		SoftwareVersion   string    `json:"software_version"`
		SoftwareConfigKey string    `json:"software_config_key"`
		Desc              string    `json:"desc"`
		CreateTime        time.Time `json:"create_time"`
		UpdateTime        time.Time `json:"update_time"`
		DelFlag           int32     `json:"del_flag"`
	}
	Manage struct {
		Id          int32     `json:"id" xorm:"autoincr"`
		ProjectId   string    `json:"project_id"`
		State       int32     `json:"state"`
		Desc        string    `json:"desc"`
		CreateTime  time.Time `json:"create_time"`
		DelFlag     int32     `json:"del_flag"`
		ConfigFiles string    `json:"config_files"`
		UpdateTime  time.Time `json:"update_time"`
	}
	ManageJobList struct {
		Id                int32     `json:"id" xorm:"autoincr"`
		ManageId          int32     `json:"manage_id"`
		ExecuteIndexOrder int32     `json:"execute_index_order"`
		State             int32     `json:"state"`
		CreateTime        time.Time `json:"create_time"`
		DelFlag           int32     `json:"del_flag"`
		Ext               string    `json:"ext"`
	}
)

//todo 涉及到外部模块的一些job临时加入
type (
	TClusterNodeJob struct {
		Id            int32     `json:"id" xorm:"autoincr"`
		ClusterNodeId int32     `json:"cluster_node_id"`
		ClusterId     int32     `json:"cluster_id"`
		ActionType    int32     `json:"action_type"`
		ExcuteState   int32     `json:"excute_state"`
		ExcuteOrder   int32     `json:"excute_order"`
		Ext           string    `json:"ext"`
		CreateTime    time.Time `json:"create_time"`
		UpdateTime    time.Time `json:"update_time"`
		DelFlag       int32     `json:"del_flag"`
	}
	TClusterNode struct {
		ClusterNodeId       int32     `json:"cluster_node_id"`
		ClusterId           int32     `json:"cluster_id"`
		ClusterNodeName     string    `json:"cluster_node_name"`
		ClusterNodeIp       string    `json:"cluster_node_ip"`
		ClusterNodeType     int32     `json:"cluster_node_type"`
		ClusterNodeUsername string    `json:"cluster_node_username"`
		ClusterNodePassword string    `json:"cluster_node_password"`
		ClusterNodeState    int32     `json:"cluster_node_state"`
		Deleted             int32     `json:"deleted"`
		Createtime          time.Time `json:"createtime"`
	}
)
