package pipedream

type source struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ConfiguredProps struct {
		HTTPInterface struct {
			EndpointURL string `json:"endpoint_url"`
		} `json:"httpInterface"`
	} `json:"configured_props"`
}

type sourceListResponse struct {
	PaginationInfo paginationInfo `json:"pagination_info"`
	Data           []source       `json:"data"`
}

type sourceCreateResponse struct {
	Data source `json:"data"`
}

type webhookEvent struct {
	Event struct {
		BodyRaw string `json:"bodyRaw"`
	} `json:"event"`
}

type paginationInfo struct {
	TotalCount  int    `json:"total_count"`
	Count       int    `json:"count"`
	StartCursor string `json:"start_cursor"`
	EndCursor   string `json:"end_cursor"`
}
