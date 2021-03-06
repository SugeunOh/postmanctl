/*
Copyright © 2020 Kevin Swiber <kswiber@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sdk

import (
	"context"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
)

// Collections returns all collections.
func (s *Service) Collections(ctx context.Context) (*resources.CollectionListItems, error) {
	var resource resources.CollectionListResponse
	if _, err := s.get(ctx, &resource, nil, "collections"); err != nil {
		return nil, err
	}

	return &resource.Collections, nil
}

// Collection returns a single collection.
func (s *Service) Collection(ctx context.Context, id string) (*resources.Collection, error) {
	var resource resources.CollectionResponse
	if _, err := s.get(ctx, &resource, nil, "collections", id); err != nil {
		return nil, err
	}

	return &resource.Collection, nil
}

// Environments returns all environments.
func (s *Service) Environments(ctx context.Context) (*resources.EnvironmentListItems, error) {
	var resource resources.EnvironmentListResponse
	if _, err := s.get(ctx, &resource, nil, "environments"); err != nil {
		return nil, err
	}

	return &resource.Environments, nil
}

// Environment returns a single environment.
func (s *Service) Environment(ctx context.Context, id string) (*resources.Environment, error) {
	var resource resources.EnvironmentResponse
	if _, err := s.get(ctx, &resource, nil, "environments", id); err != nil {
		return nil, err
	}

	return &resource.Environment, nil
}

// APIs returns all APIs.
func (s *Service) APIs(ctx context.Context, workspace string) (*resources.APIListItems, error) {
	var resource resources.APIListResponse
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	if _, err := s.get(ctx, &resource, params, "apis"); err != nil {
		return nil, err
	}

	return &resource.APIs, nil
}

// API returns a single API.
func (s *Service) API(ctx context.Context, id string) (*resources.API, error) {
	var resource resources.APIResponse
	if _, err := s.get(ctx, &resource, nil, "apis", id); err != nil {
		return nil, err
	}

	return &resource.API, nil
}

// APIVersions returns all API Versions.
func (s *Service) APIVersions(ctx context.Context, apiID string) (*resources.APIVersionListItems, error) {
	var resource resources.APIVersionListResponse
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions"); err != nil {
		return nil, err
	}

	return &resource.APIVersions, nil
}

// APIVersion returns a single API Version.
func (s *Service) APIVersion(ctx context.Context, apiID, id string) (*resources.APIVersion, error) {
	var resource resources.APIVersionResponse
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions", id); err != nil {
		return nil, err
	}

	return &resource.APIVersion, nil
}

// Schema returns a single schema for an API version.
func (s *Service) Schema(ctx context.Context, apiID, apiVersionID, id string) (*resources.Schema, error) {
	var resource resources.SchemaResponse
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions", apiVersionID, "schemas", id); err != nil {
		return nil, err
	}

	return &resource.Schema, nil
}

// APIRelations returns the linked relations of an API
func (s *Service) APIRelations(ctx context.Context, apiID, apiVersionID string) (*resources.APIRelations, error) {
	var resource resources.APIRelationsResource
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions", apiVersionID, "relations"); err != nil {
		return nil, err
	}

	return &resource.Relations, nil
}

// FormattedAPIRelationItems returns the formatted linked relations of an API
func (s *Service) FormattedAPIRelationItems(ctx context.Context, apiID, apiVersionID string) (*resources.FormattedAPIRelationItems, error) {
	r, err := s.APIRelations(ctx, apiID, apiVersionID)
	if err != nil {
		return nil, err
	}

	f := resources.NewFormattedAPIRelationItems(r)
	return &f, nil
}

// User returns the current user.
func (s *Service) User(ctx context.Context) (*resources.User, error) {
	var resource resources.UserResponse
	if _, err := s.get(ctx, &resource, nil, "me"); err != nil {
		return nil, err
	}

	return &resource.User, nil
}

// Workspaces returns the workspaces for the current user.
func (s *Service) Workspaces(ctx context.Context) (*resources.WorkspaceListItems, error) {
	var resource resources.WorkspaceListResponse
	if _, err := s.get(ctx, &resource, nil, "workspaces"); err != nil {
		return nil, err
	}

	return &resource.Workspaces, nil
}

// Workspace returns a single workspace for the current user.
func (s *Service) Workspace(ctx context.Context, id string) (*resources.Workspace, error) {
	var resource resources.WorkspaceResponse
	if _, err := s.get(ctx, &resource, nil, "workspaces", id); err != nil {
		return nil, err
	}

	return &resource.Workspace, nil
}

// Monitors returns the monitors for the current user.
func (s *Service) Monitors(ctx context.Context) (*resources.MonitorListItems, error) {
	var resource resources.MonitorListResponse
	if _, err := s.get(ctx, &resource, nil, "monitors"); err != nil {
		return nil, err
	}

	return &resource.Monitors, nil
}

// Monitor returns a single monitor for the current user.
func (s *Service) Monitor(ctx context.Context, id string) (*resources.Monitor, error) {
	var resource resources.MonitorResponse
	if _, err := s.get(ctx, &resource, nil, "monitors", id); err != nil {
		return nil, err
	}

	return &resource.Monitor, nil
}

// Mocks returns the mocks for the current user.
func (s *Service) Mocks(ctx context.Context) (*resources.MockListItems, error) {
	var resource resources.MockListResponse
	if _, err := s.get(ctx, &resource, nil, "mocks"); err != nil {
		return nil, err
	}

	return &resource.Mocks, nil
}

// Mock returns a single mock for the current user.
func (s *Service) Mock(ctx context.Context, id string) (*resources.Mock, error) {
	var resource resources.MockResponse
	if _, err := s.get(ctx, &resource, nil, "mocks", id); err != nil {
		return nil, err
	}

	return &resource.Mock, nil
}
