package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	recommendationsTable = "recommendations"
)

func (r *repository) AddRecommendation(ctx context.Context, recommendation dto.Recommendation) error {
	insertQuery, _, err := goqu.Insert(recommendationsTable).Rows(recommendation).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetRecommendations(ctx context.Context) ([]entity.Recommendation, error) {
	selectQuery, _, err := goqu.From(recommendationsTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var recommendationsDto []dto.Recommendation
	if err = r.db.DB.SelectContext(ctx, &recommendationsDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	recommendations := make([]entity.Recommendation, len(recommendationsDto))
	for i, recommendationDto := range recommendationsDto {
		recommendation := entity.Recommendation{
			Id:              recommendationDto.Id,
			Peer:            recommendationDto.Peer,
			RecommendedPeer: recommendationDto.RecommendedPeer,
		}
		recommendations[i] = recommendation
	}
	return recommendations, nil
}

func (r *repository) UpdateRecommendation(ctx context.Context, recommendation dto.Recommendation) error {
	updateQuery, _, err := goqu.Update(recommendationsTable).Set(
		goqu.Record{
			"peer":             recommendation.Peer,
			"recommended_peer": recommendation.RecommendedPeer,
		}).Where(goqu.C("id").Eq(recommendation.Id)).
		Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	id := ""
	row := r.db.DB.QueryRowxContext(ctx, updateQuery)
	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("update data: %w", err)
	}
	return nil
}

func (r *repository) DeleteRecommendation(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(recommendationsTable).
		Where(goqu.C("id").Eq(key)).
		Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	id := ""
	row := r.db.DB.QueryRowxContext(ctx, deleteQuery)
	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("delete data: %w", err)
	}
	return nil
}
