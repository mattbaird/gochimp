package gochimp

import "fmt"

const (
	authorized_apps_path       = "/authorized-apps"
	single_authorized_app_path = authorized_apps_path + "/%s"
)

type ListOfAuthorizedApps struct {
	baseList `json:""`
	Apps     []AuthorizedApp `json:""`
}

type AuthorizedAppRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthorizedApp struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Users       []string `json:"users "`
	withLinks
}

type AuthorizedAppCreateResponse struct {
	AccessToken string `json:"access_token"`
	ViewerToken string `json:"viewer_token"`
}

func (api ChimpAPI) GetAuthorizedApps(params *ExtendedQueryParams) (*ListOfAuthorizedApps, error) {
	response := new(ListOfAuthorizedApps)

	err := api.Request("GET", authorized_apps_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api ChimpAPI) CreateAuthorizedApp(body *AuthorizedAppRequest) (*AuthorizedAppCreateResponse, error) {
	response := new(AuthorizedAppCreateResponse)

	err := api.Request("GET", authorized_apps_path, nil, body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api ChimpAPI) GetAuthroizedApp(id string, params *BasicQueryParams) (*AuthorizedApp, error) {
	response := new(AuthorizedApp)
	endpoint := fmt.Sprintf(single_authorized_app_path, id)

	err := api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
