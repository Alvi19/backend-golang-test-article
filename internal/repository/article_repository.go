package repository

import (
	"context"
	"errors"

	"github.com/Alvi19/backend-golang-test/internal/config"
	"github.com/Alvi19/backend-golang-test/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(ctx context.Context, a *domain.Article) error
	GetByID(ctx context.Context, id uint) (*domain.Article, error)
	List(ctx context.Context, limit, offset int) ([]domain.Article, error)
	Update(ctx context.Context, a *domain.Article) error
	Delete(ctx context.Context, id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func NewPostgresGorm(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.PostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Article{})
}

func (r *articleRepository) Create(ctx context.Context, a *domain.Article) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *articleRepository) GetByID(ctx context.Context, id uint) (*domain.Article, error) {
	var a domain.Article
	if err := r.db.WithContext(ctx).First(&a, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *articleRepository) List(ctx context.Context, limit, offset int) ([]domain.Article, error) {
	var articles []domain.Article
	if err := r.db.WithContext(ctx).
		Order("created_date desc").
		Limit(limit).
		Offset(offset).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) Update(ctx context.Context, a *domain.Article) error {
	return r.db.WithContext(ctx).Save(a).Error
}

func (r *articleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Article{}, id).Error
}
