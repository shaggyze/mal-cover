package service

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/rl404/fairy/errors/stack"
	"github.com/shaggyze/mal-cover/internal/utils"
)

// GenerateCoverRequest is request model for generate cover.
type GenerateCoverRequest struct {
	Username string `validate:"required" mod:"no_space"`
	Type     string `validate:"required,oneof=anime manga" mod:"no_space,lcase"`
	Style    string `validate:"style" mod:"trim,unescape"`
	Size     string `mod:"no_space,lcase"`
}

// GenerateCover to generate css cover.
func (s *service) GenerateCover(ctx context.Context, data GenerateCoverRequest) (string, int, error) {
	if err := utils.Validate(&data); err != nil {
		return "", http.StatusBadRequest, stack.Wrap(ctx, err)
	}

	// Get user's anime/manga list.
	list, code, err := s.mal.GetList(ctx, data.Username, data.Type)
	if err != nil {
		return "", code, stack.Wrap(ctx, err)
	}

	// Replace css style.
	cssRow := make([]string, len(list))
	for i, l := range list {
		if data.Size == "small" {
			cssRow[i] = strings.NewReplacer("{id}", strconv.Itoa(l.ID), "{url}", strings.NewReplacer(".jpg", "t.jpg").Replace(l.Image)).Replace(data.Style)
		} else if data.Size == "medium" {
			cssRow[i] = strings.NewReplacer("{id}", strconv.Itoa(l.ID), "{url}", l.Image).Replace(data.Style)
		} else if data.Size == "large" {
			cssRow[i] = strings.NewReplacer("{id}", strconv.Itoa(l.ID), "{url}", strings.NewReplacer(".jpg", "l.jpg").Replace(l.Image)).Replace(data.Style)
		} else {
			cssRow[i] = strings.NewReplacer("{id}", strconv.Itoa(l.ID), "{url}", l.Image).Replace(data.Style)
		}
	}

	return strings.Join(cssRow, "\n"), http.StatusOK, nil
}
