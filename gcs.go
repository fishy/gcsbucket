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
	name string
}

// Open opens a gcs bucket.
//
// The bucket must already exist, otherwise all operations will fail.
//
// It doesn't do any I/O operations.
func Open(bucket string) *GCSBucket {
	return &GCSBucket{
		name: bucket,
	}
}

func (gcs *GCSBucket) Read(ctx context.Context, name string) (
	io.ReadCloser,
	error,
) {
	obj, err := gcs.getObject(ctx, name)
	if err != nil {
		return nil, err
	}
	return obj.NewReader(ctx)
}

func (gcs *GCSBucket) Write(
	ctx context.Context,
	name string,
	data io.Reader,
) error {
	obj, err := gcs.getObject(ctx, name)
	if err != nil {
		return err
	}
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, data); err != nil {
		writer.Close()
		return err
	}
	return writer.Close()
}

// Delete deletes an object from the bucket.
func (gcs *GCSBucket) Delete(ctx context.Context, name string) error {
	obj, err := gcs.getObject(ctx, name)
	if err != nil {
		return err
	}
	return obj.Delete(ctx)
}

// IsNotExist checks whether err is storage.ErrObjectNotExist.
func (gcs *GCSBucket) IsNotExist(err error) bool {
	return err == storage.ErrObjectNotExist
}

// getObject returns an ObjectHandle of a given object name.
func (gcs *GCSBucket) getObject(
	ctx context.Context,
	name string,
) (*storage.ObjectHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bkt := client.Bucket(gcs.name)
	return bkt.Object(name), nil
}
