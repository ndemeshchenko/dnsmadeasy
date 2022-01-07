package dnsmadeasy_test

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/ndemeshchenko/dnsmadeasy"
)

var (
	accessKey = flag.String("APIAccessKey", "", "DNSMadeEasy API Access Key")
	secretKey = flag.String("APISecretKey", "", "DNSMadeEasy API Secret Key")
	// cleanupTest    = flag.Bool("purge", true, "cleanup all gotest-* domains from sandbox account")
	testDomains    = make(map[string]*dnsmadeasy.Domain)
	domainBaseName = "dmetest"
)

func TestClient(t *testing.T) {
	client, err := newm()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(client)
}

func TestCreateDomain(t *testing.T) {
	client, err := newm()
	if err != nil {
		t.Fatal(err)
	}

	domain, err := genDomain(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Created SNDBX domain name %s", domain.Name)
}

func newm() (*dnsmadeasy.DMEClient, error) {
	return dnsmadeasy.New(&dnsmadeasy.DMEClient{
		APIUrl:               dnsmadeasy.SANDBOXAPI,
		APIAccessKey:         *accessKey,
		APISecretKey:         *secretKey,
		DisableTLSValidation: true,
	})
}

func genDomain(client *dnsmadeasy.DMEClient) (*dnsmadeasy.Domain, error) {
	genName := fmt.Sprintf("%s-%v.io", domainBaseName, time.Now().UnixNano())
	domain, err := client.CreateDomain(&dnsmadeasy.Domain{
		Name: genName,
	})

	if err != nil {
		return nil, err
	}

	testDomains[genName] = domain
	return domain, err
}

// commented because of domain deletion broken on DNSMadeEasy Sandbox API
//
//func purgeSandboxDomains(client *dnsmadeasy.DMEClient) error {
//	domains, err := client.Domains()
//	if err != nil {
//		return err
//	}
//
//	testDomainsPatter := regexp.MustCompile(fmt.Sprintf("^%s-\\d+\\.io$", domainBaseName))
//	var wg sync.WaitGroup
//	for _, domain := range domains {
//		if !testDomainsPatter.MatchString(domain.Name) {
//			fmt.Printf("domain %s doesn't match pattern", domain.Name)
//			continue
//		}
//
//		wg.Add(1)
//		go func(rmDomain dnsmadeasy.Domain) {
//			defer wg.Done()
//			fmt.Printf("Deleting %s\n", rmDomain.Name)
//
//			err := client.DeleteDomain(rmDomain.ID, 60*time.Second)
//			if err != nil {
//				fmt.Printf("Can't delete domain %s. error: %v\n", rmDomain.Name, err)
//			}
//		}(domain)
//
//	}
//	wg.Wait()
//
//	return nil
//}
