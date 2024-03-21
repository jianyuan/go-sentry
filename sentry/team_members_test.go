package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamMembersService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/members/member_id/teams/team_slug/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "4502349234123",
			"slug": "ancient-gabelers",
			"name": "Ancient Gabelers",
			"dateCreated": "2023-05-31T19:47:53.621181Z",
			"isMember": true,
			"teamRole": "contributor",
			"flags": {
				"idp:provisioned": false
			},
			"access": [
				"alerts:read",
				"event:write",
				"project:read",
				"team:read",
				"member:read",
				"project:releases",
				"event:read",
				"org:read"
			],
			"hasAccess": true,
			"isPending": false,
			"memberCount": 3,
			"avatar": {
				"avatarType": "letter_avatar",
				"avatarUuid": null
			}
		}`)
	})

	ctx := context.Background()
	team, _, err := client.TeamMembers.Create(ctx, "organization_slug", "member_id", "team_slug")
	assert.NoError(t, err)

	expected := &TeamMember{
		ID:          String("4502349234123"),
		Slug:        String("ancient-gabelers"),
		Name:        String("Ancient Gabelers"),
		DateCreated: Time(mustParseTime("2023-05-31T19:47:53.621181Z")),
		IsMember:    Bool(true),
		TeamRole:    String("contributor"),
		Flags: map[string]bool{
			"idp:provisioned": false,
		},
		Access: []string{
			"alerts:read",
			"event:write",
			"project:read",
			"team:read",
			"member:read",
			"project:releases",
			"event:read",
			"org:read",
		},
		HasAccess:   Bool(true),
		IsPending:   Bool(false),
		MemberCount: Int(3),
		Avatar: &Avatar{
			UUID: nil,
			Type: "letter_avatar",
		},
	}
	assert.Equal(t, expected, team)
}

func TestTeamMembersService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/members/member_id/teams/team_slug/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "4502349234123",
			"slug": "ancient-gabelers",
			"name": "Ancient Gabelers",
			"dateCreated": "2023-05-31T19:47:53.621181Z",
			"isMember": false,
			"teamRole": null,
			"flags": {
				"idp:provisioned": false
			},
			"access": [
				"alerts:read",
				"event:write",
				"project:read",
				"team:read",
				"member:read",
				"project:releases",
				"event:read",
				"org:read"
			],
			"hasAccess": true,
			"isPending": false,
			"memberCount": 3,
			"avatar": {
				"avatarType": "letter_avatar",
				"avatarUuid": null
			}
		}`)
	})

	ctx := context.Background()
	team, _, err := client.TeamMembers.Delete(ctx, "organization_slug", "member_id", "team_slug")
	assert.NoError(t, err)

	expected := &TeamMember{
		ID:          String("4502349234123"),
		Slug:        String("ancient-gabelers"),
		Name:        String("Ancient Gabelers"),
		DateCreated: Time(mustParseTime("2023-05-31T19:47:53.621181Z")),
		IsMember:    Bool(false),
		TeamRole:    nil,
		Flags: map[string]bool{
			"idp:provisioned": false,
		},
		Access: []string{
			"alerts:read",
			"event:write",
			"project:read",
			"team:read",
			"member:read",
			"project:releases",
			"event:read",
			"org:read",
		},
		HasAccess:   Bool(true),
		IsPending:   Bool(false),
		MemberCount: Int(3),
		Avatar: &Avatar{
			UUID: nil,
			Type: "letter_avatar",
		},
	}
	assert.Equal(t, expected, team)
}
