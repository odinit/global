package tables

// DockerImage 镜像信息
type DockerImage struct {
	ID              uint   `gorm:"primarykey" json:"-"`
	ImageName       string `gorm:"column:image_name;type:string;not null;unique;comment:镜像名称" json:"imageName"`                                                 // 镜像名称
	Hostname        string `gorm:"column:hostname;type:string;default:'';comment:主机名，默认为空"`                                                                     // 主机名，默认为空
	DomainName      string `gorm:"column:domain_name;type:string;default:'';comment:域名，默认为空"`                                                                   // 域名，默认为空
	User            string `gorm:"column:user;type:string;default:root;comment:用户名，默认为root"`                                                                    // 用户名，默认为root
	ExposedPorts    string `gorm:"column:exposed_ports;type:string;default:'';comment:暴露的端口，使用;分割，默认为空"`                                                        // 暴露的端口，使用;分割，默认为空
	Tty             bool   `gorm:"column:tty;type:tinyint;size:1;default:1;comment:伪终端，默认为1，即true"`                                                             // 伪终端，默认为1，即true
	OpenStdin       bool   `gorm:"column:open_stdin;type:tinyint;size:1;default:1;comment:打开标准输入，默认1，即true"`                                                    // 打开标准输入，默认1，即true
	StdinOnce       bool   `gorm:"column:stdin_once;type:tinyint;size:1;default:1;comment:当1个连接的客户端断开时关闭标准输入，默认1，即true"`                                        // 当1个连接的客户端断开时关闭标准输入，默认1，即true
	Env             string `gorm:"column:env;type:string;default:'';comment:环境变量，使用;分割，默认为空"`                                                                   // 环境变量，使用;分割，默认为空
	Cmd             string `gorm:"column:cmd;type:string;default:'';comment:启动命令，使用;分割，默认为空"`                                                                   // 启动命令，使用;分割，默认为空
	HealthCheck     string `gorm:"column:health_check;type:string;default:'';comment:健康检查，json字符串，默认为空"`                                                        // 健康检查，json字符串，默认为空
	Volumes         string `gorm:"column:volumes;type:string;default:'';comment:挂载列表，json字符串-字典，默认为空"`                                                          // 挂载列表，json字符串-字典，默认为空
	WorkingDir      string `gorm:"column:working_dir;type:string;default:'';comment:工作路径，默认为空"`                                                                 // 工作路径，默认为空
	Entrypoint      string `gorm:"column:entrypoint;type:string;default:'';comment:启动入口，使用;分割，默认为空"`                                                            // 启动入口，使用;分割，默认为空
	NetworkDisabled bool   `gorm:"column:network_disabled;type:tinyint;size:1;default:0;comment:是否禁用网络，默认0"`                                                    // 是否禁用网络，默认0
	MacAddress      string `gorm:"column:mac_addr;type:string;default:'';comment:mac地址，默认为空"`                                                                   // mac地址，默认为空
	OnBuild         string `gorm:"column:on_build;type:string;default:'';comment:Dockerfile中ONBUILD源数据，使用;分割，默认为空"`                                             // Dockerfile中ONBUILD源数据，使用;分割，默认为空
	Labels          string `gorm:"column:labels;type:string;default:'';comment:给容器设置标签，json字符串-字典，默认为空"`                                                        // 给容器设置标签，json字符串-字典，默认为空
	StopSignal      string `gorm:"column:stop_signal;type:string;default:'';comment:停止容器的信号，默认为空"`                                                              // 停止容器的信号，默认为空
	StopTimeout     int    `gorm:"column:stop_timeout;type:int;default:0;comment:容器超时停止时间，默认为null"`                                                             // 容器超时停止时间，默认为null
	Shell           string `gorm:"column:shell;type:string;default:'';comment:shell，使用;分割，默认为空"`                                                                // shell，使用;分割，默认为空
	Binds           string `gorm:"column:binds;type:string;default:'';comment:挂载路径列表，使用;分割，默认为空"`                                                               // 挂载路径列表，使用;分割，默认为空
	ContainerIDFile string `gorm:"column:container_id_file;type:string;default:'';comment:写入容器id的文件路径，默认为空"`                                                    // 写入容器id的文件路径，默认为空
	LogConfig       string `gorm:"column:log_config;type:string;default:'';comment:容器日志配置，json字符串-字典，默认为空"`                                                     // 容器日志配置，json字符串-字典，默认为空
	NetworkMode     string `gorm:"column:network_mode;type:string;default:'';comment:容器网络类型，默认为空"`                                                              // 容器网络类型，默认为空
	PortBindings    string `gorm:"column:port_bindings;type:string;default:'';comment:端口映射，形如8000:8000，使用;隔开，默认为空"`                                             // 端口映射，形如8000:8000，使用;隔开，默认为空
	RestartPolicy   string `gorm:"column:restart_policy;type:string;default:'';comment:重启策略，json字符串-字典，默认为空"`                                                   // 重启策略，json字符串-字典，默认为空
	AutoRemove      bool   `gorm:"column:auto_remove;type:tinyint;size:1;default:0;comment:自动删除，默认为0，即false"`                                                   // 自动删除，默认为0，即false
	VolumeDriver    string `gorm:"column:volume_driver;type:string;default:'';comment:挂载驱动名称，默认为空"`                                                             // 挂载驱动名称，默认为空
	VolumesFrom     string `gorm:"column:volumes_from;type:string;default:'';comment: 从其他容器挂载，使用;分割，默认为空"`                                                      // 从其他容器挂载，使用;分割，默认为空
	CapAdd          string `gorm:"column:cap_add;type:string;default:'';comment:添加到容器内核功能列表，使用;分割，默认为空"`                                                        // 添加到容器内核功能列表，使用;分割，默认为空
	CapDrop         string `gorm:"column:cap_drop;type:string;default:'';comment:从容器内核功能列表中删除的内核，使用;分割，默认为空"`                                                   // 从容器内核功能列表中删除的内核，使用;分割，默认为空
	CgroupnsMode    string `gorm:"column:cgroup_ns_mode;type:string;default:'';comment:容器的 Cgroup 命名空间模式，默认为空"`                                                 // 容器的 Cgroup 命名空间模式，默认为空
	DNS             string `gorm:"column:dns;type:string;default:'';comment:DNS服务器列表，使用;分割，默认为空"`                                                               // DNS服务器列表，使用;分割，默认为空
	DNSOptions      string `gorm:"column:dns_ptions;type:string;default:'';comment:DNSOptions列表，使用;分割，默认为空"`                                                    // DNSOptions列表，使用;分割，默认为空
	DNSSearch       string `gorm:"column:dns_search;type:string;default:'';comment:DNSSearch列表，使用;分割，默认为空"`                                                     // DNSSearch列表，使用;分割，默认为空
	ExtraHosts      string `gorm:"column:extra_hosts;type:string;default:'';comment:额外主机列表，使用;分割，默认为空"`                                                         // 额外主机列表，使用;分割，默认为空
	GroupAdd        string `gorm:"column:group_add;type:string;default:'';comment:List of additional groups that the container process will run as，使用;分割，默认为空"` // List of additional groups that the container process will run as，使用;分割，默认为空
	IpcMode         string `gorm:"column:ipc_mode;type:string;default:'';comment:用于容器的 IPC 命名空间，默认为空"`                                                          // 用于容器的 IPC 命名空间，默认为空
	Cgroup          string `gorm:"column:cgroup;type:string;default:'';comment:容器的Cgroup，默认为空"`                                                                 // 容器的Cgroup，默认为空
	Links           string `gorm:"column:links;type:string;default:'';comment:连接列表(名称:别名)，使用;分割，默认为空"`                                                          // 连接列表(名称:别名)，使用;分割，默认为空
	OomScoreAdj     int    `gorm:"column:oom_score_adj;type:int;default:0;comment:容器OOM-killing首选项，默认0"`                                                        // 容器OOM-killing首选项，默认0
	PidMode         string `gorm:"column:pid_mode;type:string;default:'';comment:pid命名空间，默认为空"`                                                                 // pid命名空间，默认为空
	Privileged      bool   `gorm:"column:privileged;type:tinyint;size:1;default:0;comment:容器是否处于特权模式，默认为0，即false"`                                              // 容器是否处于特权模式，默认为0，即false
	PublishAllPorts bool   `gorm:"column:publish_all_ports;type:tinyint;size:1;default:0;comment:是否发布容器所有暴露的端口，默认为0，即false"`                                    // 是否发布容器所有暴露的端口，默认为0，即false
	ReadonlyRootfs  bool   `gorm:"column:readonly_rootfs;type:tinyint;size:1;default:0;comment:容器根目录是否只读，默认为0，即false"`                                          // 容器根目录是否只读，默认为0，即false
	SecurityOpt     string `gorm:"column:security_opt;type:string;default:'';comment:用于为 MLS 系统（例如 SELinux）自定义标签的字符串值列表，使用;分割，默认为空"`                            // 用于为 MLS 系统（例如 SELinux）自定义标签的字符串值列表，使用;分割，默认为空
	StorageOpt      string `gorm:"column:storage_opt;type:string;default:'';comment:每个容器的存储驱动程序选项，json字符串-map，默认为空"`                                            // 每个容器的存储驱动程序选项，json字符串-map，默认为空
	Tmpfs           string `gorm:"column:tmpfs;type:string;default:'';comment:用于容器的 tmpfs（挂载）列表，json字符串-map，默认为空"`                                              // 用于容器的 tmpfs（挂载）列表，json字符串-map，默认为空
	UTSMode         string `gorm:"column:uts_mode;type:string;default:'';comment:UTS命名空间，默认为空"`                                                                 // UTS命名空间，默认为空
	UsernsMode      string `gorm:"column:user_ns_mode;type:string;default:'';comment:用户命名空间，默认为空"`                                                              // 用户命名空间，默认为空
	ShmSize         int64  `gorm:"column:shm_size;type:bigint;default:0;comment:shm 内存使用总量，默认为0"`                                                               // shm 内存使用总量，默认为0
	Sysctls         string `gorm:"column:sysctls;type:string;default:'';comment:sysctls命名空间，json字符串-map，默认为空"`                                                  // sysctls命名空间，json字符串-map，默认为空
	Runtime         string `gorm:"column:runtime;type:string;default:'';comment:Runtime，默认为空"`                                                                  // Runtime，默认为空
	Resources       string `gorm:"column:resources;type:string;default:'';comment:容器资源配置，json字符串-map，默认为空"`                                                     // 容器资源配置，json字符串-map，默认为空
	Mounts          string `gorm:"column:mounts;type:text;default:'';comment:容器挂载路径，使用;分割，默认为空"`                                                                // 容器挂载路径，使用;分割，默认为空
	MaskedPaths     string `gorm:"column:masked_paths;type:string;default:'';comment:要在容器内屏蔽的路径列表（这会覆盖默认路径集），使用;分割，默认为空"`                                       // 要在容器内屏蔽的路径列表（这会覆盖默认路径集），使用;分割，默认为空
	ReadonlyPaths   string `gorm:"column:readonly_paths;type:string;default:'';comment:只读路径列表（会覆盖默认路径集），使用;分割，默认为空"`                                            // 只读路径列表（会覆盖默认路径集），使用;分割，默认为空
}
