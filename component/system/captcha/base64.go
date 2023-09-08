package captcha

import (
	"time"

	"github.com/mojocn/base64Captcha"
)

type Store struct {
	data map[string]string
}

func (s *Store) Set(id string, value string) error {
	defer func() {
		time.AfterFunc(5*time.Minute, func() {
			delete(s.data, id)
		})
	}()
	s.data[id] = value
	return nil
}
func (s *Store) Get(id string, clear bool) string {
	defer func() {
		if clear {
			delete(s.data, id)
		}
	}()
	return s.data[id]
}
func (s *Store) Verify(id, answer string, clear bool) bool {
	return answer == s.data[id]
}

func MakeBase64CaptchaClient() *base64Captcha.Captcha {
	driver := base64Captcha.DriverDigit{}
	driver.Length = 6
	driver.Height = 60
	driver.Width = 200
	driver.MaxSkew = 0.5
	driver.DotCount = 100

	store := Store{}
	store.data = make(map[string]string)

	return base64Captcha.NewCaptcha(&driver, &store)

}
