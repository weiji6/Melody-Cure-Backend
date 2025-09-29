package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"melody_cure/config"
	"time"
)

type ImageService struct {
	qiniuConfig config.QiniuConfig
}

// NewImageService 创建ImageService实例
func NewImageService() *ImageService {
	return &ImageService{
		qiniuConfig: config.GetQiniuConfig(),
	}
}

// QiniuUploadToken 七牛云上传token响应结构
type QiniuUploadToken struct {
	Token      string `json:"token"`
	Domain     string `json:"domain"`
	Bucket     string `json:"bucket"`
	ExpiresAt  int64  `json:"expires_at"`
	UseHTTPS   bool   `json:"use_https"`
}

// GenerateQiniuUploadToken 生成七牛云上传token
func (s *ImageService) GenerateQiniuUploadToken() (*QiniuUploadToken, error) {
	// 设置过期时间
	deadline := time.Now().Unix() + s.qiniuConfig.Expires
	
	// 构建上传策略
	putPolicy := map[string]interface{}{
		"scope":    s.qiniuConfig.Bucket,
		"deadline": deadline,
		// 可以添加更多策略配置
		"returnBody": `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	
	// 将策略转换为JSON
	putPolicyJSON, err := json.Marshal(putPolicy)
	if err != nil {
		return nil, err
	}
	
	// Base64编码策略
	encodedPutPolicy := base64.URLEncoding.EncodeToString(putPolicyJSON)
	
	// 使用HMAC-SHA1签名
	h := hmac.New(sha1.New, []byte(s.qiniuConfig.SecretKey))
	h.Write([]byte(encodedPutPolicy))
	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	
	// 生成最终token
	token := s.qiniuConfig.AccessKey + ":" + sign + ":" + encodedPutPolicy
	
	return &QiniuUploadToken{
		Token:     token,
		Domain:    s.qiniuConfig.Domain,
		Bucket:    s.qiniuConfig.Bucket,
		ExpiresAt: deadline,
		UseHTTPS:  s.qiniuConfig.UseHTTPS,
	}, nil
}