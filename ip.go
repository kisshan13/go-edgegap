// IP Lookup
// IP API - IP addresses related operations.
// API Reference - https://docs.edgegap.com/api/#tag/IP-Lookup

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PublicIPResponse struct {
	IP string `json:"public_ip,omitempty"` // Public IP Address
}

type IPBulkInfoPayload struct {
	Addresses []string `json:"addresses"`
}

type IPBulkInfo struct {
	Addresses []IPInformation `json:"addresses"`
}

type IPAddressLookupLocation struct {
	Continent IPAddressLookupLocationMeta `json:"continent"`
	Country   IPAddressLookupLocationMeta `json:"country"`
	Latitude  int                         `json:"latitude"`
	Longitude int                         `json:"longitude"`
}

type IPAddressLookupLocationMeta struct {
	Code string `json:"code"` // Country Code or Continent Code
	Name string `json:"name"` // Country Name or Continent Name
}

type IPInformation struct {
	Type      string                  `json:"type,omitempty"`
	IPAddress string                  `json:"ip_address,omitempty"`
	Location  IPAddressLookupLocation `json:"location,omitempty"`
}

// Retrieve your public IP address.
func (e *EdgegapClient) IPGet() (*Response[PublicIPResponse], error) {
	var response PublicIPResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/ip")
	}, &response)
}

// Lookup an IP address and return the associated information.
func (e *EdgegapClient) IPGetInfo(ip string) (*Response[IPInformation], error) {
	var response IPInformation

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/ip/%s/lookup", ip))
	}, &response)
}

// Lookup IP addresses and return the associated information. Maximum of 20 IPs.
func (e *EdgegapClient) IPGetInfoBulk(payload IPBulkInfoPayload) (*Response[IPBulkInfo], error) {
	var response IPBulkInfo

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Post(fmt.Sprintf("/"))
	}, &response)
}
