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
	
	if err = conn.Connect(); err != nil {
		t.Fatal(err)
	}

	if err = conn.Close(); err != nil {
		t.Fatal(err)
	}
}
