package api

type TagsTagResponse struct {
	Tag string `json:"tag"`
	StatResponse
}
type TagsDeleteResponse TagsTagResponse
type TagsListResponse []TagsTagResponse

type TagsInfoResponse struct {
	TagsTagResponse
	Stats map[string]StatResponse
}

type TagsTimeSeriesResponse []struct {
	Time Time `json:"time"`
	StatResponse
}

type TagsAllTimeSeriesResponse []struct {
	Time Time `json:"time"`
	StatResponse
}
