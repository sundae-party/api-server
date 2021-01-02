package storage

import (
	"context"
	"sundae-party/api-server/pkg/apis/core/types"
	"testing"
)

func TestNewStore(t *testing.T) {
	ctx := context.Background()

	mongoOpsOK := &StoreOption{
		Type:     "mongo",
		Address:  []string{"mongo:27017"},
		User:     "sundae",
		Password: "pass",
		DbName:   "sundae",
		RsName:   "rs0",
	}

	_, err := NewStore(ctx, mongoOpsOK)
	if err != nil {
		t.Error(err)
	}
}

func TestPutIntegration(t *testing.T) {
	ctx := context.Background()

	mongoOpsOK := &StoreOption{
		Type:     "mongo",
		Address:  []string{"mongo:27017"},
		User:     "sundae",
		Password: "pass",
		DbName:   "sundae",
		RsName:   "rs0",
	}

	iOk := &types.Integration{
		Metadata: &types.Integration_IntegrationMeta{
			Name:          "Hue",
			Documentation: "https://sundae/doc/hue",
		},
		State: &types.Integration_IntegrationState{
			Connected: true,
			Version:   "v1.0.0",
			Url:       "https://github.com/sundae-party/integration/hue",
		},
		Services: []*types.IntegrationService{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}
	t.Logf("Creating new store\n")
	s, err := NewStore(ctx, mongoOpsOK)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Creating new integration\n")
	ni, err := s.PutIntegration(ctx, iOk)
	if err != nil {
		t.Logf("%s", err)
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
	}

	// Test update
	t.Logf("Updating integration\n")
	ni.State.Connected = false
	ui, err := s.PutIntegration(ctx, ni)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", ni, ui)
		t.Error(err)
	}
}

func TestGetIntegration(t *testing.T) {
	ctx := context.Background()

	mongoOpsOK := &StoreOption{
		Type:     "mongo",
		Address:  []string{"mongo:27017"},
		User:     "sundae",
		Password: "pass",
		DbName:   "sundae",
		RsName:   "rs0",
	}

	iOk := &types.Integration{
		Metadata: &types.Integration_IntegrationMeta{
			Name:          "Hue",
			Documentation: "https://sundae/doc/hue",
		},
		State: &types.Integration_IntegrationState{
			Connected: true,
			Version:   "v1.0.0",
			Url:       "https://github.com/sundae-party/integration/hue",
		},
		Services: []*types.IntegrationService{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}
	t.Logf("Creating new store\n")
	s, err := NewStore(ctx, mongoOpsOK)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Creating new integration\n")
	ni, err := s.PutIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
		t.Error(err)
	}
	t.Logf("Getting the integration\n")
	gi, err := s.GetIntegration(ctx, iOk.Metadata.Name)
	if err != nil {
		t.Fatalf("Error getting new integration -> \nWANT \n%s \ngot \n%s\n", iOk, gi)
	}
	if gi.Metadata.Name != iOk.Metadata.Name {
		t.Fatalf("Error getting new integration -> \nWANT \n%s \ngot \n%s\n", iOk, gi)
	}
}

func TestDeleteIntegration(t *testing.T) {
	ctx := context.Background()

	mongoOpsOK := &StoreOption{
		Type:     "mongo",
		Address:  []string{"mongo:27017"},
		User:     "sundae",
		Password: "pass",
		DbName:   "sundae",
		RsName:   "rs0",
	}

	iOk := &types.Integration{
		Metadata: &types.Integration_IntegrationMeta{
			Name:          "Hue",
			Documentation: "https://sundae/doc/hue",
		},
		State: &types.Integration_IntegrationState{
			Connected: true,
			Version:   "v1.0.0",
			Url:       "https://github.com/sundae-party/integration/hue",
		},
		Services: []*types.IntegrationService{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}
	t.Logf("Creating new store\n")
	s, err := NewStore(ctx, mongoOpsOK)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Creating new integration\n")
	ni, err := s.PutIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
		t.Error(err)
	}
	t.Logf("Deleting the integration\n")
	res, err := s.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error Deleting the integration -> \n%s\n", err)
	}
	t.Log(res)
}
