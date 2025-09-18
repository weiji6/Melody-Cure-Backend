package tool

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"math/rand"
	"net/smtp"
	"time"
)

type Mail struct {
	RedisClient *redis.Client
	From        string
	Key         string
}

func NewMail(redisClient *redis.Client, from string, key string) *Mail {
	return &Mail{
		RedisClient: redisClient,
		From:        from,
		Key:         key,
	}
}

// 生成验证码
func GenerateCode() string {
	rand.Seed(time.Now().UnixNano()) // 以纳米为级别
	code := rand.Intn(1000000)       // 生成6位数的验证码
	return fmt.Sprintf("%06d", code)
}

// SendEmail 发送邮件函数
func (m *Mail) SendEmail(to string, code string) error {
	from := m.From
	password := m.Key // 邮箱授权码
	smtpServer := "smtp.qq.com:465"

	// 设置 PlainAuth
	auth := smtp.PlainAuth("", from, password, "smtp.qq.com")

	// 创建 tls 配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}

	// 连接到 SMTP 服务器
	conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
	if err != nil {
		return fmt.Errorf("TLS 连接失败: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.qq.com")
	if err != nil {
		return fmt.Errorf("SMTP 客户端创建失败: %v", err)
	}
	defer client.Quit()

	// 使用 auth 进行认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("认证失败: %v", err)
	}

	// 设置发件人和收件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("发件人设置失败: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("收件人设置失败: %v", err)
	}

	// 写入邮件内容
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("数据写入失败: %v", err)
	}
	defer wc.Close()

	subject := "Lush-And-Verdant"
	body := `
		<h1>Verification Code</h1>
		<p>Your verification code is: <strong>` + code + `</strong></p >
		<p>This verification code is valid for 5 minutes</p >
		<p>If you are not doing it yourself, please ignore it !</p >
		<h1>验证码</h1>
		<p>你的验证码是: <strong>` + code + `</strong></p >
		<p>验证码的有效时间是5分钟。</p >
		<p>如非本人操作，请忽略此邮件！</p >
	`
	msg := []byte("From: Sender Name <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	_, err = wc.Write(msg)
	if err != nil {
		return fmt.Errorf("消息发送失败: %v", err)
	}

	return m.StoreCodeInRedis(to, code)
}

// StoreCodeInRedis 存储验证码到Redis
func (m *Mail) StoreCodeInRedis(email string, code string) error {
	ctx := context.Background()
	key := email
	expiration := 5 * time.Minute // 设置验证码有效期为5分钟

	err := m.RedisClient.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return fmt.Errorf("存储验证码到Redis失败:%v", err)
	}
	return nil
}

// VerifyCode 校验验证码
func (m *Mail) VerifyCode(email string, code string) (bool, error) {
	ctx := context.Background()
	key := email

	// 从Redis获取验证码
	storedCode, err := m.RedisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, fmt.Errorf("验证码不存在或已过期")
	} else if err != nil {
		return false, fmt.Errorf("查询Redis失败:%v", err)
	}

	if storedCode != code {
		return false, fmt.Errorf("验证码错误")
	}

	// 删除Redis中的验证码，防止重复使用
	err = m.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return false, fmt.Errorf("删除验证码失败:%v", err)
	}

	return true, nil
}

// 修改验证码状态
func (m *Mail) ChangeStatus(addr string) error {
	ctx := context.Background()
	key := addr

	// 删除Redis中的验证码
	err := m.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("删除验证码失败:%v", err)
	}

	return nil
}
