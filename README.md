# snail-go-sdk

A Golang wrapper for the [Snail](https://snailpay.app) API

# Authenticating

```go
package main

import "github.com/ferretcode/snail-go-sdk"

func main() {
	s := snail.NewSnail("your snail api key")
}
```

# Verifying Payments

To use this endpoint, acquire an order verification code from your user and call this function

```go
  payment, err := s.VerifyPayment("user code")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(payment)
```

# Creating Payment Links

If you want your product to have an image, you have to encode an image as base64

```go
  paymentLink, err := s.CreatePaymentLink(&snail.PaymentLinkParams{
		Image: "base64 encoded image",
		Name: "product name",
		Price: 5, // amount of usd to charge
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(paymentLink.URL)
```

# Creating Subscription Links

If you want your product to have an image, you have to encode an image as base64

```go
  paymentLink, err := s.CreateSubscriptionLink(&snail.PaymentLinkParams{
		Image: "base64 encoded image",
		Name: "product name",
		Price: 5, // amount of usd to charge
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(paymentLink.URL)
```

# List Payments

```go
p  ayments, err := s.ListPayments()

	if err != nil {
		log.Fatal(err)
	}
```

# List Subscriptions

```go
  subscriptions, err := s.ListSubscriptions()

	if err != nil {
		log.Fatal(err)
	}
```

# List Payment Links

```go
  paymentLinks, err := s.ListPaymentLinks()

	if err != nil {
		log.Fatal(err)
	}
```

# List Subscription Links

```go
  subscriptionLinks, err := s.ListSubscriptionLinks()

	if err != nil {
		log.Fatal(err)
	}
```

# List Payouts

```javascript
  payouts, err := s.ListPayouts()

	if err != nil {
		log.Fatal(err)
	}
```

# Create a Payout

```javascript
  err := s.NewPayout(5.41)

	if err != nil {
		log.Fatal(err)
	}
```

# Refund a Payment

```javascript
  err := s.RefundPayments([]string{"payment id 1", "payment id 2"})

	if err != nil {
		log.Fatal(err)
	}
```
