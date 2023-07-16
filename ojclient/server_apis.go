package ojclient

import "github.com/gookit/goutil/netutil/httpreq"

// Register client to server
func (c *Client) Register() error {
	reqData := &RegisterReq{}
	reqUrl := c.BuildAPI(ServerRegisterPath)

	resp, err := httpreq.Post(reqUrl, reqData, httpreq.WithJSONType)
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
	reqData := &UnregisterReq{}
	reqUrl := c.BuildAPI(ServerUnregisterPath)

	resp, err := httpreq.Post(reqUrl, reqData, httpreq.WithJSONType)
	if err != nil {
		return err
	}

	respData := &UnregisterResp{}
	err = httpreq.NewResp(resp).BindJSON(respData)
	if err != nil {
		return err
	}

	return nil
}

// Heartbeat send heartbeat to server
func (c *Client) Heartbeat() error {
	reqData := &HeartbeatReq{}
	reqUrl := c.BuildAPI(ServerHeartbeatPath)

	resp, err := httpreq.Post(reqUrl, reqData, httpreq.WithJSONType)
	if err != nil {
		return err
	}

	respData := &HeartbeatResp{}
	err = httpreq.NewResp(resp).BindJSON(respData)
	if err != nil {
		return err
	}

	return nil
}

// SubmitLogs submit task logs
func (c *Client) SubmitLogs() error {
	reqData := &SubmitLogsReq{}
	reqUrl := c.BuildAPI(ServerSubmitLogsPath)

	resp, err := httpreq.Post(reqUrl, reqData, httpreq.WithJSONType)
	if err != nil {
		return err
	}

	respData := &SubmitLogsResp{}
	err = httpreq.NewResp(resp).BindJSON(respData)
	if err != nil {
		return err
	}

	return nil
}

// FetchTopic client API
func (c *Client) FetchTopic() error {
	reqData := &FetchTopicReq{}
	reqUrl := c.BuildAPI(ServerFetchTopicPath)

	resp, err := httpreq.Post(reqUrl, reqData, httpreq.WithJSONType)
	if err != nil {
		return err
	}

	respData := &FetchTopicResp{}
	err = httpreq.NewResp(resp).BindJSON(respData)
	if err != nil {
		return err
	}

	return nil
}

// ReportStatus report task status
func (c *Client) ReportStatus() error {
	reqData := &ReportStatusReq{}
	reqUrl := c.BuildAPI(ServerReportStatusPath)

	resp, err := httpreq.Post(reqUrl, reqData, httpreq.WithJSONType)
	if err != nil {
		return err
	}

	respData := &ReportStatusResp{}
	err = httpreq.NewResp(resp).BindJSON(respData)
	if err != nil {
		return err
	}

	return nil
}
