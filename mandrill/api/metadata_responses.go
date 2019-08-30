package api

type MetaDataResponse struct {
	Name         string `json:"name"`
	State        string `json:"state"`
	ViewTemplate string `json:"view_template"`
}

type MetaDataListResponse []MetaDataResponse

type MetaDataAddResponse MetaDataResponse

type MetaDataUpdateResponse MetaDataResponse

type MetaDataDeleteResponse MetaDataResponse
