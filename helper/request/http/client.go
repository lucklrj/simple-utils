package http

import (
	"net/http"

	"github.com/ddliu/go-httpclient"
)

var Client *httpClient

type httpClient struct {
	handler *httpclient.HttpClient
}

func (c *httpClient) Get(url string, httpProxyUrl string, headers map[string]string, params ...interface{}) (string, error) {
	var result *httpclient.Response
	var err error

	//根据设置代理
	obj := c.handler.Begin()
	if httpProxyUrl != "" {
		obj = obj.WithOption(httpclient.OPT_PROXY, httpProxyUrl)
	}
	if headers != nil {
		obj = obj.WithHeaders(headers)
	}

	result, err = obj.Get(url, params...)
	if err != nil {
		return "", err
	} else {
		response, _ := result.ToString()
		return response, nil
	}
}
func (c *httpClient) Post(url string, headers map[string]string, postData map[string]interface{},
	cookies []*http.Cookie) (body string,
	errs error) {
	hc := c.handler
	res, err := hc.Begin().WithHeaders(headers).WithCookie(cookies...).Post(url, postData)
	if err != nil {
		//color.Red(err.Error())
		return "", err
	}
	return res.ToString()
}

func (c *httpClient) PostJson(url string, postData map[string]interface{}, headers map[string]string) (body string,
	errs error) {
	res, err := c.handler.Begin().WithHeaders(headers).PostJson(url, postData)
	if err != nil {
		return "", err
	}
	return res.ToString()
}
func (c *httpClient) GetCookie(url string) map[string]string {
	httpCookie := make(map[string]string)
	for _, cookie := range c.handler.Cookies(url) {
		httpCookie[cookie.Name] = cookie.Value
	}
	return httpCookie
}

func MakeClient() {
	Client = &httpClient{
		handler: httpclient.NewHttpClient().Defaults(httpclient.Map{
			//httpclient.OPT_REFERER:   "https://www.jianshu.com/writer",
			httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:64. 0) Gecko/20100101 Firefox/64.0",
			//httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Windows NT 6.1; rv:24.0) Gecko/20100101 Firefox/24.0",
			httpclient.OPT_UNSAFE_TLS: true,
		}),
	}
}
