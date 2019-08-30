package api

type TagsListRequest struct {
	Key string `json:"key"`
}

type TagsDeleteRequest struct {
	Key string `json:"key"`
	Tag string `json:"tag"`
}

type TagsInfoRequest struct {
	Key string `json:"key"`
	Tag string `json:"tag"`
}

type TagsTimeSeriesRequest struct {
	Key string `json:"key"`
	Tag string `json:"tag"`
}

type TagsAllTimeSeriesRequest struct {
	Key string `json:"key"`
}
