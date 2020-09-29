package bdd_spike_test

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"net/url"
	"os"
)

var hp *TaobaoHomepage
var lp *TaobaoLoginPage
var wd selenium.WebDriver

func browseTaobaoHomePage() error {
	hp = NewTaobaoHomepage(wd)
	return hp.Navigate()
}

func userLogin() error {
	err := hp.Login()
	if err != nil {
		return err
	}
	lp = NewTaobaoLoginPage(wd)
	return nil
}

func selectLoginByQrCode() error {
	return lp.SelectLoginByQRCode()
}

func thereShouldBeQrCode() error {
	_, err := lp.GetQRCode()
	return err
}

func userAlreadyLoginAtHomepage() error {
	var err error
	hp = NewTaobaoHomepage(wd)
	err = hp.Navigate()
	if err != nil {
		return err
	}

	return hp.LoginByCookie()
}

func logout() error {
	return hp.Logout()
}

func userShouldAtHomepageAsAnonymous() error {
	currentUrl, err := wd.CurrentURL()
	if err != nil {
		return err
	}
	parse, err := url.Parse(currentUrl)
	if err != nil {
		return err
	}
	if parse.Host != "www.taobao.com" {
		return fmt.Errorf("expected www.taobao.com, got %s", currentUrl)
	}
	hp = NewTaobaoHomepage(wd)
	return hp.CurrentUserIsAnonymous()
}

func FeatureContext(s *godog.ScenarioContext) {
	s.BeforeScenario(func(sc *godog.Scenario) {
		caps := selenium.Capabilities{"browserName": "chrome"}
		chromeCaps := chrome.Capabilities{
			W3C:  false,
			Args: []string{"--headless"},
		}
		caps.AddChrome(chromeCaps)
		var err error
		wd, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
		if err != nil {
			panic(err)
		}
	})

	s.Step(`^browse taobao homepage$`, browseTaobaoHomePage)
	s.Step(`^click login$`, userLogin)
	s.Step(`^select login by qr code$`, selectLoginByQrCode)
	s.Step(`^there should be qr code$`, thereShouldBeQrCode)

	s.Step(`^user already login at homepage$`, userAlreadyLoginAtHomepage)
	s.Step(`^click logout$`, logout)
	s.Step(`^user should be logout at homepage$`, userShouldAtHomepageAsAnonymous)

	s.AfterScenario(func(s *godog.Scenario, err error) {
		err = wd.Quit()
		if err != nil {
			panic(err.Error())
		}
	})
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress", // can define default values
	})
}
