//go:build integration

package storage

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/anhsbolic/kencleng/backend/internal/domain/campaign"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestPrivateReader_ReadsPrivateMediaAndClassifiesUnavailable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	reader := isolatedPrivateReader(t, ctx)
	if err := reader.Put(ctx, "campaign/test.jpg", strings.NewReader("jpeg-bytes"), int64(len("jpeg-bytes")), "image/jpeg"); err != nil {
		t.Fatalf("put private fixture: %v", err)
	}
	content, err := reader.Read(ctx, "campaign/test.jpg")
	if err != nil {
		t.Fatalf("read private fixture: %v", err)
	}
	body, err := io.ReadAll(content.Reader)
	if err != nil {
		t.Fatalf("read private fixture bytes: %v", err)
	}
	if err := content.Reader.Close(); err != nil {
		t.Fatalf("close private fixture reader: %v", err)
	}
	if string(body) != "jpeg-bytes" || content.ContentType != "image/jpeg" {
		t.Fatalf("content = %q / %q", body, content.ContentType)
	}
	for _, testCase := range []struct {
		name string
		ctx  context.Context
		key  string
	}{{name: "missing object stat failure", ctx: ctx, key: "campaign/missing.jpg"}, {name: "cancelled request context", ctx: cancelledContext(), key: "campaign/test.jpg"}} {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := reader.Read(testCase.ctx, testCase.key)
			if !errors.Is(err, campaign.ErrObjectUnavailable) {
				t.Fatalf("error = %v, want safe object-unavailable classification", err)
			}
		})
	}
}

func cancelledContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func isolatedPrivateReader(t *testing.T, ctx context.Context) *PrivateReader {
	t.Helper()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "quay.io/minio/minio:RELEASE.2025-04-22T22-12-26Z", ExposedPorts: []string{"9000/tcp"},
			Env:        map[string]string{"MINIO_ROOT_USER": "kenclengtest", "MINIO_ROOT_PASSWORD": "kenclengtestsecret"},
			Cmd:        []string{"server", "/data"},
			WaitingFor: wait.ForHTTP("/minio/health/live").WithPort("9000/tcp").WithStartupTimeout(time.Minute),
		},
		Started: true,
	})
	if err != nil {
		t.Fatalf("start isolated MinIO through the configured Docker-compatible runtime: %v", err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(context.Background()); err != nil {
			t.Errorf("terminate isolated MinIO: %v", err)
		}
	})
	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("resolve isolated MinIO host: %v", err)
	}
	port, err := container.MappedPort(ctx, "9000/tcp")
	if err != nil {
		t.Fatalf("resolve isolated MinIO port: %v", err)
	}
	client, err := minio.New(host+":"+port.Port(), &minio.Options{Creds: credentials.NewStaticV4("kenclengtest", "kenclengtestsecret", ""), Secure: false})
	if err != nil {
		t.Fatalf("create isolated MinIO client: %v", err)
	}
	const bucket = "campaign-private"
	if err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
		t.Fatalf("create isolated private bucket: %v", err)
	}
	return NewPrivateReader(client, bucket)
}
