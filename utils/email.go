package utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
}

type EmailData struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(config EmailConfig, emailData EmailData) error {
	// Setup authentication
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)

	// Compose the email
	to := []string{emailData.To}
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		emailData.To, emailData.Subject, emailData.Body))

	// Send the email
	err := smtp.SendMail(config.SMTPHost+":"+config.SMTPPort, auth, config.FromEmail, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func GenerateResetPasswordEmail(resetLink string) string {
	return `<html lang="en" style="margin:0;padding:0;">
  <body style="margin:0;padding:0;background:#f6f9fc;">
    <!-- Preheader (preview text) -->
    <div style="display:none;max-height:0;overflow:hidden;opacity:0;">
      Reset your password. This link expires in 1 hour.
    </div>

    <table role="presentation" width="100%" cellspacing="0" cellpadding="0" border="0" style="background:#f6f9fc;padding:24px 0;">
      <tr>
        <td align="center">
          <table role="presentation" width="600" cellspacing="0" cellpadding="0" border="0" style="width:600px;max-width:100%;background:#ffffff;border-radius:12px;box-shadow:0 4px 16px rgba(0,0,0,0.05);overflow:hidden;">
            <!-- Header -->
            <tr>
              <td style="padding:20px 28px;background:#111827;">
                <h1 style="margin:0;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:18px;line-height:24px;color:#ffffff;">
                  APP_NAME
                </h1>
              </td>
            </tr>

            <!-- Body -->
            <tr>
              <td style="padding:28px;">
                <h2 style="margin:0 0 8px 0;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:22px;line-height:30px;color:#111827;">
                  Reset Your Password
                </h2>
                <p style="margin:0 0 16px 0;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:14px;line-height:22px;color:#374151;">
                  You recently requested to reset your password for your APP_NAME account. Click the button below to set a new password.
                </p>

                <!-- Button -->
                <table role="presentation" cellspacing="0" cellpadding="0" border="0" style="margin:0 0 16px 0;">
                  <tr>
                    <td>
                      <a href="` + resetLink + `" target="_blank"
                         style="display:inline-block;background:#4F46E5;color:#ffffff;text-decoration:none;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:14px;line-height:20px;font-weight:600;padding:12px 20px;border-radius:8px;">
                        Reset Password
                      </a>
                    </td>
                  </tr>
                </table>

                <p style="margin:0 0 8px 0;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:12px;line-height:20px;color:#6B7280;">
                  This link will expire in <strong>1 hour</strong> for your security.
                </p>
                <p style="margin:0 0 0 0;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:12px;line-height:20px;color:#6B7280;">
                  If you did not request a password reset, you can safely ignore this email.
                </p>
              </td>
            </tr>

            <!-- Footer -->
            <tr>
              <td style="padding:16px 28px;background:#F3F4F6;">
                <p style="margin:0;font-family:Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:11px;line-height:18px;color:#6B7280;">
                  © 2025 APP_NAME • This is an automated message. Please do not reply. 
                  For assistance, contact support at <a href="mailto:support@yourapp.com" style="color:#4F46E5;text-decoration:underline;">support@yourapp.com</a>.
                </p>
              </td>
            </tr>

          </table>
        </td>
      </tr>
    </table>
  </body>
</html>`
}

func GeneratePasswordResetSuccessEmail() string {
	return `
		<html>
		<body>
			<h2>Password Reset Successful</h2>
			<p>Your password has been successfully reset.</p>
			<p>If you did not perform this action, please contact our support team immediately.</p>
		</body>
		</html>
	`
}

func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
