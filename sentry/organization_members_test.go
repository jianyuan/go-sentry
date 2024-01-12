package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationMembersService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"cursor": "100:-1:1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"inviteStatus": "approved",
				"dateCreated": "2020-01-04T00:00:00.000000Z",
				"user": {
					"username": "test@example.com",
					"lastLogin": "2020-01-02T00:00:00.000000Z",
					"isSuperuser": false,
					"emails": [
						{
							"is_verified": true,
							"id": "1",
							"email": "test@example.com"
						}
					],
					"isManaged": false,
					"experiments": {},
					"lastActive": "2020-01-03T00:00:00.000000Z",
					"isStaff": false,
					"identities": [],
					"id": "1",
					"isActive": true,
					"has2fa": false,
					"name": "John Doe",
					"avatarUrl": "https://secure.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=32&d=mm",
					"dateJoined": "2020-01-01T00:00:00.000000Z",
					"options": {
						"timezone": "UTC",
						"stacktraceOrder": -1,
						"language": "en",
						"clock24Hours": false
					},
					"flags": {
						"newsletter_consent_prompt": false
					},
					"avatar": {
						"avatarUuid": null,
						"avatarType": "letter_avatar"
					},
					"hasPasswordAuth": true,
					"email": "test@example.com"
				},
				"roleName": "Owner",
				"expired": false,
				"id": "1",
				"inviterName": null,
				"name": "John Doe",
				"role": "owner",
				"flags": {
					"sso:linked": false,
					"sso:invalid": false
				},
				"email": "test@example.com",
				"pending": false
			}
		]`)
	})

	ctx := context.Background()
	members, _, err := client.OrganizationMembers.List(ctx, "the-interstellar-jurisdiction", &ListCursorParams{
		Cursor: "100:-1:1",
	})
	assert.NoError(t, err)
	expected := []*OrganizationMember{
		{
			ID:    "1",
			Email: "test@example.com",
			Name:  "John Doe",
			User: User{
				ID:              "1",
				Name:            "John Doe",
				Username:        "test@example.com",
				Email:           "test@example.com",
				AvatarURL:       "https://secure.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=32&d=mm",
				IsActive:        true,
				HasPasswordAuth: true,
				IsManaged:       false,
				DateJoined:      mustParseTime("2020-01-01T00:00:00.000000Z"),
				LastLogin:       mustParseTime("2020-01-02T00:00:00.000000Z"),
				Has2FA:          false,
				LastActive:      mustParseTime("2020-01-03T00:00:00.000000Z"),
				IsSuperuser:     false,
				IsStaff:         false,
				Avatar: Avatar{
					Type: "letter_avatar",
					UUID: nil,
				},
				Emails: []UserEmail{
					{
						ID:         "1",
						Email:      "test@example.com",
						IsVerified: true,
					},
				},
			},
			Pending: false,
			Expired: false,
			Flags: map[string]bool{
				"sso:invalid": false,
				"sso:linked":  false,
			},
			DateCreated:  mustParseTime("2020-01-04T00:00:00.000000Z"),
			InviteStatus: "approved",
			InviterName:  nil,
		},
	}
	assert.Equal(t, expected, members)
}

func TestOrganizationMembersService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
				"inviteStatus": "approved",
				"dateCreated": "2020-01-04T00:00:00.000000Z",
				"user": {
					"username": "test@example.com",
					"lastLogin": "2020-01-02T00:00:00.000000Z",
					"isSuperuser": false,
					"emails": [
						{
							"is_verified": true,
							"id": "1",
							"email": "test@example.com"
						}
					],
					"isManaged": false,
					"experiments": {},
					"lastActive": "2020-01-03T00:00:00.000000Z",
					"isStaff": false,
					"identities": [],
					"id": "1",
					"isActive": true,
					"has2fa": false,
					"name": "John Doe",
					"avatarUrl": "https://secure.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=32&d=mm",
					"dateJoined": "2020-01-01T00:00:00.000000Z",
					"options": {
						"timezone": "UTC",
						"stacktraceOrder": -1,
						"language": "en",
						"clock24Hours": false
					},
					"flags": {
						"newsletter_consent_prompt": false
					},
					"avatar": {
						"avatarUuid": null,
						"avatarType": "letter_avatar"
					},
					"hasPasswordAuth": true,
					"email": "test@example.com"
				},
				"roleName": "Owner",
				"expired": false,
				"id": "1",
				"inviterName": null,
				"name": "John Doe",
				"role": "owner",
				"flags": {
					"sso:linked": false,
					"sso:invalid": false
				},
				"teams": [],
				"email": "test@example.com",
				"pending": false
			}`)
	})

	ctx := context.Background()
	members, _, err := client.OrganizationMembers.Get(ctx, "the-interstellar-jurisdiction", "1")
	assert.NoError(t, err)
	expected := OrganizationMember{
		ID:    "1",
		Email: "test@example.com",
		Name:  "John Doe",
		User: User{
			ID:              "1",
			Name:            "John Doe",
			Username:        "test@example.com",
			Email:           "test@example.com",
			AvatarURL:       "https://secure.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=32&d=mm",
			IsActive:        true,
			HasPasswordAuth: true,
			IsManaged:       false,
			DateJoined:      mustParseTime("2020-01-01T00:00:00.000000Z"),
			LastLogin:       mustParseTime("2020-01-02T00:00:00.000000Z"),
			Has2FA:          false,
			LastActive:      mustParseTime("2020-01-03T00:00:00.000000Z"),
			IsSuperuser:     false,
			IsStaff:         false,
			Avatar: Avatar{
				Type: "letter_avatar",
				UUID: nil,
			},
			Emails: []UserEmail{
				{
					ID:         "1",
					Email:      "test@example.com",
					IsVerified: true,
				},
			},
		},
		Pending: false,
		Expired: false,
		Flags: map[string]bool{
			"sso:invalid": false,
			"sso:linked":  false,
		},
		Teams:        []string{},
		DateCreated:  mustParseTime("2020-01-04T00:00:00.000000Z"),
		InviteStatus: "approved",
		InviterName:  nil,
	}
	assert.Equal(t, &expected, members)
}

func TestOrganizationMembersService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "1",
			"email": "test@example.com",
			"name": "test@example.com",
			"user": null,
			"role": "member",
			"roleName": "Member",
			"pending": true,
			"expired": false,
			"flags": {
				"sso:linked": false,
				"sso:invalid": false,
				"member-limit:restricted": false
			},
			"teams": [],
			"dateCreated": "2020-01-01T00:00:00.000000Z",
			"inviteStatus": "approved",
			"inviterName": "John Doe"
		}`)
	})

	createOrganizationMemberParams := CreateOrganizationMemberParams{
		Email: "test@example.com",
		Role:  OrganizationRoleMember,
	}
	ctx := context.Background()
	member, _, err := client.OrganizationMembers.Create(ctx, "the-interstellar-jurisdiction", &createOrganizationMemberParams)
	assert.NoError(t, err)

	inviterName := "John Doe"
	expected := OrganizationMember{
		ID:      "1",
		Email:   "test@example.com",
		Name:    "test@example.com",
		User:    User{},
		Pending: true,
		Expired: false,
		Flags: map[string]bool{
			"sso:linked":              false,
			"sso:invalid":             false,
			"member-limit:restricted": false,
		},
		Teams:        []string{},
		DateCreated:  mustParseTime("2020-01-01T00:00:00.000000Z"),
		InviteStatus: "approved",
		InviterName:  &inviterName,
	}

	assert.Equal(t, &expected, member)
}

func TestOrganizationMembersService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "57377908164",
			"email": "sirpenguin@antarcticarocks.com",
			"name": "Sir Penguin",
			"user": {
				"id": "280094367316",
				"name": "Sir Penguin",
				"username": "sirpenguin@antarcticarocks.com",
				"email": "sirpenguin@antarcticarocks.com",
				"avatarUrl": "https://secure.gravatar.com/avatar/16aeb26c5fdba335c7078e9e9ddb5149?s=32&d=mm",
				"isActive": true,
				"hasPasswordAuth": true,
				"isManaged": false,
				"dateJoined": "2021-07-06T21:13:58.375239Z",
				"lastLogin": "2021-08-02T18:25:00.051182Z",
				"has2fa": false,
				"lastActive": "2021-08-02T21:32:18.836829Z",
				"isSuperuser": false,
				"isStaff": false,
				"experiments": {},
				"emails": [
					{
						"id": "2153450836",
						"email": "sirpenguin@antarcticarocks.com",
						"is_verified": true
					}
				],
				"avatar": {
					"avatarType": "letter_avatar",
					"avatarUuid": null
				},
				"authenticators": [],
				"canReset2fa": true
			},
			"role": "member",
			"orgRole": "member",
			"roleName": "Member",
			"pending": false,
			"expired": false,
			"flags": {
				"idp:provisioned": false,
				"idp:role-restricted": false,
				"sso:linked": false,
				"sso:invalid": false,
				"member-limit:restricted": false,
				"partnership:restricted": false
			},
			"dateCreated": "2021-07-06T21:13:01.120263Z",
			"inviteStatus": "approved",
			"inviterName": "maininviter@antarcticarocks.com",
			"teams": [
				"cool-team",
				"ancient-gabelers"
			],
			"teamRoles": [
				{
					"teamSlug": "ancient-gabelers",
					"role": "admin"
				},
				{
					"teamSlug": "powerful-abolitionist",
					"role": "contributor"
				}
			],
			"invite_link": null,
			"isOnlyOwner": false,
			"orgRoleList": [
				{
					"id": "billing",
					"name": "Billing",
					"desc": "Can manage subscription and billing details.",
					"scopes": [
						"org:billing"
					],
					"allowed": true,
					"isAllowed": true,
					"isRetired": false,
					"is_global": false,
					"isGlobal": false,
					"minimumTeamRole": "contributor"
				},
				{
					"id": "member",
					"name": "Member",
					"desc": "Members can view and act on events, as well as view most other data within the organization.",
					"scopes": [
						"team:read",
						"project:releases",
						"org:read",
						"event:read",
						"alerts:write",
						"member:read",
						"alerts:read",
						"event:admin",
						"project:read",
						"event:write"
					],
					"allowed": true,
					"isAllowed": true,
					"isRetired": false,
					"is_global": false,
					"isGlobal": false,
					"minimumTeamRole": "contributor"
				},
				{
					"id": "admin",
					"name": "Admin",
					"desc": "Admin privileges on any teams of which they're a member. They can create new teams and projects, as well as remove teams and projects on which they already hold membership (or all teams, if open membership is enabled). Additionally, they can manage memberships of teams that they are members of. They cannot invite members to the organization.",
					"scopes": [
						"team:admin",
						"org:integrations",
						"project:admin",
						"team:read",
						"project:releases",
						"org:read",
						"team:write",
						"event:read",
						"alerts:write",
						"member:read",
						"alerts:read",
						"event:admin",
						"project:read",
						"event:write",
						"project:write"
					],
					"allowed": true,
					"isAllowed": true,
					"isRetired": true,
					"is_global": false,
					"isGlobal": false,
					"minimumTeamRole": "admin"
				},
				{
					"id": "manager",
					"name": "Manager",
					"desc": "Gains admin access on all teams as well as the ability to add and remove members.",
					"scopes": [
						"team:admin",
						"org:integrations",
						"project:releases",
						"team:write",
						"member:read",
						"org:write",
						"project:write",
						"project:admin",
						"team:read",
						"org:read",
						"event:read",
						"member:write",
						"alerts:write",
						"alerts:read",
						"event:admin",
						"project:read",
						"event:write",
						"member:admin"
					],
					"allowed": true,
					"isAllowed": true,
					"isRetired": false,
					"is_global": true,
					"isGlobal": true,
					"minimumTeamRole": "admin"
				},
				{
					"id": "owner",
					"name": "Owner",
					"desc": "Unrestricted access to the organization, its data, and its settings. Can add, modify, and delete projects and members, as well as make billing and plan changes.",
					"scopes": [
						"team:admin",
						"org:integrations",
						"project:releases",
						"org:admin",
						"team:write",
						"member:read",
						"org:write",
						"project:write",
						"project:admin",
						"team:read",
						"org:read",
						"event:read",
						"member:write",
						"alerts:write",
						"org:billing",
						"alerts:read",
						"event:admin",
						"project:read",
						"event:write",
						"member:admin"
					],
					"allowed": true,
					"isAllowed": true,
					"isRetired": false,
					"is_global": true,
					"isGlobal": true,
					"minimumTeamRole": "admin"
				}
			],
			"teamRoleList": [
				{
					"id": "contributor",
					"name": "Contributor",
					"desc": "Contributors can view and act on events, as well as view most other data within the team's projects.",
					"scopes": [
						"team:read",
						"project:releases",
						"org:read",
						"event:read",
						"member:read",
						"alerts:read",
						"project:read",
						"event:write"
					],
					"allowed": false,
					"isAllowed": false,
					"isRetired": false,
					"isMinimumRoleFor": null
				},
				{
					"id": "admin",
					"name": "Team Admin",
					"desc": "Admin privileges on the team. They can create and remove projects, and can manage the team's memberships. They cannot invite members to the organization.",
					"scopes": [
						"team:admin",
						"org:integrations",
						"project:admin",
						"team:read",
						"project:releases",
						"org:read",
						"team:write",
						"event:read",
						"alerts:write",
						"member:read",
						"alerts:read",
						"event:admin",
						"project:read",
						"event:write",
						"project:write"
					],
					"allowed": false,
					"isAllowed": false,
					"isRetired": false,
					"isMinimumRoleFor": "admin"
				}
			]
		}`)
	})

	updateOrganizationMemberParams := UpdateOrganizationMemberParams{
		OrganizationRole: OrganizationRoleMember,
	}
	ctx := context.Background()
	member, _, err := client.OrganizationMembers.Update(ctx, "the-interstellar-jurisdiction", "1", &updateOrganizationMemberParams)
	assert.NoError(t, err)

	inviterName := "maininviter@antarcticarocks.com"
	expected := OrganizationMember{
		ID:    "57377908164",
		Email: "sirpenguin@antarcticarocks.com",
		Name:  "Sir Penguin",
		User: User{
			ID:              "280094367316",
			Name:            "Sir Penguin",
			Username:        "sirpenguin@antarcticarocks.com",
			Email:           "sirpenguin@antarcticarocks.com",
			AvatarURL:       "https://secure.gravatar.com/avatar/16aeb26c5fdba335c7078e9e9ddb5149?s=32&d=mm",
			IsActive:        true,
			HasPasswordAuth: true,
			IsManaged:       false,
			DateJoined:      mustParseTime("2021-07-06T21:13:58.375239Z"),
			LastLogin:       mustParseTime("2021-08-02T18:25:00.051182Z"),
			Has2FA:          false,
			LastActive:      mustParseTime("2021-08-02T21:32:18.836829Z"),
			IsSuperuser:     false,
			IsStaff:         false,
			Avatar: Avatar{
				Type: "letter_avatar",
				UUID: nil,
			},
			Emails: []UserEmail{
				{
					ID:         "2153450836",
					Email:      "sirpenguin@antarcticarocks.com",
					IsVerified: true,
				},
			},
		},
		OrgRole: OrganizationRoleMember,
		OrgRoleList: []OrganizationRoleListItem{
			{
				ID:              "billing",
				Name:            "Billing",
				Desc:            "Can manage subscription and billing details.",
				Scopes:          []string{"org:billing"},
				IsAllowed:       true,
				IsRetired:       false,
				IsGlobal:        false,
				MinimumTeamRole: "contributor",
			},
			{
				ID:              "member",
				Name:            "Member",
				Desc:            "Members can view and act on events, as well as view most other data within the organization.",
				Scopes:          []string{"team:read", "project:releases", "org:read", "event:read", "alerts:write", "member:read", "alerts:read", "event:admin", "project:read", "event:write"},
				IsAllowed:       true,
				IsRetired:       false,
				IsGlobal:        false,
				MinimumTeamRole: "contributor",
			},
			{
				ID:              "admin",
				Name:            "Admin",
				Desc:            "Admin privileges on any teams of which they're a member. They can create new teams and projects, as well as remove teams and projects on which they already hold membership (or all teams, if open membership is enabled). Additionally, they can manage memberships of teams that they are members of. They cannot invite members to the organization.",
				Scopes:          []string{"team:admin", "org:integrations", "project:admin", "team:read", "project:releases", "org:read", "team:write", "event:read", "alerts:write", "member:read", "alerts:read", "event:admin", "project:read", "event:write", "project:write"},
				IsAllowed:       true,
				IsRetired:       true,
				IsGlobal:        false,
				MinimumTeamRole: "admin",
			},
			{
				ID:              "manager",
				Name:            "Manager",
				Desc:            "Gains admin access on all teams as well as the ability to add and remove members.",
				Scopes:          []string{"team:admin", "org:integrations", "project:releases", "team:write", "member:read", "org:write", "project:write", "project:admin", "team:read", "org:read", "event:read", "member:write", "alerts:write", "alerts:read", "event:admin", "project:read", "event:write", "member:admin"},
				IsAllowed:       true,
				IsRetired:       false,
				IsGlobal:        true,
				MinimumTeamRole: "admin",
			},
			{
				ID:              "owner",
				Name:            "Owner",
				Desc:            "Unrestricted access to the organization, its data, and its settings. Can add, modify, and delete projects and members, as well as make billing and plan changes.",
				Scopes:          []string{"team:admin", "org:integrations", "project:releases", "org:admin", "team:write", "member:read", "org:write", "project:write", "project:admin", "team:read", "org:read", "event:read", "member:write", "alerts:write", "org:billing", "alerts:read", "event:admin", "project:read", "event:write", "member:admin"},
				IsAllowed:       true,
				IsRetired:       false,
				IsGlobal:        true,
				MinimumTeamRole: "admin",
			},
		},
		Pending: false,
		Expired: false,
		Flags: map[string]bool{
			"idp:provisioned":         false,
			"idp:role-restricted":     false,
			"sso:linked":              false,
			"sso:invalid":             false,
			"member-limit:restricted": false,
			"partnership:restricted":  false,
		},
		DateCreated:  mustParseTime("2021-07-06T21:13:01.120263Z"),
		InviteStatus: "approved",
		InviterName:  &inviterName,
		TeamRoleList: []TeamRoleListItem{
			{
				ID:               "contributor",
				Name:             "Contributor",
				Desc:             "Contributors can view and act on events, as well as view most other data within the team's projects.",
				Scopes:           []string{"team:read", "project:releases", "org:read", "event:read", "member:read", "alerts:read", "project:read", "event:write"},
				IsAllowed:        false,
				IsRetired:        false,
				IsMinimumRoleFor: nil,
			},
			{
				ID:               "admin",
				Name:             "Team Admin",
				Desc:             "Admin privileges on the team. They can create and remove projects, and can manage the team's memberships. They cannot invite members to the organization.",
				Scopes:           []string{"team:admin", "org:integrations", "project:admin", "team:read", "project:releases", "org:read", "team:write", "event:read", "alerts:write", "member:read", "alerts:read", "event:admin", "project:read", "event:write", "project:write"},
				IsAllowed:        false,
				IsRetired:        false,
				IsMinimumRoleFor: String("admin"),
			},
		},
		TeamRoles: []TeamRole{
			{
				TeamSlug: "ancient-gabelers",
				Role:     String(TeamRoleAdmin),
			},
			{
				TeamSlug: "powerful-abolitionist",
				Role:     String(TeamRoleContributor),
			},
		},
		Teams: []string{
			"cool-team",
			"ancient-gabelers",
		},
	}

	assert.Equal(t, &expected, member)
}

func TestOrganizationMembersService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	resp, err := client.OrganizationMembers.Delete(ctx, "the-interstellar-jurisdiction", "1")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), resp.ContentLength)
}
