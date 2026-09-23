// Package storage provides concrete S3-compatible private-object access.
package storage

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"

	"github.com/anhsbolic/kencleng/backend/internal/domain/campaign"
)

// PrivateReader reads Campaign media exclusively from one configured private
// bucket. It deliberately exposes no public, signed, or redirect URL method.
type PrivateReader struct {
	client *minio.Client
	bucket string
}

// NewPrivateReader creates the MinIO adapter for the already-configured
// private bucket.
func NewPrivateReader(client *minio.Client, bucket string) *PrivateReader {
	return &PrivateReader{client: client, bucket: bucket}
}

// Read establishes object availability and content type before returning a
// stream, allowing the HTTP layer to send a 503 before response headers commit.
func (r *PrivateReader) Read(ctx context.Context, objectKey string) (*campaign.MediaContent, error) {
	object, err := r.client.GetObject(ctx, r.bucket, objectKey, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("storage: get private object: %w", errors.Join(campaign.ErrObjectUnavailable, err))
	}
	info, err := object.Stat()
	if err != nil {
		_ = object.Close()
		return nil, fmt.Errorf("storage: stat private object: %w", errors.Join(campaign.ErrObjectUnavailable, err))
	}
	if info.ContentType != "image/jpeg" && info.ContentType != "image/png" {
		_ = object.Close()
		return nil, fmt.Errorf("storage: unsupported private object type: %w", campaign.ErrObjectUnavailable)
	}
	return &campaign.MediaContent{Reader: objectReadCloser{object}, ContentType: info.ContentType}, nil
}

// Put stores an operator-validated object in the same private bucket. It is
// intentionally only used by the explicit seed command, never by HTTP.
func (r *PrivateReader) Put(ctx context.Context, objectKey string, body io.Reader, size int64, contentType string) error {
	if contentType != "image/jpeg" && contentType != "image/png" {
		return fmt.Errorf("storage: unsupported private object type")
	}
	_, err := r.client.PutObject(ctx, r.bucket, objectKey, body, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fmt.Errorf("storage: put private object: %w", err)
	}
	return nil
}

// Remove deletes a private object during seed replacement cleanup.
func (r *PrivateReader) Remove(ctx context.Context, objectKey string) error {
	if err := r.client.RemoveObject(ctx, r.bucket, objectKey, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("storage: remove private object: %w", err)
	}
	return nil
}

// objectReadCloser keeps MinIO's close operation explicit at the domain seam.
type objectReadCloser struct{ *minio.Object }

func (r objectReadCloser) Close() error { return r.Object.Close() }

var _ io.ReadCloser = objectReadCloser{}
var _ campaign.ObjectReader = (*PrivateReader)(nil)
