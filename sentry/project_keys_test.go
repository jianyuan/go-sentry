package sentry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectKeysService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Link", "</api/0/projects/the-interstellar-jurisdiction/pump-station/keys/?&cursor=0:0:1>; rel=\"previous\"; results=\"true\"; cursor=\"0:0:1\", </api/0/projects/the-interstellar-jurisdiction/pump-station/keys/?&cursor=1584513610301:0:1>; rel=\"next\"; results=\"false\"; cursor=\"1584513610301:0:1\"")
		fmt.Fprint(w, `[
			{
				"id": "60120449b6b1d5e45f75561e6dabd80b",
				"name": "Liked Pegasus",
				"label": "Liked Pegasus",
				"public": "60120449b6b1d5e45f75561e6dabd80b",
				"secret": "189485c3b8ccf582bf5e12c530ef8858",
				"projectId": 4505281256090153,
				"isActive": true,
				"rateLimit": {
					"window": 7200,
					"count": 1000
				},
				"dsn": {
					"secret": "https://a785682ddda742d7a8a4088810e67701:bcd99b3790b3441c85ce4b1eaa854f66@o4504765715316736.ingest.sentry.io/4505281256090153",
					"public": "https://a785682ddda742d7a8a4088810e67791@o4504765715316736.ingest.sentry.io/4505281256090153",
					"csp": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/csp-report/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"security": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/security/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"minidump": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/minidump/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"nel": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/nel/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"unreal": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/unreal/a785682ddda719b7a8a4011110d75598/",
					"cdn": "https://js.sentry-cdn.com/a785682ddda719b7a8a4011110d75598.min.js",
					"crons": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/crons/___MONITOR_SLUG___/a785682ddda719b7a8a4011110d75598/"
				},
				"browserSdkVersion": "7.x",
				"browserSdk": {
					"choices": [
						[
							"latest",
							"latest"
						],
						[
							"7.x",
							"7.x"
						]
					]
				},
				"dateCreated": "2023-06-21T19:50:26.036254Z",
				"dynamicSdkLoaderOptions": {
					"hasReplay": true,
					"hasPerformance": true,
					"hasDebug": true
				}
			},
			{
				"id": "da8d69cb17e80677b76e08fde4656b93",
				"name": "Bold Oarfish",
				"label": "Bold Oarfish",
				"public": "da8d69cb17e80677b76e08fde4656b93",
				"secret": "5c241ebc42ccfbec281cbefbedc7ab96",
				"projectId": 4505281256090153,
				"isActive": true,
				"rateLimit": null,
				"dsn": {
					"secret": "https://a785682ddda742d7a8a4088810e67701:bcd99b3790b3441c85ce4b1eaa854f66@o4504765715316736.ingest.sentry.io/4505281256090153",
					"public": "https://a785682ddda742d7a8a4088810e67791@o4504765715316736.ingest.sentry.io/4505281256090153",
					"csp": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/csp-report/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"security": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/security/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"minidump": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/minidump/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"nel": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/nel/?sentry_key=a785682ddda719b7a8a4011110d75598",
					"unreal": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/unreal/a785682ddda719b7a8a4011110d75598/",
					"cdn": "https://js.sentry-cdn.com/a785682ddda719b7a8a4011110d75598.min.js",
					"crons": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/crons/___MONITOR_SLUG___/a785682ddda719b7a8a4011110d75598/"
				},
				"browserSdkVersion": "7.x",
				"browserSdk": {
					"choices": [
						[
							"latest",
							"latest"
						],
						[
							"7.x",
							"7.x"
						]
					]
				},
				"dateCreated": "2023-06-21T19:50:26.036254Z",
				"dynamicSdkLoaderOptions": {
					"hasReplay": true,
					"hasPerformance": true,
					"hasDebug": true
				}
			}
		]`)
	})

	ctx := context.Background()
	projectKeys, _, err := client.ProjectKeys.List(ctx, "the-interstellar-jurisdiction", "pump-station", nil)
	assert.NoError(t, err)

	expected := []*ProjectKey{
		{
			ID:        "60120449b6b1d5e45f75561e6dabd80b",
			Name:      "Liked Pegasus",
			Label:     "Liked Pegasus",
			Public:    "60120449b6b1d5e45f75561e6dabd80b",
			Secret:    "189485c3b8ccf582bf5e12c530ef8858",
			ProjectID: json.Number("4505281256090153"),
			IsActive:  true,
			RateLimit: &ProjectKeyRateLimit{
				Window: 7200,
				Count:  1000,
			},
			DSN: ProjectKeyDSN{
				Secret:   "https://a785682ddda742d7a8a4088810e67701:bcd99b3790b3441c85ce4b1eaa854f66@o4504765715316736.ingest.sentry.io/4505281256090153",
				Public:   "https://a785682ddda742d7a8a4088810e67791@o4504765715316736.ingest.sentry.io/4505281256090153",
				CSP:      "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/csp-report/?sentry_key=a785682ddda719b7a8a4011110d75598",
				Security: "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/security/?sentry_key=a785682ddda719b7a8a4011110d75598",
				Minidump: "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/minidump/?sentry_key=a785682ddda719b7a8a4011110d75598",
				NEL:      "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/nel/?sentry_key=a785682ddda719b7a8a4011110d75598",
				Unreal:   "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/unreal/a785682ddda719b7a8a4011110d75598/",
				CDN:      "https://js.sentry-cdn.com/a785682ddda719b7a8a4011110d75598.min.js",
				Crons:    "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/crons/___MONITOR_SLUG___/a785682ddda719b7a8a4011110d75598/",
			},
			BrowserSDKVersion: "7.x",
			DateCreated:       mustParseTime("2023-06-21T19:50:26.036254Z"),
			DynamicSDKLoaderOptions: ProjectKeyDynamicSDKLoaderOptions{
				HasReplay:      true,
				HasPerformance: true,
				HasDebugFiles:  true,
			},
		},
		{
			ID:        "da8d69cb17e80677b76e08fde4656b93",
			Name:      "Bold Oarfish",
			Label:     "Bold Oarfish",
			Public:    "da8d69cb17e80677b76e08fde4656b93",
			Secret:    "5c241ebc42ccfbec281cbefbedc7ab96",
			ProjectID: json.Number("4505281256090153"),
			IsActive:  true,
			RateLimit: nil,
			DSN: ProjectKeyDSN{
				Secret:   "https://a785682ddda742d7a8a4088810e67701:bcd99b3790b3441c85ce4b1eaa854f66@o4504765715316736.ingest.sentry.io/4505281256090153",
				Public:   "https://a785682ddda742d7a8a4088810e67791@o4504765715316736.ingest.sentry.io/4505281256090153",
				CSP:      "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/csp-report/?sentry_key=a785682ddda719b7a8a4011110d75598",
				Security: "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/security/?sentry_key=a785682ddda719b7a8a4011110d75598",
				Minidump: "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/minidump/?sentry_key=a785682ddda719b7a8a4011110d75598",
				NEL:      "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/nel/?sentry_key=a785682ddda719b7a8a4011110d75598",
				Unreal:   "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/unreal/a785682ddda719b7a8a4011110d75598/",
				CDN:      "https://js.sentry-cdn.com/a785682ddda719b7a8a4011110d75598.min.js",
				Crons:    "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/crons/___MONITOR_SLUG___/a785682ddda719b7a8a4011110d75598/",
			},
			BrowserSDKVersion: "7.x",
			DateCreated:       mustParseTime("2023-06-21T19:50:26.036254Z"),
			DynamicSDKLoaderOptions: ProjectKeyDynamicSDKLoaderOptions{
				HasReplay:      true,
				HasPerformance: true,
				HasDebugFiles:  true,
			},
		},
	}
	assert.Equal(t, expected, projectKeys)
}

func TestProjectKeysService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/60120449b6b1d5e45f75561e6dabd80b/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "60120449b6b1d5e45f75561e6dabd80b",
			"name": "Liked Pegasus",
			"label": "Liked Pegasus",
			"public": "60120449b6b1d5e45f75561e6dabd80b",
			"secret": "189485c3b8ccf582bf5e12c530ef8858",
			"projectId": 4505281256090153,
			"isActive": true,
			"rateLimit": {
				"window": 7200,
				"count": 1000
			},
			"dsn": {
				"secret": "https://a785682ddda742d7a8a4088810e67701:bcd99b3790b3441c85ce4b1eaa854f66@o4504765715316736.ingest.sentry.io/4505281256090153",
				"public": "https://a785682ddda742d7a8a4088810e67791@o4504765715316736.ingest.sentry.io/4505281256090153",
				"csp": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/csp-report/?sentry_key=a785682ddda719b7a8a4011110d75598",
				"security": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/security/?sentry_key=a785682ddda719b7a8a4011110d75598",
				"minidump": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/minidump/?sentry_key=a785682ddda719b7a8a4011110d75598",
				"nel": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/nel/?sentry_key=a785682ddda719b7a8a4011110d75598",
				"unreal": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/unreal/a785682ddda719b7a8a4011110d75598/",
				"cdn": "https://js.sentry-cdn.com/a785682ddda719b7a8a4011110d75598.min.js",
				"crons": "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/crons/___MONITOR_SLUG___/a785682ddda719b7a8a4011110d75598/"
			},
			"browserSdkVersion": "7.x",
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"7.x",
						"7.x"
					]
				]
			},
			"dateCreated": "2023-06-21T19:50:26.036254Z",
			"dynamicSdkLoaderOptions": {
				"hasReplay": true,
				"hasPerformance": true,
				"hasDebug": true
			}
		}`)
	})

	ctx := context.Background()
	projectKey, _, err := client.ProjectKeys.Get(ctx, "the-interstellar-jurisdiction", "pump-station", "60120449b6b1d5e45f75561e6dabd80b")
	assert.NoError(t, err)

	expected := &ProjectKey{
		ID:        "60120449b6b1d5e45f75561e6dabd80b",
		Name:      "Liked Pegasus",
		Label:     "Liked Pegasus",
		Public:    "60120449b6b1d5e45f75561e6dabd80b",
		Secret:    "189485c3b8ccf582bf5e12c530ef8858",
		ProjectID: json.Number("4505281256090153"),
		IsActive:  true,
		RateLimit: &ProjectKeyRateLimit{
			Window: 7200,
			Count:  1000,
		},
		DSN: ProjectKeyDSN{
			Secret:   "https://a785682ddda742d7a8a4088810e67701:bcd99b3790b3441c85ce4b1eaa854f66@o4504765715316736.ingest.sentry.io/4505281256090153",
			Public:   "https://a785682ddda742d7a8a4088810e67791@o4504765715316736.ingest.sentry.io/4505281256090153",
			CSP:      "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/csp-report/?sentry_key=a785682ddda719b7a8a4011110d75598",
			Security: "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/security/?sentry_key=a785682ddda719b7a8a4011110d75598",
			Minidump: "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/minidump/?sentry_key=a785682ddda719b7a8a4011110d75598",
			NEL:      "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/nel/?sentry_key=a785682ddda719b7a8a4011110d75598",
			Unreal:   "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/unreal/a785682ddda719b7a8a4011110d75598/",
			CDN:      "https://js.sentry-cdn.com/a785682ddda719b7a8a4011110d75598.min.js",
			Crons:    "https://o4504765715316736.ingest.sentry.io/api/4505281256090153/crons/___MONITOR_SLUG___/a785682ddda719b7a8a4011110d75598/",
		},
		BrowserSDKVersion: "7.x",
		DateCreated:       mustParseTime("2023-06-21T19:50:26.036254Z"),
		DynamicSDKLoaderOptions: ProjectKeyDynamicSDKLoaderOptions{
			HasReplay:      true,
			HasPerformance: true,
			HasDebugFiles:  true,
		},
	}

	assert.Equal(t, expected, projectKey)
}

func TestProjectKeysService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Fabulous Key",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"4.x",
						"4.x"
					]
				]
			},
			"browserSdkVersion": "4.x",
			"dateCreated": "2018-09-20T15:48:07.397Z",
			"dsn": {
				"cdn": "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
				"csp": "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"minidump": "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"public": "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				"secret": "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				"security": "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00"
			},
			"id": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"isActive": true,
			"label": "Fabulous Key",
			"name": "Fabulous Key",
			"projectId": 2,
			"public": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"rateLimit": null,
			"secret": "a07dcd97aa56481f82aeabaed43ca448"
		}`)
	})

	params := &CreateProjectKeyParams{
		Name: "Fabulous Key",
	}
	ctx := context.Background()
	projectKey, _, err := client.ProjectKeys.Create(ctx, "the-interstellar-jurisdiction", "pump-station", params)
	assert.NoError(t, err)
	expected := &ProjectKey{
		ID:        "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Name:      "Fabulous Key",
		Label:     "Fabulous Key",
		Public:    "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Secret:    "a07dcd97aa56481f82aeabaed43ca448",
		ProjectID: json.Number("2"),
		IsActive:  true,
		DSN: ProjectKeyDSN{
			Secret:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
			Public:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
			CSP:      "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Security: "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Minidump: "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			CDN:      "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
		},
		BrowserSDKVersion: "4.x",
		DateCreated:       mustParseTime("2018-09-20T15:48:07.397Z"),
	}
	assert.Equal(t, expected, projectKey)
}

func TestProjectKeysService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/befdbf32724c4ae0a3d286717b1f8127/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Fabulous Key",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"4.x",
						"4.x"
					]
				]
			},
			"browserSdkVersion": "4.x",
			"dateCreated": "2018-09-20T15:48:07.397Z",
			"dsn": {
				"cdn": "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
				"csp": "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"minidump": "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"public": "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				"secret": "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				"security": "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00"
			},
			"id": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"isActive": true,
			"label": "Fabulous Key",
			"name": "Fabulous Key",
			"projectId": 2,
			"public": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"rateLimit": null,
			"secret": "a07dcd97aa56481f82aeabaed43ca448"
		}`)
	})

	params := &UpdateProjectKeyParams{
		Name: "Fabulous Key",
	}
	ctx := context.Background()
	projectKey, _, err := client.ProjectKeys.Update(ctx, "the-interstellar-jurisdiction", "pump-station", "befdbf32724c4ae0a3d286717b1f8127", params)
	assert.NoError(t, err)
	expected := &ProjectKey{
		ID:        "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Name:      "Fabulous Key",
		Label:     "Fabulous Key",
		Public:    "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Secret:    "a07dcd97aa56481f82aeabaed43ca448",
		ProjectID: json.Number("2"),
		IsActive:  true,
		DSN: ProjectKeyDSN{
			Secret:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
			Public:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
			CSP:      "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Security: "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Minidump: "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			CDN:      "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
		},
		BrowserSDKVersion: "4.x",
		DateCreated:       mustParseTime("2018-09-20T15:48:07.397Z"),
	}
	assert.Equal(t, expected, projectKey)
}

func TestProjectKeysService_Update_RateLimit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/befdbf32724c4ae0a3d286717b1f8127/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Fabulous Key",
			"rateLimit": map[string]interface{}{
				"window": json.Number("86400"),
				"count":  json.Number("1000"),
			},
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"browserSdk": {
				"choices": [
					[
						"latest",
						"latest"
					],
					[
						"4.x",
						"4.x"
					]
				]
			},
			"browserSdkVersion": "4.x",
			"dateCreated": "2018-09-20T15:48:07.397Z",
			"dsn": {
				"cdn": "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
				"csp": "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"minidump": "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
				"public": "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
				"secret": "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
				"security": "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00"
			},
			"id": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"isActive": true,
			"label": "Fabulous Key",
			"name": "Fabulous Key",
			"projectId": 2,
			"public": "cfc7b0341c6e4f6ea1a9d256a30dba00",
			"rateLimit": {
				"count": 1000,
				"window": 86400
			},
			"secret": "a07dcd97aa56481f82aeabaed43ca448"
		}`)
	})

	rateLimit := ProjectKeyRateLimit{
		Count:  1000,
		Window: 86400,
	}
	params := &UpdateProjectKeyParams{
		Name:      "Fabulous Key",
		RateLimit: &rateLimit,
	}
	ctx := context.Background()
	projectKey, _, err := client.ProjectKeys.Update(ctx, "the-interstellar-jurisdiction", "pump-station", "befdbf32724c4ae0a3d286717b1f8127", params)
	assert.NoError(t, err)
	expected := &ProjectKey{
		ID:        "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Name:      "Fabulous Key",
		Label:     "Fabulous Key",
		Public:    "cfc7b0341c6e4f6ea1a9d256a30dba00",
		Secret:    "a07dcd97aa56481f82aeabaed43ca448",
		ProjectID: json.Number("2"),
		IsActive:  true,
		RateLimit: &rateLimit,
		DSN: ProjectKeyDSN{
			Secret:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00:a07dcd97aa56481f82aeabaed43ca448@sentry.io/2",
			Public:   "https://cfc7b0341c6e4f6ea1a9d256a30dba00@sentry.io/2",
			CSP:      "https://sentry.io/api/2/csp-report/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Security: "https://sentry.io/api/2/security/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			Minidump: "https://sentry.io/api/2/minidump/?sentry_key=cfc7b0341c6e4f6ea1a9d256a30dba00",
			CDN:      "https://sentry.io/js-sdk-loader/cfc7b0341c6e4f6ea1a9d256a30dba00.min.js",
		},
		BrowserSDKVersion: "4.x",
		DateCreated:       mustParseTime("2018-09-20T15:48:07.397Z"),
	}
	assert.Equal(t, expected, projectKey)
}

func TestProjectKeysService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/keys/befdbf32724c4ae0a3d286717b1f8127/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
	})

	ctx := context.Background()
	_, err := client.ProjectKeys.Delete(ctx, "the-interstellar-jurisdiction", "pump-station", "befdbf32724c4ae0a3d286717b1f8127")
	assert.NoError(t, err)

}
