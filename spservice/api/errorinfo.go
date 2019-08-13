package api

import "errors"

const (
	ParameterError                    = "参数错误"
	RecordNotexistError               = "操作记录不存在"
	HostNotDestoryed                  = "主机还未销毁，记录不能删除"
	HostStateError                    = "主机状态错误，不能执行操作"
	SoftwareConfigNotExist            = "软件配置不存在"
	SoftwareVersionNotExist           = "软件版本不存在"
	SoftwareNotUninstalled            = "软件未卸载，记录无法删除"
	AnsibleInternalError              = "Ansible执行错误"
	ProjectIsRunning                  = "项目正在运行中，请稍后再试"
	DbExecuteHasNotEffect             = "Db执行不生效"
	DbExecuteHasNotMatch              = "Db插入和实际不同，错误"
	AnsibleExecuteStateNotRight       = "ansible执行状态错误"
	StateNotRight                     = "状态错误，操作失败"
	HostHasNoSoftware                 = "主机上未安装软件，无法操作"
	ManageJobHasNoRecord              = "任务记录不存在，操作失败"
	ManageJobHasNotBeenStarted        = "任务还未开始执行"
	ManageJobExtNotRight              = "任务信息错误"
	InternalError                     = "服务内部错误"
	ManageJobRunningNotExisted        = "未有执行中任务，操作失败"
	ManageJobRunningProcessNotExisted = "进程不存在，操作失败"
	ManageJobRunningKillFailed        = "停止进程失败，服务内部错误"
	HostRecordNotExist                = "主机记录不存在，请联系技术支持人员"
	SoftwareRecordNotExist            = "软件记录不存在，请联系技术支持人员"
)

func Gerr(str string) error {
	return errors.New(str)
}
