package bdd_spike_test

import (
	"encoding/json"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"strings"
	"time"
)

type Page struct {
	Wd selenium.WebDriver
}

func (p *Page) Refresh() error {
	return p.Wd.Refresh()
}

func (p *Page) WaitForNavigateToUrlContains(keyword string, timeout time.Duration) error {
	return p.Wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		currentURL, err := wd.CurrentURL()
		if err != nil {
			return false, nil
		}
		return strings.Contains(currentURL, keyword), nil
	}, timeout)
}

func (p *Page) FindElementWithTimeout(by, value string, timeout time.Duration) (selenium.WebElement, error) {
	err := p.Wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(by, value)
		return err == nil, nil
	}, timeout)

	if err != nil {
		return nil, err
	}

	element, err := p.Wd.FindElement(by, value)
	if err != nil {
		return nil, err
	}
	return element, nil
}

func (p *Page) AddCookies(cookies []Cookie) error {
	for _, cookie := range cookies {
		err := hp.Wd.AddCookie(&selenium.Cookie{
			Name:   cookie.Name,
			Value:  cookie.Value,
			Path:   cookie.Path,
			Domain: cookie.Domain,
			Secure: cookie.Secure,
			Expiry: uint(cookie.Expiry),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func readLoggedInUserCookies() ([]Cookie, error) {
	loginUserCookies, err := ioutil.ReadFile("cookie.json")
	if err != nil {
		return nil, err
	}
	var cookies []Cookie
	err = json.Unmarshal(loginUserCookies, &cookies)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}

func (*Page) HoverElement(element selenium.WebElement, xOffset, yOffset int) error {
	return element.MoveTo(xOffset, yOffset)
}

type Cookie struct {
	Domain string  `json:"domain"`
	Name   string  `json:"name"`
	Secure bool    `json:"secure"`
	Expiry float64 `json:"expirationDate"`
	Path   string  `json:"path"`
	Value  string  `json:"value"`
}
