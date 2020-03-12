[![微信小程序「安全登录」logo](https://cdn.jsdelivr.net/gh/secauth/docs@master/assets/secauth-logo.png)](https://github.com/secauth/docs)

# SecAuth

使用微信扫描二维码即可快速安全登录，接入简单灵活，轻松实现扫码登录。

- **快速**：扫码登录，无需填写
- **简单**：自助接入，申请即用
- **安全**：微信鉴权，安全可靠

## 登录流程

[![微信小程序「安全登录」登录流程](https://cdn.jsdelivr.net/gh/secauth/docs@master/assets/secauth-flow.png)](https://github.com/secauth/docs)

## 登录结果

[![微信小程序「安全登录」主题](https://cdn.jsdelivr.net/gh/secauth/docs@master/assets/secauth-theme.jpg)](https://github.com/secauth/docs)

- 保持简洁，避免打扰用户，确定即离开
- 自带多款主题配色
- 支持定制，由应用自行提供配色，与你的网站保持一致（计划中）

## 两步验证

应用提供了两步验证 TOTP，提高账号安全，可作为 Google Authenticator 等应用的替代。

## 谁在使用

- [Hamibot](https://hamibot.com/)
- 期待你的加入

## 如何使用

### 申请接入

打开微信小程序「安全登录」，点击申请接入。

[![微信小程序「安全登录」小程序码](https://cdn.jsdelivr.net/gh/secauth/docs@master/assets/secauth-wxacode.png)](https://github.com/secauth/docs)

### 获取小程序码

#### 请求格式

用实际值替换尖括号中的变量。

```http
GET /release/wxacode?id=<id> HTTP/1.1
Host: service-ggnj6gz0-1256804704.ap-hongkong.apigateway.myqcloud.com
Authorization: Bearer <token>
```

```sh
curl -X GET \
  "https://service-ggnj6gz0-1256804704.ap-hongkong.apigateway.myqcloud.com/release/wxacode?id=<id>" \
  -H "Authorization: Bearer <token>"
```

- id: 字符串，必填，最长 24 位，用户扫描小程序码后，会向开发者指定的地址发送请求，携带此 id 及用户凭证
- token：字符串，必填，填写「安全登录」提供的访问令牌

#### 响应格式

##### 正常
```
Content-Type: application/json

{
  status: "success",
  data: "小程序码 URL"
}
```
##### 异常
```
Content-Type: application/json

{
  status: "fail",
  message: "错误信息"
}
```

### 接收登录请求

当用户扫描小程序码后，「安全登录」会向开发者提供的回调地址发送一个 POST 请求，开发者可以获取到 token，使用「安全登录」提供的令牌密钥进行验证，payload 包含小程序码 id，用户凭证，用户信息等。

#### 请求格式
```
User-Agent: SecAuth-Hookshot/x
Content-Type: application/json

{
  token: "JWT TOKEN"
}
```
#### 响应格式

```
Content-Type: application/json

{
  status: "success" // 登录成功
}
```

### [代码样例](/samples)

## 计划

- [x] 增加回调地址测试
- [x] 获取用户微信头像和昵称
- [ ] 自定义结果页配色
- [ ] 访问令牌无需手动更新
- [ ] 自定义应用头像（小程序码头像）

## 其他

任何疑问，欢迎提 issue。

### 小程序码如何自定义头像

功能未开发完成，目前需要人工处理，请提 issue，注明应用头像，附带尺寸为 120x120 的图片以及令牌密钥前 6 位。

### 获取用户微信头像和昵称

会在登录时询问用户，用户允许后获得，支持获取的用户信息：

- nickName：昵称
- avatarUrl：头像
- gender：性别，0：未知、1：男、2：女
- province：省份
- city：城市
- country：国家

注意：开发者配置还未完成，请联系我启用此功能。
