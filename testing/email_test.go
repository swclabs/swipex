package test

import (
	"swclabs/swix/internal/core/x/mail"
	"testing"
)

func TestEmail(t *testing.T) {
	m := mail.New()
	err := m.SendPurchaseOrder("iduchungho@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
}
