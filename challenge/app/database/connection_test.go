package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := connect()
	if len(db.Session.LiveServers()) < 1 {
		t.Errorf("Expected to succeed database connect but there are no servers alive")
	}
}
