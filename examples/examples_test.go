package examples

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/testcontainers/testcontainers-go/modules/compose"
)

func TestMain(m *testing.M) {
	os.Exit(runAcceptanceTests(m))
}

func runAcceptanceTests(m *testing.M) int {
	dc, err := compose.NewDockerCompose(filepath.Join("..", "docker-compose.yaml"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = dc.WithOsEnv().Up(ctx, compose.Wait(true)); err != nil {
		panic(err)
	}

	defer func() {
		if err := dc.Down(context.Background(), compose.RemoveOrphans(true), compose.RemoveImagesLocal); err != nil {
			panic(err)
		}
	}()

	container, err := dc.ServiceContainer(ctx, "unifi")
	if err != nil {
		panic(err)
	}

	endpoint, err := container.PortEndpoint(ctx, "8443/tcp", "https")
	if err != nil {
		panic(err)
	}

	if err := setupEnvironment(endpoint); err != nil {
		panic(err)
	}

	return m.Run()
}

func setupEnvironment(endpoint string) error {
	settings := map[string]string{
		"UNIFI_USERNAME": "admin",
		"UNIFI_PASSWORD": "admin",
		"UNIFI_INSECURE": "true",
		"UNIFI_API":      endpoint,
	}

	for name, value := range settings {
		if err := os.Setenv(name, value); err != nil {
			return err
		}
	}

	return nil
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest:        false,
		ExpectRefreshChanges: true,
	}
}
