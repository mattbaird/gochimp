package api

type TemplatesAddRequest struct {
	Key       string   `json:"key"`
	Name      string   `json:"name"`
	FromEmail string   `json:"from_email,omitempty"`
	FromName  string   `json:"from_name,omitempty"`
	Subject   string   `json:"subject,omitempty"`
	Code      string   `json:"code,omitempty"`
	Text      string   `json:"text,omitempty"`
	Publish   bool     `json:"publish,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type TemplatesInfoRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TemplatesUpdateRequest struct {
	Key       string   `json:"key"`
	Name      string   `json:"name"`
	FromEmail string   `json:"from_email,omitempty"`
	FromName  string   `json:"from_name,omitempty"`
	Subject   string   `json:"subject,omitempty"`
	Code      string   `json:"code,omitempty"`
	Text      string   `json:"text,omitempty"`
	Publish   bool     `json:"publish,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type TemplatesPublishRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TemplatesDeleteRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TemplatesListRequest struct {
	Key   string `json:"key"`
	Label string `json:"label,omitempty"`
}

type TemplatesTimeSeriesRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type TemplatesRenderRequest struct {
	Key             string `json:"key"`
	TemplateName    string `json:"template_name"`
	TemplateContent []struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	} `json:"template_content"`
	MergeVars []struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	} `json:"merge_vars,omitempty"`
}
