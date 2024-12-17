package dto

type AsanaUsersListResponse struct {
	Data     []UserDataItem `json:"data"`
	NextPage NextPage       `json:"next_page"`
}

type UserDataItem struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type NextPage struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}
