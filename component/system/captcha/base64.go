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

func MakeBase64CaptchaClient() *base64Captcha.Captcha {
	driver := base64Captcha.DriverDigit{}
	driver.Length = 6
	driver.Height = 60
	driver.Width = 200
	driver.MaxSkew = 0.5
	driver.DotCount = 100

	store := Store{}

	return base64Captcha.NewCaptcha(&driver, &store)

}
