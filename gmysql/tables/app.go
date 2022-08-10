package tables

// AppMeta 记录应用的元数据
type AppMeta struct {
	ID         uint   `gorm:"primarykey" json:"-"`
	AppId      string `gorm:"column:app_id;type:string;size:255;not null;comment:应用id"`    // 应用id
	AppName    string `gorm:"column:app_name;type:string;size:255;not null;comment:应用名称"`  // 应用名称
	AppClass   string `gorm:"column:app_class;type:string;size:255;not null;comment:应用分类"` // 应用分类
	AppPicture string `gorm:"column:app_picture;type:string;size:255;default:default.jpg;comment:应用缩略图"`
	ImageName  string `gorm:"column:image_name;type:string;size:255;not null;comment:应用镜像id"` // 应用镜像id
}

// AppStatus 应用状态
type AppStatus struct {
	ID          uint    `gorm:"primarykey" json:"-"`
	UserId      string  `gorm:"column:user_id;type:string;size:255;not null;comment:用户id"`             // 用户id
	AppId       string  `gorm:"column:app_id;type:string;size:255;default:'';comment:应用id"`            // 应用id
	AppName     string  `gorm:"column:app_name;type:string;size:255;default:'';comment:应用名称"`          // 应用名称
	FuncName    string  `gorm:"column:func_name;type:string;size:255;default:'空';comment:方法名称"`        // 方法名称
	FuncId      string  `gorm:"column:func_id;type:string;size:255;default:'';comment:方法id"`           // 应用id
	ContainerId string  `gorm:"column:container_id;type:string;size:255;not null;unique;comment:容器id"` // 容器id
	IP          string  `gorm:"column:ip;type:string;size:255;default:'';comment:应用容器宿主机ip"`           // 应用容器宿主机ip
	Port        string  `gorm:"column:port;type:string;size:255;default:'';comment:应用容器映射端口"`          // 应用容器映射端口
	Parameter   string  `gorm:"column:parameter;type:mediumtext;comment:输入的参数"`                        // 输入的参数
	CreateTime  []uint8 `gorm:"column:create_time;type:datetime;size:255;default:'';comment:容器创建时间"`   // 容器创建时间
	StartTime   []uint8 `gorm:"column:start_time;type:datetime;size:255;default:'';comment:容器启动时间"`    // 容器启动时间
	StopTime    []uint8 `gorm:"column:stop_time;type:datetime;size:255;default:'';comment:容器停止时间"`     // 容器停止时间
	Status      int     `gorm:"column:status;type:tinyint;default:0;comment:容器状态"`                     // 容器状态
	Step        int     `gorm:"column:step;type:tinyint;default:1;comment:方法步骤,不同步骤相当于单独的方法"`          // 方法步骤
}

// AppFuncInfo 应用方法中各步骤参数详情
type AppFuncInfo struct {
	ID                uint   `gorm:"primarykey" json:"-"`
	Step              int    `gorm:"column:step;type:tinyint;default:1;comment:方法步骤,不同步骤相当于单独的方法"`                // 方法步骤
	FuncId            string `gorm:"column:func_id;type:string;size:255;not null;unique;comment:方法id"`            // 方法id
	FuncName          string `gorm:"column:func_name;type:string;size:255;not null;comment:应用名称"`                 // 方法名称
	Cmd               string `gorm:"column:cmd;type:string;size:255;default:'';comment:应用执行命令"`                   // 应用执行命令
	ParameterFileName string `gorm:"column:parameter_file_name;type:string;size:255;default:'';comment:应用参数文件路径"` // 应用参数文件路径
	ParameterFilePath string `gorm:"column:parameter_file_path;type:string;size:255;default:'';comment:应用参数文件名称"` // 应用参数文件名称
	OutputPath        string `gorm:"column:output_path;type:string;size:255;default:'';comment:应用输出文件路径"`         // 应用输出文件路径
	Parameter         string `gorm:"column:parameter;type:text;comment:应用参数信息"`                                   // 应用参数信息
	Template          string `gorm:"column:template;type:text;comment:应用参数模板信息"`                                  // 应用参数模板信息
	OutputFilter      string `gorm:"column:output_filter;type:text;comment:输出文件过滤"`                               // 输出文件过滤正则,并匹配描述信息
}

// AppFuncMeta 应用中方法的信息,同一个应用的不同方法肯定在同一个镜像中
type AppFuncMeta struct {
	ID       uint   `gorm:"primarykey" json:"-"`
	FuncId   string `gorm:"column:func_id;type:string;size:255;not null;unique;comment:方法id"` // 方法id
	FuncName string `gorm:"column:func_name;type:string;size:255;default:'';comment:方法名称"`    // 方法名称
	AppId    string `gorm:"column:app_id;type:string;size:255;not null;comment:应用id"`         // 应用id
	AppName  string `gorm:"column:app_name;type:string;size:255;not null;comment:应用名称"`       // 应用名称
	StepNum  int    `gorm:"column:step_num;type:tinyint;default:1;comment:方法有多少步骤"`           // 方法步骤
}
