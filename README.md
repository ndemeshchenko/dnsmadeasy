# dnsmadeasy

dnsmadeasy is a Go lang wrapper library for DNS Made Easy provider API.
This is a minimum implementation to provide functionality to create and delete DNS records.
Primary usage by Jetstack cert-manager as webhook https://github.com/ndemeshchenko/cert-manager-webhook-dnsmadeasy 

## Usage

Create DNSMadeEasy client with `New()` method

```go
import "github.com/ndemeshchenko/dnsmadeasy"

client, err := dnsmadeasy.New(&dnsmadeasy.DMEClient{
		APIUrl:               dnsmadeasy.PRODAPI,
		APIAccessKey:         *accessKey,
		APISecretKey:         *secretKey,
		DisableTLSValidation: false,
	})

if err != nil {
	panic(err)
}
```