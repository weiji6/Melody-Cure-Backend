package DAO

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	Image         string         `json:"image"`
	Name          string         `json:"name"`
	Password      string         `json:"-"` // 不返回密码
	Email         string         `gorm:"unique" json:"email"`
	Phone         string         `json:"phone"`
	Identity      string         `json:"identity"` // 身份类型
	Address       string         `json:"address"`
	Certificate   string         `json:"certificate"`
	Certification bool           `json:"certification"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// 机构/康复师认证信息
type Certification struct {
	ID               string         `gorm:"primaryKey" json:"id"`
	UserID           string         `gorm:"index" json:"user_id"`
	CertificateType  string         `json:"certificate_type"` // 机构认证/康复师认证
	CertificateName  string         `json:"certificate_name"`
	CertificateNo    string         `json:"certificate_no"`
	IssuingAuthority string         `json:"issuing_authority"`
	IssueDate        time.Time      `json:"issue_date"`
	ExpiryDate       *time.Time     `json:"expiry_date"`
	Status           string         `json:"status"` // pending, approved, rejected
	ReviewNotes      string         `json:"review_notes"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	User             User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// 陪伴型AI配置
type AICompanion struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	UserID       string         `gorm:"index" json:"user_id"`
	CompanionType string        `json:"companion_type"` // 类似抖音小火人/王者灵主
	Name         string         `json:"name"`
	Avatar       string         `json:"avatar"`
	Personality  string         `json:"personality"` // 性格设定
	VoiceType    string         `json:"voice_type"`
	IsActive     bool           `json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	User         User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// 虚拟疗愈导师
type VirtualTherapist struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	UserID       string         `gorm:"index" json:"user_id"`
	TherapistType string        `json:"therapist_type"` // 专业领域
	Name         string         `json:"name"`
	Avatar       string         `json:"avatar"`
	Specialization string       `json:"specialization"` // 专业特长
	Experience   int            `json:"experience"` // 经验年限
	IsActive     bool           `json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	User         User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// 儿童档案
type ChildArchive struct {
	ID              string         `gorm:"primaryKey" json:"id"`
	UserID          string         `gorm:"index" json:"user_id"` // 家长ID
	ChildName       string         `json:"child_name"`
	Gender          string         `json:"gender"`
	BirthDate       time.Time      `json:"birth_date"`
	Avatar          string         `json:"avatar"`
	Condition       string         `json:"condition"` // 病情描述
	Diagnosis       string         `json:"diagnosis"` // 诊断结果
	Treatment       string         `json:"treatment"` // 治疗方案
	Progress        string         `json:"progress"` // 康复进展
	Notes           string         `json:"notes"` // 备注
	TreatmentStartDate *time.Time  `json:"treatment_start_date"` // 治疗开始日期
	HealedDays      int            `json:"healed_days"` // 已疗愈天数
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// 课程
type Course struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Level       string         `json:"level"` // 初级/中级/高级
	Duration    int            `json:"duration"` // 课程时长(分钟)
	CoverImage  string         `json:"cover_image"`
	VideoURL    string         `json:"video_url"`
	Content     string         `json:"content"` // 课程内容
	Price       float64        `json:"price"`
	IsFree      bool           `json:"is_free"`
	Status      string         `json:"status"` // published, draft
	ViewCount   int            `json:"view_count"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// 用户收藏
type UserFavorite struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	UserID       string         `gorm:"index" json:"user_id"`
	ResourceType string         `json:"resource_type"` // course, game, article
	ResourceID   string         `json:"resource_id"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	User         User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// 游戏
type Game struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"` // 认知训练/感统训练等
	AgeRange    string         `json:"age_range"` // 适用年龄段
	Difficulty  string         `json:"difficulty"` // 难度等级
	CoverImage  string         `json:"cover_image"`
	GameURL     string         `json:"game_url"`
	Instructions string        `json:"instructions"` // 游戏说明
	Benefits    string         `json:"benefits"` // 训练效果
	PlayCount   int            `json:"play_count"`
	Rating      float64        `json:"rating"`
	Status      string         `json:"status"` // published, draft
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
