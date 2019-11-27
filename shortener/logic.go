package shortener

import (
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrRedirectNotFound = errs.New("Redirect not found")
	ErrRedirectInvalid  = errs.New("Invalid redirect")
)

type redirectService struct {
	repository RedirectRepository
}

// Instantiates a new redirect service
func NewRedirectService(repo RedirectRepository) RedirectService {
	return &redirectService{
		repository: repo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.repository.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {

	// Validation logic goes here
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()

	return r.repository.Store(redirect)
}
