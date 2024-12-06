package service

import (
	"github.com/jambo0624/blog/internal/shared/application/service"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type TagService struct {
	*service.BaseService[entity.Tag, *query.TagQuery]
}

func NewTagService(repo repository.BaseRepository[entity.Tag, *query.TagQuery]) *TagService {
	baseService := service.NewBaseService(repo)
	return &TagService{
		BaseService: baseService,
	}
}

func (s *TagService) Create(req *dto.CreateTagRequest) (*entity.Tag, error) {
	tag, err := entity.NewTag(req.Name, req.Color)
	if err != nil {
		return nil, err
	}

	if err := s.Repo.Save(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *TagService) Update(id uint, req *dto.UpdateTagRequest) (*entity.Tag, error) {
	tag, err := s.FindByID(id)
	if err != nil {
		return nil, err
	}

	tag.Update(req)

	if err := s.Repo.Update(tag); err != nil {
		return nil, err
	}

	return tag, nil
}
