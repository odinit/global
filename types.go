package global

import (
	"microbigdata/utils/mysql"
	"time"
)

type MF map[string]interface{}
type MS map[string]string

type AppParameter struct {
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
	Component string      `json:"component"`
	Choice    string      `json:"choice"`
	Tips      string      `json:"tips"`
	Required  bool        `json:"required"`
	Path      string      `json:"path"`
}

// App--------------------------------------------------

type AppCmdExecReq struct {
	ContainerId string `json:"containerId"` // 容器id
	Cmd         string `json:"cmd"`         // 执行命令
}

type AppOutputListReq struct {
	ContainerId string `json:"containerId"` // 应用容器id
	OutputPath  string `json:"outputPath"`  // 输出文件路径
}

type AppParameterGenerateReq struct {
	ContainerId string `json:"containerId"`
	FuncId      string `json:"funcId"`
	Parameter   MF     `json:"parameter"`
}

type AppInputUploadReq struct {
	ContainerId   string `json:"containerId"`   // 应用容器id
	ParameterName string `json:"parameterName"` // 上传文件对应的参数名称
}

type AppOutputDataReq struct {
	ContainerId string `json:"containerId"` // 应用容器id
	FilePath    string `json:"filePath"`    // 输出文件路径
}

type AppOutputDownloadReq = AppOutputDataReq

// OutputFilter 用于过滤输出文件,并匹配文件描述信息
type OutputFilter struct {
	Regex       string
	Description string
}

type HostInfo struct {
	IP  string  `json:"ip"`
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}

// Docker 容器管理中心--------------------------------------------------

type DockerContainerCreateReq struct {
	IP            string            `json:"ip"`            // 宿主机ip
	ContainerInfo mysql.DockerImage `json:"containerInfo"` // 容器信息
}
type DockerContainerStartReq struct {
	IP          string `json:"ip"`          // 宿主机ip
	ContainerId string `json:"containerId"` //容器id
}

type DockerContainerStartRes struct {
	Code        int    `json:"code"`        // 请求状态
	Msg         string `json:"msg"`         // 请求状态信息
	ContainerId string `json:"containerId"` // 容器id
}
type DockerContainerRunReq = DockerContainerCreateReq
type DockerContainerRunRes = DockerContainerStartRes
type DockerContainerCreateRes = DockerContainerStartRes
type DockerContainerStopReq = DockerContainerStartReq
type DockerContainerRemoveReq = DockerContainerStartReq
type DockerContainerRemoveRes = DockerContainerStartRes
type DockerContainerLogsReq = DockerContainerStartReq

// Hard 硬件中心--------------------------------------------------

type HardScoreRes struct {
	Code int     `json:"code"` // 请求状态
	Msg  string  `json:"msg"`  // 请求状态信息
	Cpu  float64 `json:"cpu"`
	Mem  float64 `json:"mem"`
}

type HardPortRes struct {
	Code int    `json:"code"` // 请求状态
	Msg  string `json:"msg"`  // 请求状态信息
	Port string `json:"port"`
}

type HardPortRes1 struct {
	Code int      `json:"code"` // 请求状态
	Msg  string   `json:"msg"`  // 请求状态信息
	Port []string `json:"port"`
}

type CPU struct {
	Time  string
	Value float64
}

type Mem struct {
	Name  string
	Type  string
	Value interface{}
}

type Mems []Mem

type HardResourceRes struct {
	Code int    `json:"code"` // 请求状态
	Msg  string `json:"msg"`  // 请求状态信息
	Cpu  []CPU  `json:"cpu"`
	Mem  Mems   `json:"mem"`
}

// Hook 钩子程序--------------------------------------------------

type HookCmdExecReq struct {
	Cmd string `json:"cmd"` // 启动命令
}

type HookCmdExecRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type HookOutputDataReq struct {
	FilePath string `json:"filePath"`
}

type HookOutputDownloadReq = HookOutputDataReq

type ContainerStopRes struct {
	Code        int    `json:"code"`        // 请求状态
	Msg         string `json:"msg"`         // 请求状态信息
	ContainerId string `json:"containerId"` // 容器id
}

type ContainerLogsRes struct {
	Code        int    `json:"code"`        // 请求状态
	Msg         string `json:"msg"`         // 请求状态信息
	ContainerId string `json:"containerId"` // 容器id
	Logs        string `json:"logs"`        // 容器日志
}

type StatusGetReq struct {
	ContainerId string `json:"containerId"` // 应用容器id
}

type StatusGetRes struct {
	Code      int    `json:"code"`      // 请求状态
	Msg       string `json:"msg"`       // 请求状态信息
	AppStatus int    `json:"appStatus"` // 容器id
}

type PathInfo struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`    // 名称
	Size    string    `json:"size"`    // 大小
	IsDir   bool      `json:"isDir"`   // 是否是文件夹
	ModTime time.Time `json:"modTime"` // 修改时间
	Path    string    `json:"path"`    // 基于网盘根目录的路径
}

type OutputDataRes struct {
	Code     int         `json:"code"`     // 请求状态
	Msg      string      `json:"msg"`      // 请求状态信息
	FileData interface{} `json:"fileData"` // 文件内容
}

type HookOutputListReq struct {
	OutputPath string  `json:"outputPath"` // 输出文件路径
	Options    Options `json:"options"`    // 输出文件选项
}

type HookOutputListRes struct {
	Code       int              `json:"code"`       // 请求状态
	Msg        string           `json:"msg"`        // 请求状态信息
	OutputList []OutputFileInfo `json:"outputList"` // 输出列表信息
}

type OutputFileInfo struct {
	FileName        string `json:"fileName"`        // 文件名称
	FilePath        string `json:"filePath"`        // 文件路径
	FileSize        string `json:"fileSize"`        // 文件大小
	FileDescription string `json:"fileDescription"` // 文件描述
}

type Options struct {
	ExtendInclude  []string `json:"extendInclude"`  // 需要包含的扩展名
	ExtendExclude  []string `json:"extendExclude"`  // 需要排除的扩展名
	KeywordInclude []string `json:"keywordInclude"` // 需要包含的关键字
	KeywordExclude []string `json:"keywordExclude"` // 需要排除的关键字
}

type HookParameterGenerateReq struct {
	Parameter         MF     `json:"parameter"`         // 参数
	Template          string `json:"template"`          // 参数模板
	ParameterFilePath string `json:"parameterFilePath"` // 参数文件路径
	ParameterFileName string `json:"parameterFileName"` // 参数文件名称
}

// Disk 网盘管理服务--------------------------------------------------

type DiskPathListReq struct {
	Path string `json:"path"`
}

type DiskPathMoveReq struct {
	OldPath string `json:"oldPath"`
	NewDir  string `json:"newDir"`
}

type DiskPathRenameReq struct {
	OldPath string `json:"oldPath"`
	NewName string `json:"newName"`
}

type DiskPathDuplicateReq = DiskPathListReq
type DiskPathRemoveReq = DiskPathListReq
type DiskFileListReq = DiskPathListReq
type DiskDirListReq = DiskPathListReq
type DiskFileDownloadReq = DiskPathListReq

type DiskDirMakeReq struct {
	Path    string `json:"path"`
	DirName string `json:"dirName"`
}

type DiskFileSearchReq struct {
	Path     string `json:"path"`
	Keywords string `json:"keywords"`
	Extends  string `json:"extends"`
}

// Hdfs 文件管理系统服务--------------------------------------------------

type HdfsDownloadRes struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    bool   `json:"data"`
}

// Gateway

type GatewayNginxPortMapRes struct {
	Code int    `json:"code"` // 请求状态
	Msg  string `json:"msg"`  // 请求状态信息
	Port string `json:"port"` // 映射后的端口
}

type AppManageNewReq struct {
	AppNameKeyword string `json:"appNameKeyword"` // 搜索模型名称的关键词
	SortField      string `json:"sortField"`      // 排序条件
	IsAsc          bool   `json:"isAsc"`          // 是否升序
	StatusFilter   int    `json:"statusFilter"`   // 运行状态筛选
	PageNum        int    `json:"pageNum"`        // 页码
	PageSize       int    `json:"pageSize"`       // 每页的数量
}
