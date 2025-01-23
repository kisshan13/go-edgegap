package edgegap

import "github.com/go-resty/resty/v2"

type EdgegapClient struct {
	client *resty.Client
}

func NewEdgegapClient(token string) *EdgegapClient {

	client := resty.New()

	client.SetHeaders(
		map[string]string{
			"Content-Type":  "application/json",
			"Authorization": token,
		},
	)

	client.SetBaseURL(EDGEGAP_BASE_URL + "/" + string(VersionOne))

	return &EdgegapClient{
		client: client,
	}
}
