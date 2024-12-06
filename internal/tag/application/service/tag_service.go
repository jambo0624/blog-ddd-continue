package service

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/domain/query"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type TagService struct {
	tagRepo tagRepository.TagRepository
}

func NewTagService(tr tagRepository.TagRepository) *TagService {
	return &TagService{
		tagRepo: tr,
	}
}

func (s *TagService) Create(req *dto.CreateTagRequest) (*tagEntity.Tag, error) {
	tag, err := tagEntity.NewTag(req.Name, req.Color)
	if err != nil {
		return nil, err
	}

	if err := s.tagRepo.Save(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *TagService) FindByID(id uint) (*tagEntity.Tag, error) {
	return s.tagRepo.FindByID(id)
}

func (s *TagService) FindAll(q *query.TagQuery) ([]*tagEntity.Tag, error) {
	if q == nil {
		q = query.NewTagQuery()
	}
	return s.tagRepo.FindAll(q)
}

func (s *TagService) Update(id uint, req *dto.UpdateTagRequest) (*tagEntity.Tag, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	tag.Update(req)

	if err := s.tagRepo.Update(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *TagService) Delete(id uint) error {
	return s.tagRepo.Delete(id)
}
