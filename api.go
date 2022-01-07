package dnsmadeasy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Domains wrapper for API call to fetch domains
func (dme *DMEClient) Domains() ([]Domain, error) {
	req, err := dme.requestTemplate("GET", "dns/managed", nil)
	if err != nil {
		return nil, err
	}

	genericParsingResponse := &genericParsingResponse{}
	err = dme.fireRequest(req, &genericParsingResponse)
	if err != nil {
		return nil, err
	}
	var data []Domain
	err = json.Unmarshal(genericParsingResponse.Data, &data)

	return data, err
}

// Domain wrapper for API call to fetch domain
func (dme *DMEClient) Domain(domainID int) (*Domain, error) {
	uri := fmt.Sprintf("dns/managed/%v", domainID)
	req, err := dme.requestTemplate("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	domain := &Domain{}
	err = dme.fireRequest(req, domain)
	if err != nil {
		return nil, err
	}

	return domain, err
}

// CreateDomain wrapper for API call to create a domain
func (dme *DMEClient) CreateDomain(reqDomain *Domain) (*Domain, error) {
	uri := "dns/managed"
	payload, err := json.Marshal(reqDomain)
	if err != nil {
		return nil, err
	}

	payloadBuffer := bytes.NewReader(payload)
	req, err := dme.requestTemplate("POST", uri, payloadBuffer)
	if err != nil {
		return nil, err
	}

	retDomain := &Domain{}
	err = dme.fireRequest(req, retDomain)
	if err != nil {
		return nil, err
	}

	return retDomain, err
}

// DeleteDomain wrapper for API call to delete a domain
func (dme *DMEClient) DeleteDomain(domainID int, timeout time.Duration) error {
	return dme.delete(fmt.Sprintf("dns/managed/%v", domainID), timeout)
}

// Records wrapper for API call to fetch list of records
func (dme *DMEClient) Records(domainID int) ([]Record, error) {
	uri := fmt.Sprintf("dns/managed/%v/records", domainID)
	req, err := dme.requestTemplate("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	genericParsingResponse := &genericParsingResponse{}
	err = dme.fireRequest(req, &genericParsingResponse)
	if err != nil {
		return nil, err
	}

	var data []Record
	err = json.Unmarshal(genericParsingResponse.Data, &data)

	return data, err
}

func (dme *DMEClient) Record(domainID, recordID int) (*Record, error) {
	records, err := dme.Records(domainID)
	if err != nil {
		return nil, err
	}

	for i := range records {
		if records[i].ID == recordID {
			return &records[i], nil
		}
	}

	return nil, fmt.Errorf("no records with id: %d found", recordID)
}

func (dme *DMEClient) CreateRecord(domainID int, reqRecord *Record) (*Record, error) {
	uri := fmt.Sprintf("dns/managed/%v/records", domainID)
	payload, err := json.Marshal(reqRecord)
	if err != nil {
		return nil, err
	}

	payloadBuffer := bytes.NewReader(payload)
	req, err := dme.requestTemplate("POST", uri, payloadBuffer)
	if err != nil {
		return nil, err
	}

	retRecord := &Record{}
	err = dme.fireRequest(req, retRecord)
	if err != nil {
		return nil, err
	}

	return retRecord, err
}

func (dme *DMEClient) UpdateRecord(domainID int, record *Record) error {
	uri := fmt.Sprintf("dns/managed/%v/records/%v", domainID, record.ID)
	payload, err := json.Marshal(record)
	if err != nil {
		return err
	}

	return dme.update(uri, payload)
}

func (dme *DMEClient) DeleteRecord(domainID, recordID int) error {
	uri := fmt.Sprintf("dns/managed/%v/records/%v", domainID, recordID)
	req, err := dme.requestTemplate("DELETE", uri, nil)
	if err != nil {
		return err
	}
	return dme.fireRequest(req, nil)
}

// PUT call wrapper
func (dme *DMEClient) update(uri string, payload []byte) error {
	payloadBuffer := bytes.NewReader(payload)
	req, err := dme.requestTemplate("PUT", uri, payloadBuffer)
	if err != nil {
		return err
	}

	return dme.fireRequest(req, nil)
}

// DELETE call wrapper
func (dme *DMEClient) delete(uri string, timeout time.Duration) error {
	req, err := dme.requestTemplate("DELETE", uri, nil)
	if err != nil {
		return err
	}

	deleteError := dme.fireRequest(req, nil)
	if deleteError == nil || timeout == 0 {
		return deleteError
	}

	timeoutAt := time.Now().Add(timeout)
	for time.Now().Before(timeoutAt) {
		req, _ := dme.requestTemplate("DELETE", uri, nil)
		deleteError := dme.fireRequest(req, nil)

		if deleteError == nil {
			return deleteError
		}

		if deleteError.Error() != pendingDeleteError {
			return deleteError
		}

		time.Sleep(15 * time.Second)
	}

	return err
}
