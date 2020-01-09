package payment

type paymentContext struct{}
type paymentTransaction struct{}

type PayFunc func(paymentCtx paymentContext) (paymentTransaction, error)

func Pay(c clienter) PayFunc {
	return func(paymentCtx paymentContext) (paymentTransaction, error) {
		return paymentTransaction{}, nil
	}
}
