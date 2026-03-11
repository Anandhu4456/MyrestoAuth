package smtp

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type SMTPConfig struct {
	Host        string
	Port        int
	SenderEmail string
	AppPassword string
	FromName    string
	BaseURL     string
}

type SMTPService struct {
	smtpC SMTPConfig
}

func NewSMTPService(smtpc SMTPConfig) *SMTPService {
	return &SMTPService{
		smtpC: smtpc,
	}
}

func (s *SMTPService) SendVerificationEmail(to, restaurantName, token string) error {

	link := fmt.Sprintf("%s/verify-email?token=%s", s.smtpC.BaseURL, token)

	body, err := render(verificationTmpl, map[string]string{
		"RestaurantName": restaurantName,
		"Link":           link,
	})
	if err != nil {
		return err
	}

	return s.send(to, "Verify Your MyRestoToday Account", body)

}

func (s *SMTPService) SendPasswordSetupEmail(to, restaurantName, token string) error {
	link := fmt.Sprintf("%s/set-password?token=%s", s.smtpC.BaseURL, token)
	body, err := render(passwordSetupTmpl, map[string]string{
		"RestaurantName": restaurantName,
		"Link":           link,
	})
	if err != nil {
		return err
	}
	return s.send(to, "Set up your MyRestoToday password", body)
}

func (s *SMTPService) send(to, subject, htmlBody string) error {
	auth := smtp.PlainAuth("", s.smtpC.SenderEmail, s.smtpC.AppPassword, s.smtpC.Host)
	from := fmt.Sprintf("%s <%s>", s.smtpC.FromName, s.smtpC.SenderEmail)
	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		from, to, subject, htmlBody,
	))
	return smtp.SendMail(fmt.Sprintf("%s:%d", s.smtpC.Host, s.smtpC.Port), auth, s.smtpC.SenderEmail, []string{to}, msg)
}

func render(tmplStr string, data map[string]string) (string, error) {
	t, err := template.New("email").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

const verificationTmpl = `<!DOCTYPE html><html><head><meta charset="UTF-8">
<style>body{font-family:Arial,sans-serif;background:#f4f4f4}.container{max-width:600px;margin:40px auto;background:#fff;border-radius:8px;overflow:hidden;box-shadow:0 2px 8px rgba(0,0,0,.1)}.header{background:#E74C3C;padding:30px;text-align:center}.header h1{color:#fff;margin:0}.body{padding:30px}.body p{color:#555;line-height:1.6}.btn{display:inline-block;margin:20px 0;padding:14px 28px;background:#E74C3C;color:#fff;text-decoration:none;border-radius:5px;font-weight:700}.footer{padding:20px;text-align:center;font-size:12px;color:#aaa}</style>
</head><body><div class="container">
<div class="header"><h1>MyRestoToday</h1></div>
<div class="body"><h2>Welcome, {{.RestaurantName}}!</h2>
<p>Verify your email address to continue setting up your account.</p>
<a href="{{.Link}}" class="btn">Verify Email Address</a>
<p>This link expires in 24 hours.</p></div>
<div class="footer">© 2024 MyRestoToday</div></div></body></html>`

const passwordSetupTmpl = `<!DOCTYPE html><html><head><meta charset="UTF-8">
<style>body{font-family:Arial,sans-serif;background:#f4f4f4}.container{max-width:600px;margin:40px auto;background:#fff;border-radius:8px;overflow:hidden;box-shadow:0 2px 8px rgba(0,0,0,.1)}.header{background:#E74C3C;padding:30px;text-align:center}.header h1{color:#fff;margin:0}.body{padding:30px}.body p{color:#555;line-height:1.6}.btn{display:inline-block;margin:20px 0;padding:14px 28px;background:#E74C3C;color:#fff;text-decoration:none;border-radius:5px;font-weight:700}.footer{padding:20px;text-align:center;font-size:12px;color:#aaa}</style>
</head><body><div class="container">
<div class="header"><h1>MyRestoToday</h1></div>
<div class="body"><h2>Set Your Password, {{.RestaurantName}}!</h2>
<p>Your email is verified. Click below to set your password.</p>
<a href="{{.Link}}" class="btn">Set Password</a>
<p>This link expires in 24 hours.</p></div>
<div class="footer">© 2024 MyRestoToday</div></div></body></html>`
