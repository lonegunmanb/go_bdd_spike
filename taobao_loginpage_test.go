package bdd_spike_test

import (
	"github.com/tebeka/selenium"
	"time"
)

type TaobaoLoginPage struct {
	*Page
}

func NewTaobaoLoginPage(wd selenium.WebDriver) *TaobaoLoginPage {
	return &TaobaoLoginPage{Page: &Page{Wd: wd}}
}

func (lp *TaobaoLoginPage) GetQRCode() (selenium.WebElement, error) {
	return lp.FindElementWithTimeout(selenium.ByXPATH, "//canvas", 10*time.Second)
}

func (lp *TaobaoLoginPage) SelectLoginByQRCode() error {
	qrCode, err := lp.FindElementWithTimeout(selenium.ByXPATH, "//i[@class='iconfont icon-qrcode']", 10*time.Second)
	if err != nil {
		return err
	}
	return qrCode.Click()
}
