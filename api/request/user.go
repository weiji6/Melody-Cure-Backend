package request

import "time"

type Register struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Identity string `json:"identity"`
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfile struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Image   string `json:"image"`
}

type UpdateProfileRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Image   string `json:"image"`
}

type CertificationRequest struct {
	CertificateType  string    `json:"certificate_type" binding:"required"` // 机构认证/康复师认证
	CertificateName  string    `json:"certificate_name" binding:"required"`
	CertificateNo    string    `json:"certificate_no" binding:"required"`
	IssuingAuthority string    `json:"issuing_authority" binding:"required"`
	IssueDate        time.Time `json:"issue_date" binding:"required"`
	ExpiryDate       *time.Time `json:"expiry_date"`
}

type AICompanionRequest struct {
	CompanionType string `json:"companion_type" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Avatar       string `json:"avatar"`
	Personality  string `json:"personality"`
	VoiceType    string `json:"voice_type"`
}

type VirtualTherapistRequest struct {
	TherapistType  string `json:"therapist_type" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Avatar        string `json:"avatar"`
	Specialization string `json:"specialization"`
	Experience    int    `json:"experience"`
}

type ChildArchiveRequest struct {
	ChildName          string     `json:"child_name" binding:"required"`
	Gender             string     `json:"gender" binding:"required"`
	BirthDate          time.Time  `json:"birth_date" binding:"required"`
	Avatar             string     `json:"avatar"`
	Condition          string     `json:"condition"`
	Diagnosis          string     `json:"diagnosis"`
	Treatment          string     `json:"treatment"`
	Progress           string     `json:"progress"`
	Notes              string     `json:"notes"`
	TreatmentStartDate *time.Time `json:"treatment_start_date"`
}

type FavoriteRequest struct {
	ResourceType string `json:"resource_type" binding:"required"` // course, game, article
	ResourceID   string `json:"resource_id" binding:"required"`
}
