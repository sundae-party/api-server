package light

import (
	context "context"
	"io"
	"log"
	"net"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/server/utils"
	"github.com/sundae-party/api-server/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	rootDir()
	// Create new mongo store
	mongo := &storage.StoreOption{
		Type:     "mongo",
		Address:  []string{os.Getenv("MONGO_ADDR")},
		User:     os.Getenv("MONGO_USR"),
		Password: os.Getenv("MONGO_PWD"),
		DbName:   os.Getenv("MONGO_DB"),
		RsName:   os.Getenv("MONGO_RS"),
	}
	ctx := context.Background()
	mongoStore, err := storage.NewStore(ctx, mongo)
	if err != nil {
		log.Fatalln(err)
	}
	// Create mock grpc server
	lis = bufconn.Listen(bufSize)
	tlsConfig, err := utils.BuildServerTlsConf([]string{"ssl/ca.pem"}, "ssl/sundae-apiserver.pem", "ssl/sundae-apiserver.key")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))
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

func rootDir() {
	_, b, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(b), "../../../../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestGetAll(t *testing.T) {

	// Prepare GRPC client
	ctx := context.Background()
	// Build tls client conf

	cliTlsConf, err := utils.LoadKeyPair("ssl/integration01.pem", "ssl/integration01.key", "ssl/ca.pem")
	if err != nil {
		t.Fatalf("TestGetAll failed, gRPC, failed to create client tls config: %v", err)
	}
	// Create grpc connexion with tls conf
	conn, err := grpc.DialContext(ctx, "sundae.com", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(cliTlsConf))
	if err != nil {
		t.Fatalf("TestGetAll failed, gRPC, failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	// Create new client handler client
	lh := types.NewLightHandlerClient(conn)

	// Insert mock light
	mockLight1 := &types.Light{Name: "l1", IntegrationName: "i1"}
	mockLight2 := &types.Light{Name: "l2", IntegrationName: "i1"}

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
				break
			}
			t.Fatalf("TestGetAll failed reading stream resp: %s", err)
		}
		if light.Name == mockLight1.Name {
			count++
		}
		if light.Name == mockLight2.Name {
			count++
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
