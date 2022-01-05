package dnsmadeasy_test

import (
	"fmt"
	"testing"

	"github.com/ndemeshchenko/dnsmadeasy"
	"github.com/stretchr/testify/assert"
)

func TestDMEClient_CreateRecord(t *testing.T) {

	client, err := newm()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(client)

	// purgeSandboxDomains(client)
	// fmt.Println(err)

	domain, err := genDomain(client)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("test domain name: %s, id: %d", domain.Name, domain.ID)

	rcrd := &dnsmadeasy.Record{
		Name:        "testme",
		Type:        "A",
		Value:       "1.1.1.1",
		TTL:         180,
		GtdLocation: "DEFAULT",
	}

	newRecord, err := client.CreateRecord(domain.ID, rcrd)
	if err != nil {
		t.Errorf("Error creating record %s on domain %s. %s", rcrd.Name, domain.Name, err)
	}

	record, err := client.Record(domain.ID, newRecord.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, newRecord.ID, record.ID, "ID should be equal")
	assert.Equal(t, rcrd.Name, record.Name, "Name should be equal")
	assert.Equal(t, rcrd.Type, record.Type, "Type should be equal")
	assert.Equal(t, rcrd.Value, record.Value, "Value should be equal")
	assert.Equal(t, rcrd.Value, record.Value, "Value should be equal")
	assert.Equal(t, rcrd.TTL, record.TTL, "TTL should be equal")
	assert.Equal(t, rcrd.GtdLocation, record.GtdLocation, "GtdLocation should be equal")

}
