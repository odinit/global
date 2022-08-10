package tables

type Job struct {
	Id            string `gorm:"column:id;type:string;size:255" json:"id"`
	JobId         string `gorm:"column:jobId;type:string;size:255" json:"jobId"`
	JobName       string `gorm:"column:jobName;type:string;size:255" json:"jobName"`
	UserName      string `gorm:"column:userName;type:string;size:255" json:"userName"`
	Queue         string `gorm:"column:queue;type:string;size:255" json:"queue"`
	JobState      string `gorm:"column:jobState;type:string;size:255" json:"jobState"`
	Start         string `gorm:"column:start;type:string;size:255" json:"start"`
	End           string `gorm:"column:end;type:string;size:255" json:"end"`
	Type          string `gorm:"column:type;type:string;size:255" json:"type"`
	RunNodes      string `gorm:"column:runNodes;type:string;size:255" json:"runNodes"`           // 使用几个运行节点
	RunNodesCores string `gorm:"column:runNodesCores;type:string;size:255" json:"runNodesCores"` // 运行每个节点的核心数
	JobMem        string `gorm:"column:jobMem;type:string;size:255" json:"jobMem"`               // job使用的Mem,单位MB
	JobSign       string `gorm:"column:jobSign;type:string;size:255" json:"jobSign"`
}
