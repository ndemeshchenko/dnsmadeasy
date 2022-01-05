package dnsmadeasy

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const PRODAPI = "https://api.dnsmadeeasy.com/V2.0/"
const SANDBOXAPI = "https://api.sandbox.dnsmadeeasy.com/V2.0/"
const pendingDeleteError = "Cannot delete a domain that is pending a create or delete action."

type DMEClient struct {
	// set with Option calls
	APIUrl               string
	APIAccessKey         string
	APISecretKey         string
	DisableTLSValidation bool

	//internal use
	dmeClient *http.Client
}

func New(dme *DMEClient) (*DMEClient, error) {
	if dme.APIAccessKey == "" {
		return nil, fmt.Errorf("API Access Key is empty. can't proceed")
	}

	if dme.APISecretKey == "" {
		return nil, fmt.Errorf("API Secret Key is empty. can't proceed")
	}

	if dme.APIUrl == "" {
		dme.APIUrl = PRODAPI
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: dme.DisableTLSValidation,
		},
	}

	dme.dmeClient = &http.Client{Transport: tr}

	return dme, nil
}

func (dme *DMEClient) requestTemplate(method, endpoint string, body io.Reader) (*http.Request, error) {
	tNow := time.Now().UTC()
	tNowString := tNow.Format(time.RFC1123)
	key := []byte(dme.APISecretKey)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(tNowString))
	hmacSha := hex.EncodeToString(h.Sum(nil))

	URI := dme.APIUrl + endpoint
	req, err := http.NewRequest(method, URI, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-dnsme-apiKey", dme.APIAccessKey)
	req.Header.Set("x-dnsme-requestDate", tNowString)
	req.Header.Set("x-dnsme-hmac", hmacSha)

	return req, nil
}

func (dme *DMEClient) fireRequest(req *http.Request, dst interface{}) error {
	res, err := dme.dmeClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if body != nil {
		body = []byte(strings.Replace(string(body), "{error:", "{\"error\":", 1))
	}

	if res.StatusCode == http.StatusForbidden {
		return fmt.Errorf("access forbidden (%s)", req.URL.String())
	}
	if res.StatusCode == http.StatusNotFound {
		return fmt.Errorf("404 Not Found (%s)", req.URL.String())
	}

	genericParsingError := &GenericParsingError{}

	json.Unmarshal(body, genericParsingError)
	if len(genericParsingError.Error) > 0 {
		return fmt.Errorf(strings.Join(genericParsingError.Error, "\n"))
	}

	if req.Method == "PUT" || req.Method == "DELETE" {
		return nil
	}

	err = json.Unmarshal(body, dst)
	return err
}

type GenericParsingError struct {
	Error []string `json:"error"`
}
