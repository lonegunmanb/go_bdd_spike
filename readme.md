首先要安装selenuim server以及chrome webdriver，其次启动selenuim server，然后安装[godog](github.com/cucumber/godog)。

执行登出测试需要先手动登录淘宝，登录后使用名为"EditThisCookie"的Chrome插件将.taobao.com的Cookie导出到项目目录下名为"cookie.json"的文件才能顺利执行。

执行测试：

```shell script
godog
Feature: 登录登出淘宝账户
  为了能够购物
  作为一个淘宝用户
  我需要扫码登录淘宝以及登出淘宝账户

  Scenario: Display Login QRCode # features/login.feature:6
    Given browse taobao homepage # login_test.go:18 -> _test.browseTaobaoHomePage
    And click login              # login_test.go:23 -> _test.userLogin
    When select login by qr code # login_test.go:32 -> _test.selectLoginByQrCode
    Then there should be qr code # login_test.go:36 -> _test.thereShouldBeQrCode

  Scenario: Logout                         # features/login.feature:12
    Given user already login at homepage   # login_test.go:41 -> _test.userAlreadyLoginAtHomepage
    When click logout                      # login_test.go:52 -> _test.logout
    Then user should be logout at homepage # login_test.go:56 -> _test.userShouldAtHomepageAsAnonymous

2 scenarios (2 passed)
7 steps (7 passed)
```
