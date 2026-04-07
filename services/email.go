package services

import (
	"context"
	"fmt"
	"net/smtp"
	"vehiclesales/models"
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
}

type emailService struct {
	companyService CompanySettingsService
}

func NewEmailService() EmailService {
	return &emailService{
		companyService: NewCompanySettingsService(),
	}
}

func (s *emailService) SendSalesOrderConfirmation(ctx context.Context, companyCode string, order *models.SalesOrder) error {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil || settings == nil || !settings.EmailSettings.EnableEmail {
		return nil // Not enabled or error
	}

	mail := settings.EmailSettings
	auth := smtp.PlainAuth("", mail.EmailUsername, mail.EmailPassword, mail.SMTPHost)

	preview, _ := s.PreviewSalesOrderEmail(ctx, companyCode, order)
	subject := preview.Subject
	body := preview.Body

	msg := []byte("To: " + order.Email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	addr := fmt.Sprintf("%s:%d", mail.SMTPHost, mail.SMTPPort)
	return smtp.SendMail(addr, auth, mail.SenderEmail, []string{order.Email}, msg)
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

Order Details:
Order Number: %s
Vehicle: %s %s %s
Total Amount: %.2f
Balance Due: %.2f

Status: %s

Best regards,
%s
`, order.CustomerName, settings.CompanyName, order.SalesOrderCode, order.Brand, order.Model, order.Variant, order.TotalAmount, order.BalanceAmount, order.Status, settings.EmailSettings.SenderName)

	return &EmailPreview{
		To:      order.Email,
		Subject: subject,
		Body:    body,
	}, nil
}

func (s *emailService) SendPaymentReceipt(ctx context.Context, companyCode string, payment *models.Payment, customer *models.Customer) error {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil || settings == nil || !settings.EmailSettings.EnableEmail || !settings.EmailSettings.AutoSendReceipt {
		return nil
	}

	mail := settings.EmailSettings
	auth := smtp.PlainAuth("", mail.EmailUsername, mail.EmailPassword, mail.SMTPHost)

	preview, _ := s.PreviewPaymentReceiptEmail(ctx, companyCode, payment, customer)
	subject := preview.Subject
	body := preview.Body

	msg := []byte("To: " + customer.Email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	addr := fmt.Sprintf("%s:%d", mail.SMTPHost, mail.SMTPPort)
	return smtp.SendMail(addr, auth, mail.SenderEmail, []string{customer.Email}, msg)
}

func (s *emailService) PreviewPaymentReceiptEmail(ctx context.Context, companyCode string, payment *models.Payment, customer *models.Customer) (*EmailPreview, error) {
	settings, err := s.companyService.Get(ctx, companyCode)
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprintf("Payment Receipt - %s", payment.PaymentCode)
	body := fmt.Sprintf(`
Dear %s,

We have received your payment of %.2f.

Payment Details:
Receipt Number: %s
Date: %s
Payment Mode: %s
Amount: %.2f

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
