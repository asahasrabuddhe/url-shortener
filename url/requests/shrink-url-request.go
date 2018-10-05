package requests

import (
	"errors"
	"net/url"
)

type ShrinkURLRequest struct {
	URL string `json:"url"`
}

func (sur *ShrinkURLRequest) Validate() error {
	u, err := url.Parse(sur.URL)
	if err != nil {
		return errors.New("error: foo is not a valid URL")
	} else if u.Scheme == "" || u.Host == "" {
		return errors.New("error: foo must be an absolute URL")
	} else if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("error: foo must begin with http or https")
	}

	return nil
}
