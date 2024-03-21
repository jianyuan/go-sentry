package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectSymbolSourcesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/organization_slug/project_slug/symbol-sources/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodGet, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "27c5692e-de41-4087-bc14-74ed0fa421ba",
				"name": "s3",
				"bucket": "bucket",
				"region": "us-east-2",
				"access_key": "access_key",
				"type": "s3",
				"layout": {
					"casing": "default",
					"type": "native"
				},
				"secret_key": {
					"hidden-secret": true
				}
			},
				{
				"private_key": {
					"hidden-secret": true
				},
				"id": "f9df862d-45f7-496c-bf9b-ecade4c9f136",
				"layout": {
					"type": "native",
					"casing": "default"
				},
				"name": "gcs",
				"bucket": "gcs-bucket-name",
				"client_email": "test@example.com",
				"type": "gcs"
			},
			{
				"id": "1ccb6083-91ac-4394-a276-40fe0bb10ece",
				"name": "http",
				"url": "https://example.com",
				"layout": {
					"type": "native",
					"casing": "default"
				},
				"username": "admin",
				"password": {
					"hidden-secret": true
				},
				"type": "http"
			}
		]`)
	})

	ctx := context.Background()
	sources, _, err := client.ProjectSymbolSources.List(ctx, "organization_slug", "project_slug", nil)
	assert.NoError(t, err)

	expected := []*ProjectSymbolSource{
		{
			ID:   String("27c5692e-de41-4087-bc14-74ed0fa421ba"),
			Type: String("s3"),
			Name: String("s3"),
			Layout: &ProjectSymbolSourceLayout{
				Type:   String("native"),
				Casing: String("default"),
			},
			Bucket:    String("bucket"),
			Region:    String("us-east-2"),
			AccessKey: String("access_key"),
			SecretKey: &ProjectSymbolSourceHiddenSecret{
				HiddenSecret: Bool(true),
			},
		},
		{
			ID:   String("f9df862d-45f7-496c-bf9b-ecade4c9f136"),
			Type: String("gcs"),
			Name: String("gcs"),
			Layout: &ProjectSymbolSourceLayout{
				Type:   String("native"),
				Casing: String("default"),
			},
			Bucket:      String("gcs-bucket-name"),
			ClientEmail: String("test@example.com"),
			PrivateKey: &ProjectSymbolSourceHiddenSecret{
				HiddenSecret: Bool(true),
			},
		},
		{
			ID:   String("1ccb6083-91ac-4394-a276-40fe0bb10ece"),
			Type: String("http"),
			Name: String("http"),
			Layout: &ProjectSymbolSourceLayout{
				Type:   String("native"),
				Casing: String("default"),
			},
			Url:      String("https://example.com"),
			Username: String("admin"),
			Password: &ProjectSymbolSourceHiddenSecret{
				HiddenSecret: Bool(true),
			},
		},
	}
	assert.Equal(t, expected, sources)
}

func TestProjectSymbolSourcesService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/organization_slug/project_slug/symbol-sources/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodPost, r)
		assertPostJSON(t, map[string]interface{}{
			"name":       "s3",
			"bucket":     "bucket",
			"region":     "us-east-2",
			"access_key": "access_key",
			"type":       "s3",
			"layout": map[string]interface{}{
				"casing": "default",
				"type":   "native",
			},
			"secret_key": "secret_key",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "27c5692e-de41-4087-bc14-74ed0fa421ba",
			"name": "s3",
			"bucket": "bucket",
			"region": "us-east-2",
			"access_key": "access_key",
			"type": "s3",
			"layout": {
				"casing": "default",
				"type": "native"
			},
			"secret_key": {
				"hidden-secret": true
			}
		}`)
	})

	ctx := context.Background()
	params := &CreateProjectSymbolSourceParams{
		Type: String("s3"),
		Name: String("s3"),
		Layout: &ProjectSymbolSourceLayout{
			Type:   String("native"),
			Casing: String("default"),
		},
		Bucket:    String("bucket"),
		Region:    String("us-east-2"),
		AccessKey: String("access_key"),
		SecretKey: String("secret_key"),
	}
	source, _, err := client.ProjectSymbolSources.Create(ctx, "organization_slug", "project_slug", params)
	assert.NoError(t, err)

	expected := &ProjectSymbolSource{
		ID:   String("27c5692e-de41-4087-bc14-74ed0fa421ba"),
		Type: String("s3"),
		Name: String("s3"),
		Layout: &ProjectSymbolSourceLayout{
			Type:   String("native"),
			Casing: String("default"),
		},
		Bucket:    String("bucket"),
		Region:    String("us-east-2"),
		AccessKey: String("access_key"),
		SecretKey: &ProjectSymbolSourceHiddenSecret{
			HiddenSecret: Bool(true),
		},
	}
	assert.Equal(t, expected, source)
}

func TestProjectSymbolSourcesService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/organization_slug/project_slug/symbol-sources/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "27c5692e-de41-4087-bc14-74ed0fa421ba", r.URL.Query().Get("id"))
		assertMethod(t, http.MethodPut, r)
		assertPostJSON(t, map[string]interface{}{
			"id":         "27c5692e-de41-4087-bc14-74ed0fa421ba",
			"name":       "s3",
			"bucket":     "bucket",
			"region":     "us-east-2",
			"access_key": "access_key",
			"type":       "s3",
			"layout": map[string]interface{}{
				"casing": "default",
				"type":   "native",
			},
			"secret_key": "secret_key",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "27c5692e-de41-4087-bc14-74ed0fa421ba",
			"name": "s3",
			"bucket": "bucket",
			"region": "us-east-2",
			"access_key": "access_key",
			"type": "s3",
			"layout": {
				"casing": "default",
				"type": "native"
			},
			"secret_key": {
				"hidden-secret": true
			}
		}`)
	})

	ctx := context.Background()
	params := &UpdateProjectSymbolSourceParams{
		ID:   String("27c5692e-de41-4087-bc14-74ed0fa421ba"),
		Type: String("s3"),
		Name: String("s3"),
		Layout: &ProjectSymbolSourceLayout{
			Type:   String("native"),
			Casing: String("default"),
		},
		Bucket:    String("bucket"),
		Region:    String("us-east-2"),
		AccessKey: String("access_key"),
		SecretKey: String("secret_key"),
	}
	source, _, err := client.ProjectSymbolSources.Update(ctx, "organization_slug", "project_slug", "27c5692e-de41-4087-bc14-74ed0fa421ba", params)
	assert.NoError(t, err)

	expected := &ProjectSymbolSource{
		ID:   String("27c5692e-de41-4087-bc14-74ed0fa421ba"),
		Type: String("s3"),
		Name: String("s3"),
		Layout: &ProjectSymbolSourceLayout{
			Type:   String("native"),
			Casing: String("default"),
		},
		Bucket:    String("bucket"),
		Region:    String("us-east-2"),
		AccessKey: String("access_key"),
		SecretKey: &ProjectSymbolSourceHiddenSecret{
			HiddenSecret: Bool(true),
		},
	}
	assert.Equal(t, expected, source)
}
