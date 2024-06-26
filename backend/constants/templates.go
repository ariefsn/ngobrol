package constants

type MailTemplate string

const (
	TemplateAuthVerificationCode  MailTemplate = "auth_verification_code.mjml"
	TemplateAuthResetPasswordLink MailTemplate = "auth_reset_password.mjml"
	TemplateEmailAccountCreated   MailTemplate = "email_account_created.mjml"
	TemplateEmailAccountApproved  MailTemplate = "email_account_approved.mjml"
	TemplateEmailAccountDeclined  MailTemplate = "email_account_declined.mjml"
	TemplateEnquiry               MailTemplate = "enquiry.mjml"
)
