package bdd_spike_test

import (
	"github.com/tebeka/selenium"
	"time"
)

type TaobaoHomepage struct {
	*Page
}

func (hp *TaobaoHomepage) LoginByCookie() error {
	cookies, err := readLoggedInUserCookies()
	if err != nil {
		return err
	}
	err = hp.AddCookies(cookies)
	if err != nil {
		return err
	}
	err = hp.Refresh()
	if err != nil {
		return err
	}
	return err
}

func (hp *TaobaoHomepage) Login() error {
	href, err := hp.FindElementWithTimeout(selenium.ByXPATH, "//a[.='亲，请登录']", 10*time.Second)
	if err != nil {
		return err
	}
	err = href.Click()
	if err != nil {
		return err
	}

	return hp.WaitForNavigateToUrlContains("login.taobao.com", 10*time.Second)
}

func (hp *TaobaoHomepage) Logout() error {
	err := hp.PopupLogoutRef()
	if err != nil {
		return err
	}
	logoutRef, err := hp.FindElementWithTimeout(selenium.ByXPATH, "//a[contains(@href, 'logout.jhtml')]", 10*time.Second)
	if err != nil {
		return err
	}
	err = logoutRef.Click()
	if err != nil {
		return err
	}
	return err
}

func (hp *TaobaoHomepage) PopupLogoutRef() error {
	userNameRef, err := hp.FindElementWithTimeout(selenium.ByXPATH, "//a[@class='site-nav-login-info-nick ' and .!='undefined']", 10*time.Second)
	if err != nil {
		return err
	}
	err = hp.HoverElement(userNameRef, 5, 5)
	if err != nil {
		return err
	}

	_, err = hp.FindElementWithTimeout(selenium.ByXPATH, "//a[contains(@href, 'logout.jhtml')]", 10*time.Second)
	return err
}

func NewTaobaoHomepage(wd selenium.WebDriver) *TaobaoHomepage {
	hp := &TaobaoHomepage{Page: &Page{Wd: wd}}
	return hp
}

func (hp *TaobaoHomepage) Navigate() error {
	return hp.Wd.Get("http://www.taobao.com")
}

func (hp *TaobaoHomepage) CurrentUserIsAnonymous() error {
	_, err := hp.FindElementWithTimeout(selenium.ByXPATH, "//a[.='亲，请登录']", 10*time.Second)
	return err
}
