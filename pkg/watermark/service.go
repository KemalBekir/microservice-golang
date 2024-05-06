package watermark

import (
	"context"

	"Go-Tutorials/microservice-golang/internal"
)

type Service interface {
	// Get the list of all documents
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, tickedID string) (internal.Status, error)
	Watermark(ctx context.Context, tickedID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
