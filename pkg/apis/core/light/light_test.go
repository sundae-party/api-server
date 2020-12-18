package light

import (
	"sundae-party/api-server/pkg/apis/core/integration"
)

var (
	mockIntegration = &integration.Integration{
		Name:          "Hue",
		Documentation: "https://sundae/doc/hue",
		Version:       "v1.0.0",
		Url:           "https://github.com/sundae-party/integration/hue",
		State: &integration.State{
			Connected: true,
		},
		StorePath: "/integration/store",
		Services: []*integration.Service{
			{
				Name: "refresh_entities",
				Data: "",
			},
		},
	}

	mockState = &State{
		Brightness: 33,
		ColorRGB: &ColorRGB{
			Red:   125,
			Blue:  135,
			Green: 120,
		},
		On:     true,
		Kelvin: 2500,
	}

	mockLight = &Light{
		Name:          "light",
		DisplayedName: "Desk light",
		Integration:   mockIntegration,
		State:         mockState,
		Room:          "Desk",
		Device:        "Light strip",
	}
)

// func TestCreate(t *testing.T) {
// 	s := etcd3.NewStore()
// 	lh := &LightHandler{Store: s}

// 	got0, err := lh.Create(context.TODO(), mockLight)
// 	if err != nil {
// 		println(err)
// 		t.Fatal(err)
// 	}

// 	mockLight.Name = "Light 1"
// 	_, err = lh.Create(context.TODO(), mockLight)
// 	if err != nil {
// 		println(err)
// 		t.Fatal(err)
// 	}
// 	t.Log(lh.Store.GetByIntegration("/Hue"))

// 	mockLight.Name = "Light 2"
// 	_, err = lh.Create(context.TODO(), mockLight)
// 	if err != nil {
// 		println(err)
// 		t.Fatal(err)
// 	}

// 	mockLight.Name = "Light 0"
// 	mockLight.Integration.Name = "Zwave"
// 	_, err = lh.Create(context.TODO(), mockLight)
// 	if err != nil {
// 		println(err)
// 		t.Fatal(err)
// 	}

// 	data := lh.Store.GetByIntegration("Zwave")
// 	for i, e := range data {
// 		t.Logf("========> %d  ----  %s\n", i, e)
// 	}

// 	if cmp.Equal(got0, &mockLight) == true {
// 		t.Errorf("(context.TODO(), &mockLight) = %+v; want %+v\n", got0, mockLight)
// 	}
// }
