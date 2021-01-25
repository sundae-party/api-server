package storage

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
)

var (
	m_store Store
)

func init() {
	ctx := context.Background()
	mongoOpsOK := &StoreOption{
		Type:     "mongo",
		Address:  []string{os.Getenv("MONGO_ADDR")},
		User:     os.Getenv("MONGO_USR"),
		Password: os.Getenv("MONGO_PWD"),
		DbName:   os.Getenv("MONGO_DB"),
		RsName:   os.Getenv("MONGO_RS"),
	}

	store, err := NewStore(ctx, mongoOpsOK)
	if err != nil {
		log.Fatalf("Init err : %s", err)
	}
	m_store = store
}

func TestPutIntegration(t *testing.T) {
	ctx := context.Background()

	iOk := &types.Integration{
		Name:          "Hue",
		Documentation: "https://sundae/doc/hue",
		Version:       "v1.0.0",
		Url:           "https://github.com/sundae-party/integration/hue",
		State: &types.IntegrationState{
			Connected: true,
		},
		Services: []*types.IntegrationService{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}

	t.Logf("Creating new integration\n")
	ni, err := m_store.PutIntegration(ctx, iOk)
	if err != nil {
		t.Logf("%s", err)
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
	}

	// Test update
	t.Logf("Updating integration\n")
	ni.State.Connected = false
	ui, err := m_store.PutIntegration(ctx, ni)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", ni, ui)
		t.Error(err)
	}

	t.Logf("Deleting the integration\n")
	res, err := m_store.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error Deleting the integration -> \n%s\n", err)
	}
	t.Log(res)
}

func TestGetIntegration(t *testing.T) {
	ctx := context.Background()

	iOk := &types.Integration{
		Name:          "Hue",
		Documentation: "https://sundae/doc/hue",
		Version:       "v1.0.0",
		Url:           "https://github.com/sundae-party/integration/hue",
		State: &types.IntegrationState{
			Connected: true,
		},
		Services: []*types.IntegrationService{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}

	t.Logf("Creating new integration\n")
	ni, err := m_store.PutIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
		t.Error(err)
	}
	t.Logf("Getting the integration\n")
	gi, err := m_store.GetIntegration(ctx, iOk.Name)
	if err != nil {
		t.Fatalf("Error getting new integration -> \nWANT \n%s \ngot \n%s\n", iOk, gi)
	}
	if gi.Name != iOk.Name {
		t.Fatalf("Error getting new integration -> \nWANT \n%s \ngot \n%s\n", iOk, gi)
	}
	t.Logf("Deleting the integration\n")
	res, err := m_store.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error Deleting the integration -> \n%s\n", err)
	}
	t.Log(res)
}

func TestUpdateIntegrationState(t *testing.T) {
	ctx := context.Background()

	iOk := &types.Integration{
		Name:          "i1",
		Documentation: "https://sundae/doc/hue",
		Version:       "v1.0.0",
		Url:           "https://github.com/sundae-party/integration/i1",
		State: &types.IntegrationState{
			Connected: true,
		},
		Services: []*types.IntegrationService{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}

	// Create ne integration
	ni, err := m_store.PutIntegration(ctx, iOk)
	if err != nil {
		t.Errorf("TestUpdateIntegrationState Error : creating new integration -> %s\n", err)
	}

	// Update the integration state
	ni.State = &types.IntegrationState{Connected: true, ServiceName: ni.Services[0].Name, ServiceData: ni.Services[0].Data, ServiceIdle: false}
	_, err = m_store.UpdateIntegrationState(ctx, ni)
	if err != nil {
		t.Errorf("TestUpdateIntegrationState Error : updating integration state -> %s\n", err)
	}

	// Get the integration
	gi, err := m_store.GetIntegration(ctx, iOk.Name)
	if err != nil {
		t.Errorf("TestUpdateIntegrationState Error : getting the updated integration -> %s\n", err)
	}

	// Check if updated state is applyed
	if !reflect.DeepEqual(ni.State, gi.State) {
		t.Errorf("TestUpdateIntegrationState Error the integration state should be -> %v but have -> %v", ni.State, gi.State)
	}

	// Clean, remove integration
	_, err = m_store.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Errorf("Error Deleting the integration -> \n%s\n", err)
	}
}

// Light test
func TestCreateInvalidLight(t *testing.T) {
	ctx := context.Background()

	invalidMockLight1 := &types.Light{Name: "l1"}
	invalidMockLight2 := &types.Light{Integration: &types.Integration{Name: "i1"}}

	l1, err := m_store.PutLight(ctx, invalidMockLight1)
	if err != nil {
		t.Logf("OK => fail to create invalid mock light : %s", err)
	}
	if l1 != nil {
		t.Fatalf("ERROR this light is invalid => %v\n", l1)
	}

	l2, err := m_store.PutLight(ctx, invalidMockLight2)
	if err != nil {
		t.Logf("OK => fail to create invalid mock light :  %s", err)
	}
	if l2 != nil {
		t.Fatalf("ERROR this light is invalid => %v\n", l2)
	}
}

func TestGetAllLight(t *testing.T) {

	ctx := context.Background()

	mockLight1 := &types.Light{Name: "l1", Integration: &types.Integration{Name: "i1"}}
	mockLight2 := &types.Light{Name: "l2", Integration: &types.Integration{Name: "i1"}}

	// Create mock light
	l1, err := m_store.PutLight(ctx, mockLight1)
	if err != nil {
		t.Fatalf("Fail to create mock light %s", err)
	}
	l2, err := m_store.PutLight(ctx, mockLight2)
	if err != nil {
		t.Fatalf("Fail to create mock light %s", err)
	}

	// Try to get it
	t.Logf("Getting all lights\n")
	lights, err := m_store.GetAllLight(ctx)
	if err != nil {
		t.Fatalf("Error getting lights -> %s\n", err)
	}
	count := 0
	for _, light := range lights {
		if light.Name == mockLight1.Name {
			t.Log(light.Name)
			count++
		}
		if light.Name == mockLight2.Name {
			t.Log(light.Name)
			count++
		}
	}

	// Clean created light
	_, err = m_store.DeleteLight(ctx, l1)
	if err != nil {
		t.Fatalf("Error Deleting the light l1 -> \n%s\n", err)
	}
	_, err = m_store.DeleteLight(ctx, l2)
	if err != nil {
		t.Fatalf("Error Deleting the light l2 -> \n%s\n", err)
	}

	if count != 2 {
		t.Log("Error, should have l1 & l2 but have -> \n")
		for _, light := range lights {
			t.Logf("%s\n", light.Name)
		}
		t.Fatalf("All light not found\n")
	}
}
