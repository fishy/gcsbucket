// Package gcsbucket provides an implementation of Bucket for Google Cloud
// Storage.
package gcsbucket

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/fishy/fsdb/bucket"
)

// Make sure *GCSBucket satisifies bucket.Bucket interface.
var _ bucket.Bucket = (*GCSBucket)(nil)

// GCSBucket is an implementation of bucket with Google Cloud Storage.
type GCSBucket struct {
	client *storage.Client
	bkt    *storage.BucketHandle
}

// Open opens a gcs bucket.
//
// The bucket must already exist, otherwise all operations will fail.
func Open(ctx context.Context, bucket string) (*GCSBucket, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bkt := client.Bucket(bucket)
	return &GCSBucket{
		client: client,
		bkt:    bkt,
	}, nil
}

// Close releases resources associated with this bucket.
//
// All operations will panic after Close is called.
//
// Close need not be called at program exit.
func (gcs *GCSBucket) Close() error {
	return gcs.client.Close()
}

func (gcs *GCSBucket) Read(ctx context.Context, name string) (
	io.ReadCloser,
	error,
) {
	return gcs.bkt.Object(name).NewReader(ctx)
}

func (gcs *GCSBucket) Write(
	ctx context.Context,
	name string,
	data io.Reader,
) error {
	writer := gcs.bkt.Object(name).NewWriter(ctx)
	if _, err := io.Copy(writer, data); err != nil {
		writer.Close()
		return err
	}
	return writer.Close()
}

// Delete deletes an object from the bucket.
func (gcs *GCSBucket) Delete(ctx context.Context, name string) error {
	return gcs.bkt.Object(name).Delete(ctx)
}

// IsNotExist checks whether err is storage.ErrObjectNotExist.
func (gcs *GCSBucket) IsNotExist(err error) bool {
	return err == storage.ErrObjectNotExist
}
