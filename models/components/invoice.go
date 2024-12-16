// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"HathoraCloud/internal/utils"
	"time"
)

type Invoice struct {
	Status    InvoiceStatus `json:"status"`
	AmountDue float64       `json:"amountDue"`
	PdfURL    string        `json:"pdfUrl"`
	DueDate   time.Time     `json:"dueDate"`
	Year      float64       `json:"year"`
	Month     float64       `json:"month"`
	ID        string        `json:"id"`
}

func (i Invoice) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Invoice) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Invoice) GetStatus() InvoiceStatus {
	if o == nil {
		return InvoiceStatus("")
	}
	return o.Status
}

func (o *Invoice) GetAmountDue() float64 {
	if o == nil {
		return 0.0
	}
	return o.AmountDue
}

func (o *Invoice) GetPdfURL() string {
	if o == nil {
		return ""
	}
	return o.PdfURL
}

func (o *Invoice) GetDueDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.DueDate
}

func (o *Invoice) GetYear() float64 {
	if o == nil {
		return 0.0
	}
	return o.Year
}

func (o *Invoice) GetMonth() float64 {
	if o == nil {
		return 0.0
	}
	return o.Month
}

func (o *Invoice) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
