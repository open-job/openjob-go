package ojclient

import "github.com/gookit/goutil/netutil/httpreq"

// Register client to server
func (c *Client) Register() error {
	reqData := &RegisterReq{}

	resp, err := httpreq.Post(c.BuildAPI(ServerRegisterPath), reqData, httpreq.WithJSONType)
	if err != nil {
		return err
	}

	respData := &RegisterResp{}
	err = httpreq.NewResp(resp).BindJSON(respData)
	if err != nil {
		return err
	}

	return nil
}

// Unregister client from server
func (c *Client) Unregister() error {
	return nil
}

// Heartbeat send heartbeat to server
func (c *Client) Heartbeat() error {
	return nil
}

// SubmitLogs submit task logs
func (c *Client) SubmitLogs() error {
	return nil
}

// FetchTopic client API
func (c *Client) FetchTopic() error {
	return nil
}

// ReportStatus report task status
func (c *Client) ReportStatus() error {
	return nil
}
