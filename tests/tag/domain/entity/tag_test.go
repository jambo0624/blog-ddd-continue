package entity_test

import (
	"testing"

	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	tests := []struct {
		name        string
		tagName     string
		color       string
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "valid tag",
			tagName: "Test Tag",
			color:   "#FF0000",
			wantErr: false,
		},
		{
			name:        "empty name",
			tagName:     "",
			color:       "#FF0000",
			wantErr:     true,
			expectedErr: validate.ErrNameRequired,
		},
		{
			name:        "empty color",
			tagName:     "Test Tag",
			color:       "",
			wantErr:     true,
			expectedErr: validate.ErrColorRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := entity.NewTag(tt.tagName, tt.color)
			if tt.wantErr {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.tagName, tag.Name)
			assert.Equal(t, tt.color, tag.Color)
		})
	}
}

func TestTag_Update(t *testing.T) {
	tag, _ := entity.NewTag("Original", "#000000")
	req := &dto.UpdateTagRequest{
		Name:  "Updated",
		Color: "#FFFFFF",
	}

	tag.Update(req)

	assert.Equal(t, "Updated", tag.Name)
	assert.Equal(t, "#FFFFFF", tag.Color)
}
