package sentry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NotificationActionsService service

type CreateNotificationActionParams struct {
	TriggerType      *string      `json:"triggerType"`
	ServiceType      *string      `json:"serviceType"`
	IntegrationId    *json.Number `json:"integrationId,omitempty"`
	TargetIdentifier interface{}  `json:"targetIdentifier,omitempty"`
	TargetDisplay    *string      `json:"targetDisplay,omitempty"`
	TargetType       *string      `json:"targetType,omitempty"`
	Projects         []string     `json:"projects"`
}

type NotificationAction struct {
	ID               *json.Number  `json:"id"`
	TriggerType      *string       `json:"triggerType"`
	ServiceType      *string       `json:"serviceType"`
	IntegrationId    *json.Number  `json:"integrationId"`
	TargetIdentifier interface{}   `json:"targetIdentifier"`
	TargetDisplay    *string       `json:"targetDisplay"`
	TargetType       *string       `json:"targetType"`
	Projects         []json.Number `json:"projects"`
}

func (s *NotificationActionsService) Get(ctx context.Context, organizationSlug string, actionId string) (*NotificationAction, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/notifications/actions/%v/", organizationSlug, actionId)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	action := &NotificationAction{}
	resp, err := s.client.Do(ctx, req, action)
	if err != nil {
		return nil, resp, err
	}
	return action, resp, nil
}

func (s *NotificationActionsService) Create(ctx context.Context, organizationSlug string, params *CreateNotificationActionParams) (*NotificationAction, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/notifications/actions/", organizationSlug)
	req, err := s.client.NewRequest(http.MethodPost, u, params)
	if err != nil {
		return nil, nil, err
	}

	action := &NotificationAction{}
	resp, err := s.client.Do(ctx, req, action)
	if err != nil {
		return nil, resp, err
	}
	return action, resp, nil
}

type UpdateNotificationActionParams = CreateNotificationActionParams

func (s *NotificationActionsService) Update(ctx context.Context, organizationSlug string, actionId string, params *UpdateNotificationActionParams) (*NotificationAction, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/notifications/actions/%v/", organizationSlug, actionId)
	req, err := s.client.NewRequest(http.MethodPut, u, params)
	if err != nil {
		return nil, nil, err
	}

	action := &NotificationAction{}
	resp, err := s.client.Do(ctx, req, action)
	if err != nil {
		return nil, resp, err
	}
	return action, resp, nil
}

func (s *NotificationActionsService) Delete(ctx context.Context, organizationSlug string, actionId string) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%v/notifications/actions/%v/", organizationSlug, actionId)
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
