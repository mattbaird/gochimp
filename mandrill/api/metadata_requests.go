package api

type MetaDataListRequest struct {
	Key string `json:"key"`
}

type MetaDataAddRequest struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	ViewTemplate string `json:"view_template,omitempty"`
}

type MetaDataUpdateRequest struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	ViewTemplate string `json:"view_template"`
}

type MetaDataDeleteRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}
