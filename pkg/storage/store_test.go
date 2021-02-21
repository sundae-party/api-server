package storage

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

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

//
// Integration tests
//

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

	// Creating new integration
	ni, err := m_store.PutIntegration(ctx, iOk)
	if err != nil {
		t.Logf("%s", err)
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
	}

	// Test update
	ni.State.Connected = false
	ui, err := m_store.PutIntegration(ctx, ni)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", ni, ui)
		t.Error(err)
	}

	// Delete
	_, err = m_store.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error Deleting the integration -> \n%s\n", err)
	}
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

	// Creating new integration
	ni, err := m_store.PutIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error creating new integration -> \nWANT \n%s \ngot \n%s\n", iOk, ni)
		t.Error(err)
	}
	// Getting the integration
	gi, err := m_store.GetIntegration(ctx, iOk.Name)
	if err != nil {
		t.Fatalf("Error getting new integration -> \nWANT \n%s \ngot \n%s\n", iOk, gi)
	}
	if gi.Name != iOk.Name {
		t.Fatalf("Error getting new integration -> \nWANT \n%s \ngot \n%s\n", iOk, gi)
	}
	// Deleting the integration
	_, err = m_store.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Fatalf("Error Deleting the integration -> \n%s\n", err)
	}
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

func TestUpdateIntegrationDesiredState(t *testing.T) {
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
		t.Errorf("TestUpdateIntegrationDesiredState Error : creating new integration -> %s\n", err)
	}

	// Update the integration desired state
	ni.DesiredState = &types.IntegrationState{Connected: true, ServiceName: ni.Services[0].Name, ServiceData: ni.Services[0].Data, ServiceIdle: false}
	_, err = m_store.UpdateIntegrationDesiredState(ctx, ni)
	if err != nil {
		t.Errorf("TestUpdateIntegrationDesiredState Error : updating integration desired state -> %s\n", err)
	}

	// Get the integration
	gi, err := m_store.GetIntegration(ctx, iOk.Name)
	if err != nil {
		t.Errorf("TestUpdateIntegrationDesiredState Error : getting the updated integration -> %s\n", err)
	}

	// Check if updated desired state is applyed
	if !reflect.DeepEqual(ni.DesiredState, gi.DesiredState) {
		t.Errorf("TestUpdateIntegrationDesiredState Error the integration desired state should be -> %v but have -> %v", ni.DesiredState, gi.DesiredState)
	}

	// Clean, remove integration
	_, err = m_store.DeleteIntegration(ctx, iOk)
	if err != nil {
		t.Errorf("Error Deleting the integration -> \n%s\n", err)
	}
}

//
// Light test
//

func TestCreateInvalidLight(t *testing.T) {
	ctx := context.Background()

	invalidMockLight1 := &types.Light{Name: "l1"}
	invalidMockLight2 := &types.Light{IntegrationName: "i1"}

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

	mockLight1 := &types.Light{Name: "l1", IntegrationName: "i1"}
	mockLight2 := &types.Light{Name: "l2", IntegrationName: "i1"}

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
			count++
		}
		if light.Name == mockLight2.Name {
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

//
// binarySensor test
//

func TestGetAllBinarySensor(t *testing.T) {

	ctx := context.Background()

	mockBinarySensor1 := &types.BinarySensor{Name: "bs1", IntegrationName: "i1"}
	mockBinarySensor2 := &types.BinarySensor{Name: "bs2", IntegrationName: "i1"}

	// Create mock binary sensor
	bs1, err := m_store.PutBinarySensor(ctx, mockBinarySensor1)
	if err != nil {
		t.Fatalf("Fail to create mock binary sensor %s", err)
	}
	bs2, err := m_store.PutBinarySensor(ctx, mockBinarySensor2)
	if err != nil {
		t.Fatalf("Fail to create mock binary sensor %s", err)
	}

	// Try to get it
	binarySensors, err := m_store.GetAllBinarySensor(ctx)
	if err != nil {
		t.Fatalf("Error getting binary sensor -> %s\n", err)
	}
	count := 0
	for _, binarySensor := range binarySensors {
		if binarySensor.Name == mockBinarySensor1.Name {
			count++
		}
		if binarySensor.Name == mockBinarySensor2.Name {
			count++
		}
	}

	// Clean created binary sensor
	_, err = m_store.DeleteBinarySensor(ctx, bs1)
	if err != nil {
		t.Fatalf("Error Deleting the binary sensor bs1 -> \n%s\n", err)
	}
	_, err = m_store.DeleteBinarySensor(ctx, bs2)
	if err != nil {
		t.Fatalf("Error Deleting the binary sensor bs2 -> \n%s\n", err)
	}

	if count != 2 {
		t.Log("Error, should have bs1 & bs2 but have -> \n")
		for _, bs := range binarySensors {
			t.Logf("%s\n", bs.Name)
		}
		t.Fatalf("All binary sensors not found\n")
	}
}

//
// Sensor test
//

func TestGetAllSensor(t *testing.T) {

	ctx := context.Background()

	mockSensor1 := &types.Sensor{Name: "s1", IntegrationName: "i1"}
	mockSensor2 := &types.Sensor{Name: "s2", IntegrationName: "i1"}

	// Create mock sensor
	s1, err := m_store.PutSensor(ctx, mockSensor1)
	if err != nil {
		t.Fatalf("Fail to create mock sensor %s", err)
	}
	s2, err := m_store.PutSensor(ctx, mockSensor2)
	if err != nil {
		t.Fatalf("Fail to create mock sensor %s", err)
	}

	// Try to get it
	sensors, err := m_store.GetAllSensor(ctx)
	if err != nil {
		t.Fatalf("Error getting sensors -> %s\n", err)
	}
	count := 0
	for _, sensor := range sensors {
		if sensor.Name == mockSensor1.Name {
			count++
		}
		if sensor.Name == mockSensor2.Name {
			count++
		}
	}

	// Clean created sensor
	_, err = m_store.DeleteSensor(ctx, s1)
	if err != nil {
		t.Fatalf("Error Deleting the sensor s1 -> \n%s\n", err)
	}
	_, err = m_store.DeleteSensor(ctx, s2)
	if err != nil {
		t.Fatalf("Error Deleting the sensor s2 -> \n%s\n", err)
	}

	if count != 2 {
		t.Log("Error, should have s1 & s2 but have -> \n")
		for _, s := range sensors {
			t.Logf("%s\n", s.Name)
		}
		t.Fatalf("All sensors not found\n")
	}
}

//
// Sun test
//

func TestGetSun(t *testing.T) {

	ctx := context.Background()

	strDate := time.Now().Format(time.RFC3339)

	mockSunState := &types.SunState{
		NextRising:   strDate,
		NextSetting:  strDate,
		NextNoon:     strDate,
		NextMidnight: strDate,
		Elevation:    12.3,
		Azimuth:      36.12,
		State:        types.SunState_above_horizon,
	}

	// Create mock sun
	_, err := m_store.PutSun(ctx, mockSunState)
	if err != nil {
		t.Fatalf("Fail to create mock sun %s", err)
	}

	// Try to get it
	sun, err := m_store.GetSun(ctx)
	if err != nil {
		t.Fatalf("Error getting sun -> %s\n", err)
	}

	// Clean created sun
	_, err = m_store.DeleteSun(ctx)
	if err != nil {
		t.Fatalf("Error Deleting the sun -> \n%s\n", err)
	}

	if sun.Name != "sun" || sun.IntegrationName != "sun" || sun.Mutation != "sun" {
		t.Errorf("TestGetSun Error, sun invalid format, name, integration name and mutation sould be sun but have => %s", sun)
	}
	if !reflect.DeepEqual(mockSunState, sun.State) {
		t.Errorf("TestGetSun Error, sun invalid sun state format, want => %s have => %s", mockSunState, sun.State)
	}
}
