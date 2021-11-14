package sentry

import "time"

// User represents a Sentry User.
// https://github.com/getsentry/sentry/blob/275e6efa0f364ce05d9bfd09386b895b8a5e0671/src/sentry/api/serializers/models/user.py#L35
type User struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Username        string      `json:"username"`
	Email           string      `json:"email"`
	AvatarURL       string      `json:"avatarUrl"`
	IsActive        bool        `json:"isActive"`
	HasPasswordAuth bool        `json:"hasPasswordAuth"`
	IsManaged       bool        `json:"isManaged"`
	DateJoined      time.Time   `json:"dateJoined"`
	LastLogin       time.Time   `json:"lastLogin"`
	Has2FA          bool        `json:"has2fa"`
	LastActive      time.Time   `json:"lastActive"`
	IsSuperuser     bool        `json:"isSuperuser"`
	IsStaff         bool        `json:"isStaff"`
	Avatar          UserAvatar  `json:"avatar"`
	Emails          []UserEmail `json:"emails"`
}

// UserAvatar represents a user's avatar.
type UserAvatar struct {
	AvatarType string  `json:"avatarType"`
	AvatarUUID *string `json:"avatarUuid"`
}

// UserEmail represents a user's email and its verification status.
type UserEmail struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}
