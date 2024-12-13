// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PaymentMethod - Make all properties in T optional
type PaymentMethod struct {
	Card *CardPaymentMethod `json:"card,omitempty"`
	Ach  *AchPaymentMethod  `json:"ach,omitempty"`
	Link *LinkPaymentMethod `json:"link,omitempty"`
}

func (o *PaymentMethod) GetCard() *CardPaymentMethod {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *PaymentMethod) GetAch() *AchPaymentMethod {
	if o == nil {
		return nil
	}
	return o.Ach
}

func (o *PaymentMethod) GetLink() *LinkPaymentMethod {
	if o == nil {
		return nil
	}
	return o.Link
}
