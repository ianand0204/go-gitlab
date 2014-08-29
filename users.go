package gogitlab

// import "fmt"

type UsersService struct {
	client *Client
}

type User struct {
	Id            int    `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Name          string `json:"name,omitempty"`
	State         string `json:"state,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	Bio           string `json:"bio,omitempty"`
	Skype         string `json:"skype,omitempty"`
	LinkedIn      string `json:"linkedin,omitempty"`
	Twitter       string `json:"twitter,omitempty"`
	ExternUid     string `json:"extern_uid,omitempty"`
	Provider      string `json:"provider,omitempty"`
	ThemeId       int    `json:"theme_id,omitempty"`
	ColorSchemeId int    `json:"color_scheme_id,omitempty"`
	Password      string `json:"password"`
	Admin         bool   `json:"admin,omitempty"`
	CreateGroup   bool   `json:"can_create_group,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

func (s *UsersService) Users(opt *UserListOptions ) ([]User, *Response, error) {
	u, err := addOptions("users", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := new([]User)
	resp, err := s.client.Do(req, users)
	if err != nil {
		return nil, nil, err
	}

	return *users, resp, err
}

type UserListOptions struct {
  Since int `url:"since,omitempty"`
}
