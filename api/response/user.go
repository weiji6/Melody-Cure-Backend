package response

import (
	"melody_cure/DAO"
	"time"
)

// 用户信息响应
type UserProfile struct {
	ID            string    `json:"id"`
	Image         string    `json:"image"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Identity      string    `json:"identity"`
	Address       string    `json:"address"`
	Certificate   string    `json:"certificate"`
	Certification bool      `json:"certification"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// 认证状态响应
type CertificationStatus struct {
	ID               string     `json:"id"`
	CertificateType  string     `json:"certificate_type"`
	CertificateName  string     `json:"certificate_name"`
	CertificateNo    string     `json:"certificate_no"`
	IssuingAuthority string     `json:"issuing_authority"`
	IssueDate        time.Time  `json:"issue_date"`
	ExpiryDate       *time.Time `json:"expiry_date"`
	Status           string     `json:"status"`
	ReviewNotes      string     `json:"review_notes"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// AI陪伴响应
type AICompanionResponse struct {
	ID           string    `json:"id"`
	CompanionType string   `json:"companion_type"`
	Name         string    `json:"name"`
	Avatar       string    `json:"avatar"`
	Personality  string    `json:"personality"`
	VoiceType    string    `json:"voice_type"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

// 虚拟疗愈导师响应
type VirtualTherapistResponse struct {
	ID           string    `json:"id"`
	TherapistType string   `json:"therapist_type"`
	Name         string    `json:"name"`
	Avatar       string    `json:"avatar"`
	Specialization string  `json:"specialization"`
	Experience   int       `json:"experience"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

// 儿童档案响应
type ChildArchiveResponse struct {
	ID                 string     `json:"id"`
	ChildName          string     `json:"child_name"`
	Gender             string     `json:"gender"`
	BirthDate          time.Time  `json:"birth_date"`
	Age                int        `json:"age"` // 计算得出的年龄
	Avatar             string     `json:"avatar"`
	Condition          string     `json:"condition"`
	Diagnosis          string     `json:"diagnosis"`
	Treatment          string     `json:"treatment"`
	Progress           string     `json:"progress"`
	Notes              string     `json:"notes"`
	TreatmentStartDate *time.Time `json:"treatment_start_date"`
	HealedDays         int        `json:"healed_days"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

// 收藏响应
type FavoriteResponse struct {
	ID           string    `json:"id"`
	ResourceType string    `json:"resource_type"`
	ResourceID   string    `json:"resource_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// 课程响应
type CourseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Level       string    `json:"level"`
	Duration    int       `json:"duration"`
	CoverImage  string    `json:"cover_image"`
	VideoURL    string    `json:"video_url"`
	Content     string    `json:"content"`
	Price       float64   `json:"price"`
	IsFree      bool      `json:"is_free"`
	Status      string    `json:"status"`
	ViewCount   int       `json:"view_count"`
	CreatedAt   time.Time `json:"created_at"`
}

// 游戏响应
type GameResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	AgeRange    string    `json:"age_range"`
	Difficulty  string    `json:"difficulty"`
	CoverImage  string    `json:"cover_image"`
	GameURL     string    `json:"game_url"`
	Instructions string   `json:"instructions"`
	Benefits    string    `json:"benefits"`
	PlayCount   int       `json:"play_count"`
	Rating      float64   `json:"rating"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// 转换函数
func ToUserProfile(user *DAO.User) UserProfile {
	return UserProfile{
		ID:            user.ID,
		Image:         user.Image,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Identity:      user.Identity,
		Address:       user.Address,
		Certificate:   user.Certificate,
		Certification: user.Certification,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

// 转换儿童档案到响应结构体
func ToChildArchiveResponse(archive *DAO.ChildArchive) ChildArchiveResponse {
	// 计算年龄
	age := calculateAge(archive.BirthDate)
	
	// 计算已疗愈天数
	healedDays := archive.HealedDays
	if archive.TreatmentStartDate != nil {
		healedDays = int(time.Since(*archive.TreatmentStartDate).Hours() / 24)
	}
	
	return ChildArchiveResponse{
		ID:                 archive.ID,
		ChildName:          archive.ChildName,
		Gender:             archive.Gender,
		BirthDate:          archive.BirthDate,
		Age:                age,
		Avatar:             archive.Avatar,
		Condition:          archive.Condition,
		Diagnosis:          archive.Diagnosis,
		Treatment:          archive.Treatment,
		Progress:           archive.Progress,
		Notes:              archive.Notes,
		TreatmentStartDate: archive.TreatmentStartDate,
		HealedDays:         healedDays,
		CreatedAt:          archive.CreatedAt,
		UpdatedAt:          archive.UpdatedAt,
	}
}

// 计算年龄的辅助函数
func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	
	// 如果还没到生日，年龄减1
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	
	return age
}

func ToCertificationStatus(cert *DAO.Certification) CertificationStatus {
	return CertificationStatus{
		ID:               cert.ID,
		CertificateType:  cert.CertificateType,
		CertificateName:  cert.CertificateName,
		CertificateNo:    cert.CertificateNo,
		IssuingAuthority: cert.IssuingAuthority,
		IssueDate:        cert.IssueDate,
		ExpiryDate:       cert.ExpiryDate,
		Status:           cert.Status,
		ReviewNotes:      cert.ReviewNotes,
		CreatedAt:        cert.CreatedAt,
		UpdatedAt:        cert.UpdatedAt,
	}
}