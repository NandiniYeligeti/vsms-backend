package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/smtp"
	"vehiclesales/models"

	"github.com/jordan-wright/email"
)

type EmailPreview struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailService interface {
	SendSalesOrderConfirmation(ctx context.Context, companyCode string, order *models.SalesOrder) error
	SendPaymentReceipt(ctx context.Context, companyCode string, payment *models.Payment, customer *models.Customer) error
	PreviewSalesOrderEmail(ctx context.Context, companyCode string, order *models.SalesOrder) (*EmailPreview, error)
	PreviewPaymentReceiptEmail(ctx context.Context, companyCode string, payment *models.Payment, customer *models.Customer) (*EmailPreview, error)
	SendForgotPasswordEmail(ctx context.Context, user *models.User) error
}

type emailService struct {
	companyService CompanySettingsService
	pdfService     PDFService
}

func NewEmailService() EmailService {
	return &emailService{
		companyService: NewCompanySettingsService(),
		pdfService:     NewPDFService(),
	}
}

func (s *emailService) SendSalesOrderConfirmation(ctx context.Context, companyCode string, order *models.SalesOrder) error {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil || settings == nil || !settings.EmailSettings.EnableEmail {
		return nil
	}

	mail := settings.EmailSettings
	preview, _ := s.PreviewSalesOrderEmail(ctx, companyCode, order)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", mail.SenderName, mail.SenderEmail)
	e.To = []string{order.Email}
	e.Subject = preview.Subject
	e.Text = []byte(preview.Body)

	// Generate and attach PDF
	pdfBytes, err := s.pdfService.GenerateSalesOrder(settings, order)
	if err == nil {
		e.Attach(bytes.NewReader(pdfBytes), fmt.Sprintf("Order_%s.pdf", order.SalesOrderCode), "application/pdf")
	}

	auth := smtp.PlainAuth("", mail.EmailUsername, mail.EmailPassword, mail.SMTPHost)
	addr := fmt.Sprintf("%s:%d", mail.SMTPHost, mail.SMTPPort)
	return e.Send(addr, auth)
}

func (s *emailService) PreviewSalesOrderEmail(ctx context.Context, companyCode string, order *models.SalesOrder) (*EmailPreview, error) {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Order Confirmation - %s", order.SalesOrderCode)
	body := fmt.Sprintf(`
Dear %s,

Thank you for your order with %s!

Please find the attached sales order confirmation for your reference.

Order Details:
Order Number: %s
Vehicle: %s %s %s
Total Amount: Rs. %.2f
Status: %s

Best regards,
%s
`, order.CustomerName, settings.CompanyName, order.SalesOrderCode, order.Brand, order.Model, order.Variant, order.TotalAmount, order.Status, settings.EmailSettings.SenderName)

	return &EmailPreview{
		To:      order.Email,
		Subject: subject,
		Body:    body,
	}, nil
}

func (s *emailService) SendPaymentReceipt(ctx context.Context, companyCode string, payment *models.Payment, customer *models.Customer) error {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil || settings == nil || !settings.EmailSettings.EnableEmail {
		return nil
	}

	mail := settings.EmailSettings
	preview, _ := s.PreviewPaymentReceiptEmail(ctx, companyCode, payment, customer)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", mail.SenderName, mail.SenderEmail)
	e.To = []string{customer.Email}
	e.Subject = preview.Subject
	e.Text = []byte(preview.Body)

	// Generate and attach PDF Receipt
	pdfBytes, err := s.pdfService.GeneratePaymentReceipt(settings, payment, customer)
	if err == nil {
		e.Attach(bytes.NewReader(pdfBytes), fmt.Sprintf("Receipt_%s.pdf", payment.PaymentCode), "application/pdf")
	}

	auth := smtp.PlainAuth("", mail.EmailUsername, mail.EmailPassword, mail.SMTPHost)
	addr := fmt.Sprintf("%s:%d", mail.SMTPHost, mail.SMTPPort)
	return e.Send(addr, auth)
}

func (s *emailService) PreviewPaymentReceiptEmail(ctx context.Context, companyCode string, payment *models.Payment, customer *models.Customer) (*EmailPreview, error) {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Payment Receipt - %s", payment.PaymentCode)
	body := fmt.Sprintf(`
Dear %s,

We have received your payment of Rs. %.2f. 

Please find the official payment receipt attached to this email.

Payment Details:
Receipt Number: %s
Date: %s
Payment Mode: %s
Amount: Rs. %.2f

Thank you for your business!

Best regards,
%s
`, customer.CustomerName, payment.PaymentAmount, payment.PaymentCode, payment.PaymentDate.Format("02-Jan-2006"), payment.PaymentMode, payment.PaymentAmount, settings.EmailSettings.SenderName)

	return &EmailPreview{
		To:      customer.Email,
		Subject: subject,
		Body:    body,
	}, nil
}

func (s *emailService) SendForgotPasswordEmail(ctx context.Context, user *models.User) error {
	settings, err := s.companyService.Get(ctx, user.CompanyCode)
	if err != nil || settings == nil || !settings.EmailSettings.EnableEmail {
		return errors.New("email service not configured for this company. please contact administrator")
	}

	mail := settings.EmailSettings
	subject := "Password Recovery - DDR AutoPro"
	body := fmt.Sprintf(`
Dear %s,

You requested your password for DDR AutoPro.

Your Current Password: %s

Please change your password after logging in for security.

Best regards,
%s
`, user.Username, user.Password, settings.EmailSettings.SenderName)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", mail.SenderName, mail.SenderEmail)
	e.To = []string{user.Email}
	e.Subject = subject
	e.Text = []byte(body)

	auth := smtp.PlainAuth("", mail.EmailUsername, mail.EmailPassword, mail.SMTPHost)
	addr := fmt.Sprintf("%s:%d", mail.SMTPHost, mail.SMTPPort)
	return e.Send(addr, auth)
}
