package data_manager

import (
	"context"
	"io"
)

type Service interface {
	Import(ctx context.Context, table string, body io.Reader) error
	Export(ctx context.Context, table string) ([]byte, error)
	RawQuery(ctx context.Context, query string) ([]map[string]interface{}, error)
}
