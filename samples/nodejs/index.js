const express = require('express');
const app = express();
const bodyParser = require('body-parser');
const jwt = require('jsonwebtoken');
const fetch = require('node-fetch');

// 令牌密钥
const secret = '';
// 访问令牌
const accessToken = '';

// 获取小程序码
const getWxaCode = async (id) => {
  const res = await fetch(`https://service-ggnj6gz0-1256804704.ap-hongkong.apigateway.myqcloud.com/release/wxacode?id=${id}`, {
    headers: { 'Authorization': `Bearer ${accessToken}` },
  });
  const body = await res.json();
  return body.data;
};

getWxaCode('abcd1234'); // 字符串 1-24 位

// 登录回调
app.use(bodyParser.json());

app.post('/', function (req, res) {
  const { token } = req.body;
  const { id, username, password } = jwt.verify(token, secret);
  res.json({
    status: 'success' // 登录成功
  });
});

app.listen(3000);