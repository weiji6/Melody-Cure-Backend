package request

// UploadImageRequest 上传图片请求
type UploadImageRequest struct {
	// 图片文件 (multipart/form-data)
	File interface{} `form:"file" binding:"required" swaggerignore:"true"`
	// 图片分类 (avatar, content, document等)
	Category string `form:"category" binding:"required" example:"avatar"`
	// 图片描述
	Description string `form:"description" example:"用户头像"`
}

// GetImageTokenRequest 获取图床token请求
type GetImageTokenRequest struct {
	// token用途 (upload, view, delete)
	Purpose string `json:"purpose" binding:"required" example:"upload"`
	// 有效期（分钟）
	ExpireMinutes int `json:"expire_minutes" binding:"min=1,max=1440" example:"60"`
}

// DeleteImageRequest 删除图片请求
type DeleteImageRequest struct {
	// 图片ID或文件名
	ImageID string `json:"image_id" binding:"required" example:"avatar_123456.jpg"`
	// 删除token
	Token string `json:"token" binding:"required" example:"delete_token_xxx"`
}