package sentry

// https://github.com/getsentry/sentry/blob/23.12.1/src/sentry/api/serializers/models/role.py#L62-L74
type OrganizationRoleListItem struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Desc            string   `json:"desc"`
	Scopes          []string `json:"scopes"`
	IsAllowed       bool     `json:"isAllowed"`
	IsRetired       bool     `json:"isRetired"`
	IsGlobal        bool     `json:"isGlobal"`
	MinimumTeamRole string   `json:"minimumTeamRole"`
}

// https://github.com/getsentry/sentry/blob/23.12.1/src/sentry/api/serializers/models/role.py#L77-L85
type TeamRoleListItem struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Desc             string   `json:"desc"`
	Scopes           []string `json:"scopes"`
	IsAllowed        bool     `json:"isAllowed"`
	IsRetired        bool     `json:"isRetired"`
	IsMinimumRoleFor *string  `json:"isMinimumRoleFor"`
}
