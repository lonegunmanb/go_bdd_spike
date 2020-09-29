Feature: 登录登出淘宝账户
  为了能够购物
  作为一个淘宝用户
  我需要扫码登录淘宝以及登出淘宝账户

  Scenario: Display Login QRCode
    Given browse taobao homepage
    And click login
    When select login by qr code
    Then there should be qr code

  Scenario: Logout
    Given user already login at homepage
    When click logout
    Then user should be logout at homepage
