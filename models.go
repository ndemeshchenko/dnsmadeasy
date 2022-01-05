package dnsmadeasy

import "encoding/json"

type genericParsingResponse struct {
	Page         int `json:"page"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
	Data         json.RawMessage
}

// Domain structure for API type
type Domain struct {
	Name                string        `json:"name"`
	NameServer          []string      `json:"nameServer,omitempty"`
	GtdEnabled          bool          `json:"gtdEnabled,omitempty"`
	ID                  int           `json:"id,omitempty"`
	FolderID            int           `json:"folderId,omitempty"`
	NameServers         []NameServer  `json:"nameServers"`
	Updated             int64         `json:"updated,omitempty"`
	TemplateID          int           `json:"templateId,omitempty"`
	DelegateNameServers []string      `json:"delegateNameServers,omitempty"`
	Created             int64         `json:"created,omitempty"`
	TransferAclID       int           `json:"transferAclId,omitempty"`
	ActiveThirdParties  []interface{} `json:"activeThirdParties,omitempty"`
	VanityID            int           `json:"vanityId,omitempty"`
	PendingActionID     int           `json:"pendingActionId,omitempty"`
	SoaID               int           `json:"soaId,omitempty"`
	ProcessMulti        bool          `json:"processMulti,omitempty"`
}

// NameServer structure for API type
type NameServer struct {
	Fqdn string `json:"fqdn"`
	Ipv6 string `json:"ipv6"`
	Ipv4 string `json:"ipv4"`
}

// Record structure for API type
type Record struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	Type         string `json:"type"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	DynamicDNS   bool   `json:"dynamicDns"`
	Failed       bool   `json:"failed"`
	GtdLocation  string `json:"gtdLocation"`
	HardLink     bool   `json:"hardLink"`
	TTL          int    `json:"ttl"`
	SourceID     int    `json:"sourceId"`
	MxLevel      int    `json:"mxLevel,omitempty"`
	Failover     bool   `json:"failover"`
	Monitor      bool   `json:"monitor"`
	Keywords     string `json:"keywords,omitempty"`
	Source       int    `json:"source"`
	Priority     int    `json:"priority,omitempty"`
	Port         int    `json:"port,omitempty"`
	Weight       int    `json:"weight,omitempty"`
	RedirectType string `json:"redirectType,omitempty"`
}
