package light

import (
	context "context"
	"io"
	"log"
	"net"
	"testing"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	// Create new mongo store
	mongo := &storage.StoreOption{
		Type:     "mongo",
		Address:  []string{"gogs.connan.pro:27018"},
		User:     "sundae",
		Password: "pass",
		DbName:   "sundae",
		RsName:   "rs0",
	}
	ctx := context.Background()
	mongoStore, err := storage.NewStore(ctx, mongo)
	if err != nil {
		log.Fatalln(err)
	}
	// Create mock grpc server
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	lh := &LightHandler{
		Store: mongoStore,
	}
	// Register light handler
	types.RegisterLightHandlerServer(s, lh)
	// Server mock grpc server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetAll(t *testing.T) {

	// Prepare GRPC client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("TestGetAll failed, gRPC, failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	lh := types.NewLightHandlerClient(conn)

	// Insert mock light
	mockLight1 := &types.Light{Name: "l1", Integration: &types.Integration{Name: "i1"}}
	mockLight2 := &types.Light{Name: "l2", Integration: &types.Integration{Name: "i1"}}

	l1, err := lh.Create(ctx, mockLight1)
	if err != nil {
		t.Fatalf("TestGetAll failed, create mockLight1 failed: %s", err)
	}
	l2, err := lh.Create(ctx, mockLight2)
	if err != nil {
		t.Fatalf("TestGetAll failed, create mockLight2 failed: %s", err)
	}

	// Try to get all light
	resp, err := lh.GetAll(ctx, &types.GetAllRequest{})
	if err != nil {
		t.Fatalf("TestGetAll failed: %s", err)
	}

	count := 0
	for {
		light, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				// End of stream
				// This Recv method returns (nil, io.EOF) once the server-to-client stream has been completely read through.
				log.Println("End of stream")
				break
			}
			t.Fatalf("TestGetAll failed reading stream resp: %s", err)
		}
		if light.Name == mockLight1.Name {
			count++
			log.Printf("%s", light)
		}
		if light.Name == mockLight2.Name {
			count++
			log.Printf("%s", light)
		}
	}

	// Clean l1 & l2
	_, err = lh.Delete(ctx, l1)
	if err != nil {
		t.Fatalf("TestGetAll failed, delete l1 failed: %s", err)
	}
	_, err = lh.Delete(ctx, l2)
	if err != nil {
		t.Fatalf("TestGetAll failed, delete l2 failed: %s", err)
	}

	if count != 2 {
		t.Fatalf("TestGetAll failed, should have l1 & l2 (count = 2) but have -> count = %d \n", count)
	}
}

// var (
// 	mockIntegration = &types.Integration{
// 		Name:          "Hue",
// 		Documentation: "https://sundae/doc/hue",
// 		Version:       "v1.0.0",
// 		Url:           "https://github.com/sundae-party/integration/hue",
// 		State: &types.IntegrationState{
// 			Connected: true,
// 		},
// 		Services: []*types.IntegrationService{
// 			{
// 				Name: "refresh_entities",
// 				Data: "",
// 			},
// 		},
// 	}

// 	mockState = &types.LightState{
// 		Brightness: 33,
// 		ColorRGB: &types.LightColorRGB{
// 			Red:   125,
// 			Blue:  135,
// 			Green: 120,
// 		},
// 		On:     true,
// 		Kelvin: 2500,
// 	}

// 	mockLight = &types.Light{
// 		Name:          "light",
// 		DisplayedName: "Desk light",
// 		Integration:   mockIntegration,
// 		State:         mockState,
// 		Room:          "Desk",
// 		Device:        "Light strip",
// 	}
// )

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
