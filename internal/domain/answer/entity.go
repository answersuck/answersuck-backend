package answer

import (
	"errors"

	"github.com/answersuck/host/internal/domain/media"
	"github.com/answersuck/host/internal/pkg/pagination"
)

var (
	ErrMediaTypeNotAllowed = errors.New("not allowed media type for answer")
	ErrLanguageNotFound    = errors.New("language with provided id not found")
)

type Answer struct {
	Id         int     `json:"id"`
	Text       string  `json:"text"`
	MediaId    *string `json:"media_id"`
	LanguageId uint    `json:"language_id"`
}

var allowedMediaType = [3]media.Type{media.TypeImageJPEG, media.TypeImagePNG, media.TypeImageWEBP}

// mediaTypeAllowed checks if media for answer is in allowed media type array
func mediaTypeAllowed(mt string) bool {
	var allowed bool
	for _, t := range allowedMediaType {
		if media.Type(mt) == t {
			allowed = true
			break
		}
	}
	return allowed
}

type Filter struct {
	Text       string
	LanguageId uint
}

type ListParams struct {
	Pagination pagination.Params
	Filter     Filter
}

const maxLimit = 100

func NewListParams(lastId uint32, limit uint64, f Filter) ListParams {
	if limit == 0 || limit > maxLimit {
		limit = pagination.DefaultLimit
	}
	return ListParams{
		Pagination: pagination.Params{
			LastId: lastId,
			Limit:  limit,
		},
		Filter: f,
	}
}
