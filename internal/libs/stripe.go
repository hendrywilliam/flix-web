package libs

import (
	"errors"
	"os"
	"sync"

	"github.com/stripe/stripe-go/v78/client"
)

var once sync.Once

func StripeClient() *client.API {
	var stripeClient *client.API

	once.Do(func() {
		sk := os.Getenv("STRIPE_SECRET_KEY")
		if len(sk) == 0 {
			panic(errors.New("no stripe secret key in .env"))
		}
		sc := &client.API{}
		sc.Init(sk, nil)
		stripeClient = sc
	})

	return stripeClient
}
