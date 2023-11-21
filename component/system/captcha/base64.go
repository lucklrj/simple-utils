package captcha

import (
	"sync"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/spf13/cast"
)

type Store struct {
	data sync.Map
}

func (s *Store) Set(id string, value string) error {
	defer func() {
		time.AfterFunc(5*time.Minute, func() {
			s.data.Delete(id)
		})
	}()
	s.data.Store(id, value)
	return nil
}
func (s *Store) Get(id string, clear bool) string {
	defer func() {
		if clear {
			s.data.Delete(id)
		}
	}()
	value, isExists := s.data.Load(id)
	if isExists {
		return cast.ToString(value)
	} else {
		return ""
	}
}
func (s *Store) Verify(id, answer string, clear bool) bool {
	return answer == s.Get(id, clear)
}

func MakeBase64CaptchaClient(length, height, width, dotCount int, maxSkew float64) *base64Captcha.Captcha {
	driver := base64Captcha.DriverDigit{}
	driver.Length = length
	driver.Height = height
	driver.Width = width
	driver.MaxSkew = maxSkew
	driver.DotCount = dotCount

	store := Store{}

	return base64Captcha.NewCaptcha(&driver, &store)

}
