package models

type Payment struct {
	// Internal ID for this payment in Metronero
	InvoiceID string `json:"invoice_id"`
	// Merchant provided optional ID
	OrderID string `json:"order_id,omitempty"`
	MerchantID string `json:"merchant_id"`
	Amount int64 `json:"amount_atomic"`
	Status string `json:"status"`
	Description string `json:"description"`
	Date time.Time `json:"date"`
}
