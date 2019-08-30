package api

type SendersListRequest struct {
	Key string `json:"key"`
}

type SendersDomainsRequest struct {
	Key string `json:"key"`
}

type SendersAddDomainRequest struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
}

type SendersCheckDomainRequest struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
}

type SendersVerifyDomainRequest struct {
	Key     string `json:"key"`
	Domain  string `json:"domain"`
	Mailbox string `json:"mailbox"`
}

type SendersInfoRequest struct {
	Key     string `json:"key"`
	Address string `json:"address"`
}

type SendersTimeSeriesRequest struct {
	Key     string `json:"key"`
	Address string `json:"address"`
}
