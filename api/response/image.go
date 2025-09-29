package response

import "time"

// UploadImageResponse 上传图片响应
type UploadImageResponse struct {
	// 图片ID
	ImageID string `json:"image_id" example:"img_123456789"`
	// 图片URL
	ImageURL string `json:"image_url" example:"https://example.com/images/avatar_123456.jpg"`
	// 图片文件名
	FileName string `json:"file_name" example:"avatar_123456.jpg"`
	// 图片大小（字节）
	FileSize int64 `json:"file_size" example:"1024000"`
	// 图片类型
	ContentType string `json:"content_type" example:"image/jpeg"`
	// 图片分类
	Category string `json:"category" example:"avatar"`
	// 上传时间
	UploadTime time.Time `json:"upload_time" example:"2024-01-01T12:00:00Z"`
}

// ImageTokenResponse 图床token响应
type ImageTokenResponse struct {
	// 访问token
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// token类型
	TokenType string `json:"token_type" example:"Bearer"`
	// 过期时间
	ExpiresAt time.Time `json:"expires_at" example:"2024-01-01T13:00:00Z"`
	// token用途
	Purpose string `json:"purpose" example:"upload"`
	// 剩余有效时间（秒）
	ExpiresIn int64 `json:"expires_in" example:"3600"`
}

// ImageInfoResponse 图片信息响应
type ImageInfoResponse struct {
	// 图片ID
	ImageID string `json:"image_id" example:"img_123456789"`
	// 图片URL
	ImageURL string `json:"image_url" example:"https://example.com/images/avatar_123456.jpg"`
	// 图片文件名
	FileName string `json:"file_name" example:"avatar_123456.jpg"`
	// 图片大小（字节）
	FileSize int64 `json:"file_size" example:"1024000"`
	// 图片类型
	ContentType string `json:"content_type" example:"image/jpeg"`
	// 图片分类
	Category string `json:"category" example:"avatar"`
	// 图片描述
	Description string `json:"description" example:"用户头像"`
	// 上传时间
	UploadTime time.Time `json:"upload_time" example:"2024-01-01T12:00:00Z"`
	// 上传用户ID
	UserID uint `json:"user_id" example:"1"`
}

// ImageListResponse 图片列表响应
type ImageListResponse struct {
	// 图片列表
	Images []ImageInfoResponse `json:"images"`
	// 总数
	Total int64 `json:"total" example:"100"`
	// 当前页
	Page int `json:"page" example:"1"`
	// 每页数量
	PageSize int `json:"page_size" example:"10"`
}