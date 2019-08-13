package user

type EnumPrivilege uint8

const (
	SUPER                 EnumPrivilege = 1  //超级用户
	INFAR                 EnumPrivilege = 2  //infar模块用户
	PROJECT               EnumPrivilege = 3  //项目用户
	ONLYLOOK              EnumPrivilege = 4  //仅查看
	ONLYPUBLISH           EnumPrivilege = 5  //仅发布
	VMPROJECT             EnumPrivilege = 6  //vm项目管理
	VMONLYLOOK            EnumPrivilege = 7  //vm仅查看
	VMONLYPUBLISH         EnumPrivilege = 8  //vm仅发布
	APPAdministrator      EnumPrivilege = 9  //应用管理员
	APPView               EnumPrivilege = 10 //应用查看者
	APPSuperAdministrator EnumPrivilege = 11 //应用超级管理员
	MaintenanceSupport    EnumPrivilege = 12 //运维管理员
)

func GetEnumPrivilegeString(p EnumPrivilege) string {
	var jobstatusArray = [MaintenanceSupport]string{"SUPER", "INFAR", "PROJECT", "ONLYLOOK", "ONLYPUBLISH", "VMPROJECT",
		"VMONLYLOOK", "VMONLYPUBLISH", "APPAdministrator", "APPView", "APPSuperAdministrator", "MaintenanceSupport"}
	return jobstatusArray[uint8(p)-1]
}
