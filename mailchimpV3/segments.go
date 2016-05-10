package mailchimpV3

import "fmt"

const (
	segments_path       = "/lists/%s/segments"
	single_segment_path = segments_path + "/%s"
)

type ListOfSegments struct {
	baseList

	Segments []Segment `json:"segments"`
	ListID   string    `json:"list_id"`
}

type SegmentRequest struct {
	Name          string         `json:"name"`
	StaticSegment []string       `json:"static_segments"`
	Options       SegmentOptions `json:"options"`
}

type Segment struct {
	SegmentRequest

	ID          string `json:"id"`
	MemberCount int    `json:"member_count"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	ListID      string `json:"list_id"`

	withLinks
}

type SegmentOptions struct {
	Match      string               `json:"match"`
	Conditions []SegmentConditional `json:"conditions"`
}

// SegmentConditional represents parameters to filter by
type SegmentConditional struct {
	Field string  `json:"field"`
	OP    string  `json:"op"`
	Value float64 `json:"value"`
}

type SegmentQueryParams struct {
	ExtendedQueryParams

	Type            string
	SinceCreatedAt  string
	BeforeCreatedAt string
	SinceUpdatedAt  string
	BeforeUpdatedAt string
}

func (q SegmentQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()

	m["type"] = q.Type
	m["since_created_at"] = q.SinceCreatedAt
	m["since_updated_at"] = q.SinceUpdatedAt
	m["before_created_at"] = q.BeforeCreatedAt
	m["before_updated_at"] = q.BeforeUpdatedAt

	return m
}

func (list ListResponse) GetSegments(params *SegmentQueryParams) (*ListOfSegments, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(segments_path, list.ID)
	response := new(ListOfSegments)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) GetSegment(id string, params *BasicQueryParams) (*Segment, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	response := new(Segment)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) CreateSegment(body *SegmentRequest) (*Segment, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(segments_path, list.ID)
	response := new(Segment)

	return response, list.api.Request("POST", endpoint, nil, &body, response)
}

func (list ListResponse) UpdateSegment(id string, body *SegmentRequest) (*Segment, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	response := new(Segment)

	return response, list.api.Request("PATCH", endpoint, nil, &body, response)
}

func (list ListResponse) DeleteSegment(id string) (bool, error) {
	if err := list.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	return list.api.Do("DELETE", endpoint)
}
