/*
数据表结构映射
*/

package tables

// StandardClassMeta 标准与规范分类元数据
type StandardClassMeta struct {
	ID                       uint   `gorm:"primarykey" json:"-"`
	StandardClassId          string `gorm:"column:standard_class_id;type:string;not null;unique;comment:标准类别id" json:"standardClassId"`
	StandardClassName        string `gorm:"column:standard_class_name;type:string;default:null;comment:标准类别名称" json:"standardClassName"`
	StandardClassDescription string `gorm:"column:standard_class_description;type:string;default:null;comment:标准类别描述" json:"standardClassDescription"`
}

// StandardMetaSimple 标准与规范元数据-简化版
type StandardMetaSimple struct {
	ID               uint   `gorm:"primarykey" json:"-"`
	StandardId       string `gorm:"column:standard_id;type:string;not null;unique;comment:标准id" json:"standardId"`
	StandardSerial   string `gorm:"column:standard_serial;type:string;default:'';comment:标准编号" json:"standardSerial"`
	StandardUnit     string `gorm:"column:standard_unit;type:string;default:'';comment:标准开发单位" json:"standardUnit"`
	StandardName     string `gorm:"column:standard_name;type:string;default:'';comment:标准名称" json:"standardName"`
	StandardURL      string `gorm:"column:standard_url;type:string;default:'';comment:标准URL" json:"standardURL"`
	StandardFileName string `gorm:"column:standard_file_name;type:string;default:'';comment:标准文件名称" json:"standardFileName"`
	StandardIsFile   bool   `gorm:"column:standard_is_file;type:bool;default:false;comment:标准是否为文件" json:"standardIsFile"`
	StandardClassId  string `gorm:"column:standard_class_id;type:string;not null;comment:标准层级" json:"standardClassId"`
}

// StandardMeta 标准与规范元数据
type StandardMeta struct {
	ID                                   uint   `gorm:"primarykey" json:"-"`
	StandardId                           string `gorm:"column:standard_id;type:string;not null;unique;comment:标准id" json:"standardId"`
	StandardName                         string `gorm:"column:standard_name;type:string;not null;comment:标准名称" json:"standardName"`
	StandardPicture                      string `gorm:"column:standard_picture;type:string;default:null;comment:文档缩略图" json:"standardPicture"`
	StandardUnit                         string `gorm:"column:standard_unit;type:string;default:null;comment:开发单位" json:"standardUnit"`
	StandardURL                          string `gorm:"column:standard_url;type:string;default:null;comment:标准URL" json:"standardURL"`
	StandardSerial                       string `gorm:"column:standard_serial;type:string;default:null;comment:标准编号" json:"standardSerial"`
	StandardClassId                      string `gorm:"column:standard_class_id;type:string;default:null;comment:标准层级" json:"standardClassId"`
	StandardProfessionalCategory         string `gorm:"column:standard_professional_category;type:string;default:null;comment:专业类别" json:"standardProfessionalCategory"`
	StandardChineseClassificationNumber  string `gorm:"column:standard_chinese_classification_number;type:string;default:null;comment:中国标准分类号" json:"standardChineseClassificationNumber"`
	StandardNationalClassificationNumber string `gorm:"column:standard_national_classification_number;type:string;default:null;comment:国际标准分类号" json:"standardNationalClassificationNumber"`
	StandardTechnologyCentralized        string `gorm:"column:standard_technology_centralized;type:string;default:null;comment:技术归口" json:"standardTechnologyCentralized"`
	StandardCompetentDepartment          string `gorm:"column:standard_competent_department;type:string;default:null;comment:主管部门" json:"standardCompetentDepartment"`
	StandardReleaseDate                  string `gorm:"column:standard_release_date;type:string;default:null;comment:发布日期" json:"standardReleaseDate"`
	StandardImplementationDate           string `gorm:"column:standard_implementation_date;type:string;default:null;comment:实施日期" json:"standardImplementationDate"`
	StandardExpirationDate               string `gorm:"column:standard_expiration_date;type:string;default:null;comment:废止日期" json:"standardExpirationDate"`
	StandardCurrentState                 string `gorm:"column:standard_current_state;type:string;default:null;comment:当前状态" json:"standardCurrentState"`
	StandardDraftingUnit                 string `gorm:"column:standard_drafting_unit;type:text;default:null;comment:起草单位" json:"standardDraftingUnit"`
	StandardDrafter                      string `gorm:"column:standard_drafter;type:text;default:null;comment:起草人" json:"standardDrafter"`
	StandardReferenceStandard            string `gorm:"column:standard_reference_standard;type:text;default:null;comment:引用标准" json:"standardReferenceStandard"`
}

// SharedClassMeta 共享资源
type SharedClassMeta struct {
	ID                     uint   `gorm:"primarykey" json:"-"`
	SharedClassId          string `gorm:"column:shared_class_id;type:string;not null;unique;comment:类别id" json:"sharedClassId"`
	SharedClassName        string `gorm:"column:shared_class_name;type:string;default:null;comment:类别名称" json:"sharedClassName"`
	SharedClassDescription string `gorm:"column:shared_class_description;type:string;default:null;comment:类别描述" json:"sharedClassDescription"`
}

// SharedLaboratoryMeta 共享资源
type SharedLaboratoryMeta struct {
	ID             uint   `gorm:"primarykey" json:"-"`
	LaboratoryId   string `gorm:"column:laboratory_id;type:string;not null;unique;comment:实验室id" json:"laboratoryId"`
	SharedClassId  string `gorm:"column:shared_class_id;type:string;not null;comment:类别id" json:"sharedClassId"`
	LaboratoryName string `gorm:"column:laboratory_name;type:string;default:null;comment:实验室名称" json:"laboratoryName"`
	LaboratoryType string `gorm:"column:laboratory_type;type:string;default:null;comment:实验室类型" json:"laboratoryType"`
	LaboratoryURL  string `gorm:"column:laboratory_url;type:string;default:null;comment:实验室url" json:"laboratoryURL"`
}

type NewsMeta struct {
	ID          uint   `gorm:"primarykey" json:"-"`
	NewsId      string `gorm:"column:news_id;type:string;not null;unique;comment:新闻id" json:"newsId"`
	NewsClass   string `gorm:"column:news_class;type:string;default:'';comment:新闻分类" json:"newsClass"`
	NewsTitle   string `gorm:"column:news_title;type:string;default:'';comment:新闻标题" json:"newsTitle"`
	NewsURL     string `gorm:"column:news_url;type:string;default:'';comment:新闻链接" json:"newsURL"`
	NewsContent string `gorm:"column:news_content;type:text;default:'';comment:新闻内容" json:"newsContent"`
	NewsPicture string `gorm:"column:news_picture;type:string;default:'';comment:新闻配图" json:"newsPicture"`
	CreateTime  string `gorm:"column:create_time;type:string;default:'';comment:创建时间" json:"createTime"`
	UpdateTime  string `gorm:"column:update_time;type:string;default:'';comment:创建时间" json:"updateTime"`
}

type ServiceCustomizedClassMeta struct {
	ID                      uint   `gorm:"primarykey" json:"-"`
	ServiceClassId          string `gorm:"column:service_class_id;type:string;not null;unique;comment:定制化服务分类id" json:"serviceClassId"`
	ServiceClassName        string `gorm:"column:service_class_name;type:string;default:'';comment:定制化服务分类" json:"serviceClassName"`
	ServiceClassDescription string `gorm:"column:service_class_description;type:string;default:'';comment:定制化服务分类描述" json:"serviceClassDescription"`
	ServiceClassPicture     string `gorm:"column:service_class_picture;type:string;default:'';comment:定制化服务分类缩略图" json:"serviceClassPicture"`
}

type ServiceCustomizedMeta struct {
	ID             uint   `gorm:"primarykey" json:"-"`
	ServiceId      string `gorm:"column:service_id;type:string;not null;unique;comment:定制化服务id" json:"serviceId"`
	ServiceName    string `gorm:"column:service_name;type:string;default:'';comment:定制化服务名称" json:"serviceName"`
	ServicePicture string `gorm:"column:service_picture;type:string;default:'';comment:定制化服务图片地址" json:"servicePicture"`
	ServiceClassId string `gorm:"column:service_class_id;type:string;default:'';comment:定制化服务分类id" json:"serviceClassId"`
}
