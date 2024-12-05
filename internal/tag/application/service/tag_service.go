package service

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
)

type TagService struct {
	tagRepo tagRepository.TagRepository
}

func NewTagService(tr tagRepository.TagRepository) *TagService {
	return &TagService{
		tagRepo: tr,
	}
}

func (s *TagService) CreateTag(name string, color string) (*tagEntity.Tag, error) {
	tag := &tagEntity.Tag{
		Name:  name,
		Color: color,
	}
	
	if err := s.tagRepo.Save(tag); err != nil {
		return nil, err
	}
	
	return tag, nil
}

func (s *TagService) GetTagByID(id uint) (*tagEntity.Tag, error) {
	return s.tagRepo.FindByID(id)
}

func (s *TagService) UpdateTag(id uint, name string, color string) (*tagEntity.Tag, error) {
	tag, err := s.tagRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	
	tag.Name = name
	tag.Color = color
	
	if err := s.tagRepo.Update(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *TagService) DeleteTag(id uint) error {
	return s.tagRepo.Delete(id)
}
