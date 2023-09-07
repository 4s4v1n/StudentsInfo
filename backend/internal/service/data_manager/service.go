package data_manager

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/sav1nbrave4code/APG3/internal/repository"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *service {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		w := csv.NewWriter(out)
		w.Comma = ';'
		return gocsv.NewSafeCSVWriter(w)
	})

	return &service{
		repo: repo,
	}
}

func (s *service) Import(ctx context.Context, table string, body io.Reader) error {
	if err := s.repo.Import(ctx, table, body); err != nil {
		return fmt.Errorf("import data: %w", err)
	}
	return nil
}

func (s *service) Export(ctx context.Context, table string) ([]byte, error) {
	out, err := s.repo.Export(ctx, table)
	if err != nil {
		return nil, fmt.Errorf("export data: %w", err)
	}

	return out, nil
}

func (s *service) RawQuery(ctx context.Context, query string) ([]map[string]interface{}, error) {
	res, err := s.repo.RawQuery(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("exec query %w", err)
	}

	return res, nil
}
