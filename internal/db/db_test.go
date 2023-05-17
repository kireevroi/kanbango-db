package db

import (
	"os"
	"testing"

	"github.com/kireevroi/kanbango/auth/pkg/onstart"
)

func TestGetAllSessions(t *testing.T) {
	d := NewDB()
	onstart.LoadEnv("../../.env")
	d.Connect(os.Getenv("DBURLM"))
	t.Log(os.Getenv("DBURLM"))
	x, _ := d.GetAllSessions(1)
	for i, _ := range x {
		t.Log(x[i].Session)
	}
}