package client

import (
	"simulation/conn"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	conn, err := conn.NewEdgeConn("../config/clientcfg.yaml")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, conn, "expected non-nil edgeConn instance")
	
	err = conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
}
