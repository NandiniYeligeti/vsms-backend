package services

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"vehiclesales/models"

	"github.com/jung-kurt/gofpdf"
)

type PDFService interface {
	GeneratePaymentReceipt(settings *models.CompanySettings, payment *models.Payment, customer *models.Customer) ([]byte, error)
	GenerateSalesOrder(settings *models.CompanySettings, order *models.SalesOrder) ([]byte, error)
}

type pdfService struct{}

func NewPDFService() PDFService {
	return &pdfService{}
}

func (s *pdfService) GeneratePaymentReceipt(settings *models.CompanySettings, payment *models.Payment, customer *models.Customer) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Logo
	if settings.LogoURL != "" {
		logoPath := settings.LogoURL
		if !filepath.IsAbs(logoPath) {
			cwd, _ := os.Getwd()
			logoPath = filepath.Join(cwd, logoPath)
		}

		if _, err := os.Stat(logoPath); err == nil {
			pdf.ImageOptions(logoPath, 10, 10, 30, 0, false, gofpdf.ImageOptions{ImageType: "", ReadDpi: true}, 0, "")
			pdf.SetY(40)
		} else {
			pdf.SetY(10)
		}
	} else {
		pdf.SetY(10)
	}

	// Header
	yAfterLogo := pdf.GetY()
	pdf.SetY(yAfterLogo)
	pdf.SetFont("Arial", "B", 24)
	pdf.SetTextColor(37, 99, 235) // #2563eb
	pdf.CellFormat(0, 10, settings.CompanyName, "", 0, "L", false, 0, "")
	
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(0, 10, "RECEIPT", "", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(0, 5, "Payment Receipt", "", 0, "L", false, 0, "")
	
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(0, 5, fmt.Sprintf("No: %s", payment.PaymentCode), "", 1, "R", false, 0, "")
	pdf.CellFormat(0, 5, fmt.Sprintf("Date: %s", payment.PaymentDate.Format("02-Jan-2006")), "", 1, "R", false, 0, "")

	pdf.Ln(10)
	pdf.SetDrawColor(238, 238, 238)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(10)

	// Grid: Customer Details & Reference Details
	yStart := pdf.GetY()
	
	// Left Column: Customer Details
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(95, 5, "CUSTOMER DETAILS", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(95, 5, customer.CustomerName, "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 5, customer.MobileNumber, "", 1, "L", false, 0, "")
	pdf.CellFormat(95, 5, customer.Email, "", 1, "L", false, 0, "")

	// Right Column: Reference Details
	pdf.SetY(yStart)
	pdf.SetX(110)
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(90, 5, "REFERENCE DETAILS", "", 1, "L", false, 0, "")
	
	pdf.SetX(110)
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(45, 5, "Invoice/SO:", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(45, 5, payment.InvoiceNumber, "", 1, "R", false, 0, "")
	
	pdf.SetX(110)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(45, 5, "Category:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, payment.PaymentType, "", 1, "R", false, 0, "")
	
	pdf.SetX(110)
	pdf.CellFormat(45, 5, "Mode:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, payment.PaymentMode, "", 1, "R", false, 0, "")

	if payment.ReferenceNumber != "" {
		pdf.SetX(110)
		pdf.CellFormat(45, 5, "Ref No:", "", 0, "L", false, 0, "")
		pdf.CellFormat(45, 5, payment.ReferenceNumber, "", 1, "R", false, 0, "")
	}

	pdf.Ln(20)

	// Payment Summary
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(0, 5, "PAYMENT SUMMARY", "B", 1, "L", false, 0, "")
	
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(140, 10, "Description", "B", 0, "L", false, 0, "")
	pdf.CellFormat(0, 10, "Amount", "B", 1, "R", false, 0, "")
	
	pdf.CellFormat(140, 10, fmt.Sprintf("%s Credit", payment.PaymentType), "B", 0, "L", false, 0, "")
	pdf.CellFormat(0, 10, fmt.Sprintf("Rs. %.2f", payment.PaymentAmount), "B", 1, "R", false, 0, "")

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(37, 99, 235)
	pdf.CellFormat(140, 15, "Total Amount Received", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 15, fmt.Sprintf("Rs. %.2f", payment.PaymentAmount), "", 1, "R", false, 0, "")

	// Signatures
	pdf.SetY(240)
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(51, 51, 51)
	
	pdf.SetX(10)
	pdf.CellFormat(60, 5, "Customer Signature", "T", 0, "C", false, 0, "")
	
	pdf.SetX(140)
	pdf.CellFormat(60, 5, "Authorized Signatory", "T", 1, "C", false, 0, "")

	// Footer
	pdf.SetY(260)
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(153, 153, 153)
	pdf.CellFormat(0, 5, "This is a computer-generated receipt and does not require a physical signature.", "", 1, "C", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *pdfService) GenerateSalesOrder(settings *models.CompanySettings, order *models.SalesOrder) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Logo
	if settings.LogoURL != "" {
		logoPath := settings.LogoURL
		if !filepath.IsAbs(logoPath) {
			cwd, _ := os.Getwd()
			logoPath = filepath.Join(cwd, logoPath)
		}

		if _, err := os.Stat(logoPath); err == nil {
			pdf.ImageOptions(logoPath, 10, 10, 30, 0, false, gofpdf.ImageOptions{ImageType: "", ReadDpi: true}, 0, "")
			pdf.SetY(40)
		} else {
			pdf.SetY(10)
		}
	} else {
		pdf.SetY(10)
	}

	// Header
	yAfterLogo := pdf.GetY()
	pdf.SetY(yAfterLogo)
	pdf.SetFont("Arial", "B", 24)
	pdf.SetTextColor(37, 99, 235)
	pdf.CellFormat(0, 10, settings.CompanyName, "", 0, "L", false, 0, "")
	
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(0, 10, "SALES ORDER", "", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(0, 5, "Order Confirmation", "", 0, "L", false, 0, "")
	
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(0, 5, fmt.Sprintf("Order: %s", order.SalesOrderCode), "", 1, "R", false, 0, "")
	pdf.CellFormat(0, 5, fmt.Sprintf("Date: %s", order.SaleDate.Format("02-Jan-2006")), "", 1, "R", false, 0, "")

	pdf.Ln(10)
	pdf.SetDrawColor(238, 238, 238)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(10)

	// Customer Info
	yStart := pdf.GetY()
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(95, 5, "CUSTOMER DETAILS", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(95, 5, order.CustomerName, "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 5, order.MobileNumber, "", 1, "L", false, 0, "")
	pdf.CellFormat(95, 5, order.Email, "", 1, "L", false, 0, "")
	pdf.MultiCell(95, 5, order.Address, "", "L", false)

	// Vehicle Info
	pdf.SetY(yStart)
	pdf.SetX(110)
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(90, 5, "VEHICLE DETAILS", "", 1, "L", false, 0, "")
	
	pdf.SetX(110)
	pdf.SetFont("Arial", "B", 10)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(90, 5, fmt.Sprintf("%s %s", order.Brand, order.Model), "", 1, "L", false, 0, "")
	
	pdf.SetX(110)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(45, 5, "Variant:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, order.Variant, "", 1, "R", false, 0, "")
	
	pdf.SetX(110)
	pdf.CellFormat(45, 5, "Color:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, order.Color, "", 1, "R", false, 0, "")
	
	pdf.SetX(110)
	pdf.CellFormat(45, 5, "Chassis:", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, order.ChassisNumber, "", 1, "R", false, 0, "")

	pdf.Ln(20)

	// Pricing Table
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(102, 102, 102)
	pdf.CellFormat(0, 5, "ORDER SUMMARY", "B", 1, "L", false, 0, "")
	
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(51, 51, 51)
	
	rows := [][]string{
		{"Vehicle Price", fmt.Sprintf("%.2f", order.VehiclePrice)},
		{"Registration Charges", fmt.Sprintf("%.2f", order.RegistrationCharges)},
		{"Insurance", fmt.Sprintf("%.2f", order.Insurance)},
		{"Accessories", fmt.Sprintf("%.2f", order.Accessories)},
		{"Discount", fmt.Sprintf("-%.2f", order.DiscountAmount)},
	}

	for _, row := range rows {
		pdf.CellFormat(140, 8, row[0], "B", 0, "L", false, 0, "")
		pdf.CellFormat(0, 8, fmt.Sprintf("Rs. %s", row[1]), "B", 1, "R", false, 0, "")
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.SetTextColor(37, 99, 235)
	pdf.CellFormat(140, 10, "Total Amount", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 10, fmt.Sprintf("Rs. %.2f", order.TotalAmount), "", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(51, 51, 51)
	pdf.CellFormat(140, 8, "Down Payment", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 8, fmt.Sprintf("Rs. %.2f", order.DownPayment), "", 1, "R", false, 0, "")
	
	pdf.CellFormat(140, 8, "Loan Amount", "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 8, fmt.Sprintf("Rs. %.2f", order.LoanAmount), "", 1, "R", false, 0, "")

	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(37, 99, 235)
	pdf.CellFormat(140, 15, "Balance Amount", "T", 0, "L", false, 0, "")
	pdf.CellFormat(0, 15, fmt.Sprintf("Rs. %.2f", order.BalanceAmount), "T", 1, "R", false, 0, "")

	// Signatures
	pdf.SetY(240)
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(51, 51, 51)
	pdf.SetX(10)
	pdf.CellFormat(60, 5, "Customer Signature", "T", 0, "C", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(60, 5, "Authorized Signatory", "T", 1, "C", false, 0, "")

	// Footer
	pdf.SetY(260)
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(153, 153, 153)
	pdf.CellFormat(0, 5, "This is a computer-generated document and does not require a physical signature.", "", 1, "C", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
