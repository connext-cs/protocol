package api

type TUserPrivilege struct {
	Id          int32 `json:"id"  xorm:"autoincr"`
	UserId      int32 `json:"user_id"`
	PrivilegeId int32 `json:"privilege_id"`
	Deleted     int32 `json:"deleted"`
}
