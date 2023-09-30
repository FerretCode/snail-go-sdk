package snail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Snail struct {
	ApiKey string
}

func NewSnail(apiKey string) Snail {
	return Snail{
		ApiKey: apiKey,
	}
}

type Payment struct {
	Created int64 `json:"created"`	
	Customer string `json:"customer"`
	Email string `json:"email"`
	Status string `json:"status"`
	Amount int64 `json:"amount"`
	Product string `json:"product"`
	Subscription bool `json:"subscription"`
}

type ErrCodeInvalid struct{}

func (e *ErrCodeInvalid) Error() string {
	return "The payment was invalid."
}

func (s *Snail) VerifyPayment(code string) (Payment, error) {
	if len(code) != 10 {
		return Payment{}, &ErrCodeInvalid{}
	}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://snailpay.app/verify-payment?code=%s",
			code,
		),
		nil,
	)

	if err != nil {
		return Payment{}, err 
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	payment := Payment{}

	_, status, err := doRequest(req, &payment)	

	if err != nil {
		return Payment{}, err
	}

	if status != 200 {
		return Payment{}, &ErrCodeInvalid{}
	}

	return payment, nil
}

type PaymentLinkParams struct {
	Image string `json:"image"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type PaymentLink struct {
	URL string `json:"url"`
}

func (s *Snail) CreatePaymentLink(params *PaymentLinkParams) (PaymentLink, error) {
	stringified, err := json.Marshal(&params)

	if err != nil {
		return PaymentLink{}, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://snailpay.app/payment-link",
		bytes.NewBuffer(stringified),
	)

	if err != nil {
		return PaymentLink{}, err 
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	paymentLink := PaymentLink{}

	text, status, err := doRequest(req, &paymentLink)	

	if err != nil {
		return PaymentLink{}, err
	}

	if status != 200 {
		return PaymentLink{}, errors.New(text)
	}

	return paymentLink, nil
} 

type SubscriptionLinkParams struct {
	Image string `json:"image"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type SubscriptionLink struct {
	URL string `json:"url"`
}

func (s *Snail) SubscriptionLink(params *SubscriptionLink) (SubscriptionLink, error) {
	stringified, err := json.Marshal(&params)

	if err != nil {
		return SubscriptionLink{}, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://snailpay.app/subscription-link",
		bytes.NewBuffer(stringified),
	)

	if err != nil {
		return SubscriptionLink{}, err 
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	subscriptionLink := SubscriptionLink{}

	text, status, err := doRequest(req, &subscriptionLink)	

	if err != nil {
		return SubscriptionLink{}, err
	}

	if status != 200 {
		return SubscriptionLink{}, errors.New(text)
	}

	return subscriptionLink, nil
} 

type ListPayment struct {
	Amount int64 `json:"amount"`
	Customer string `json:"customer"`
	Email string `json:"email"`
	ID string `json:"id"`
	Timestamp int64 `json:"timestamp"`
	Status string `json:"status"`
}

func (s *Snail) ListPayments() ([]ListPayment, error) {
	req, err := http.NewRequest(
		"GET",
		"https://snailpay.app/payment-list",
		nil,
	)

	if err != nil {
		return []ListPayment{}, err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	var listPayments []ListPayment

	text, status, err := doRequest(req, &listPayments)

	if err != nil {
		return []ListPayment{}, err
	}

	if status != 200 {
		return []ListPayment{}, errors.New(text)
	}

	return listPayments, nil
}

type ListSubscription struct {
	Amount int64 `json:"amount"`
	Customer string `json:"customer"`
	Email string `json:"email"`
	ID string `json:"id"`
	Timestamp int64 `json:"timestamp"`
	Status string `json:"status"`
}

func (s *Snail) ListSubscriptions() ([]ListSubscription, error) {
	req, err := http.NewRequest(
		"GET",
		"https://snailpay.app/subscription-list",
		nil,
	)

	if err != nil {
		return []ListSubscription{}, err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	var listSubscriptions []ListSubscription

	text, status, err := doRequest(req, &listSubscriptions)

	if err != nil {
		return []ListSubscription{}, err
	}

	if status != 200 {
		return []ListSubscription{}, errors.New(text)
	}

	return listSubscriptions, nil
}

func (s *Snail) ListPaymentLinks() ([]string, error) {
	req, err := http.NewRequest(
		"GET",
		"https://snailpay.app/payment-link-list",
		nil,
	)

	if err != nil {
		return []string{}, err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	var links []string

	text, status, err := doRequest(req, &links)

	if err != nil {
		return []string{}, err
	}

	if status != 200 {
		return []string{}, errors.New(text)
	}

	return links, nil
}

func (s *Snail) ListSubscriptionLinks() ([]string, error) {
	req, err := http.NewRequest(
		"GET",
		"https://snailpay.app/subscription-link-list",
		nil,
	)

	if err != nil {
		return []string{}, err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	var links []string

	text, status, err := doRequest(req, &links)

	if err != nil {
		return []string{}, err
	}

	if status != 200 {
		return []string{}, errors.New(text)
	}

	return links, nil
}

type Payouts struct { 
	PayoutList []Payout
	Withdrawn float64 
	Balance float64
	Pending float64
}

type Payout struct {
	Date int64 `json:"date"`
	Amount int64 `json:"amount"`
	ArrivalDate int64 `json:"arrival_date"`
	Status string `json:"status"`
}

func (s *Snail) ListPayouts() (Payouts, error) {
	req, err := http.NewRequest(
		"GET",
		"https://snailpay.app/payout",
		nil,
	)

	if err != nil {
		return Payouts{}, err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	payouts := Payouts{}

	text, status, err := doRequest(req, &payouts)

	if err != nil {
		return Payouts{}, err
	}

	if status != 200 {
		return Payouts{}, errors.New(text)
	}

	return payouts, err
}

func (s *Snail) NewPayout(amount float64) error {
	req, err := http.NewRequest(
		"POST",
		"https://snailpay.app/new-payout",
		nil,
	)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	return nil
}

type Payments struct {
	Payments []string `json:"payments"`
}

func (s *Snail) RefundPayments(paymentIds []string) error {
	payments := Payments{
		Payments: paymentIds,
	}	

	stringified, err := json.Marshal(payments)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://snailpay.app/refund-payment",
		bytes.NewBuffer(stringified),
	)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", s.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	return nil
}

func doRequest(req *http.Request, destination interface{}) (string, int, error) {
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return res.Status, res.StatusCode, err
	}

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return res.Status, res.StatusCode, err
	}

	if err := json.Unmarshal(bytes, &destination); err != nil {
		return res.Status, res.StatusCode, err
	} 

	return res.Status, res.StatusCode, nil
}
