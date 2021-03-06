---
description: 面试高频指数：★★☆☆☆
---

# 客户端为什么信任第三方证书

假设中间人篡改了证书原文，由于他没有 CA 机构的私钥，所以无法得到此时加密后的签名，因此无法篡改签名。客户端浏览器收到该证书后会发现原文和签名解密后的值不一致，则说明证书被中间人篡改，证书不可信，从而终止向服务器传输信息。

上述过程说明证书无法被篡改，我们考虑更严重的情况，例如中间人拿到了 CA 机构认证的证书，它想窃取网站 A 发送给客户端的信息，于是它成为中间人拦截到了 A 传给客户端的证书，然后将其替换为自己的证书。此时客户端浏览器收到的是被中间人掉包后的证书，但由于证书里包含了客户端请求的网站信息，因此客户端浏览器只需要把证书里的域名与自己请求的域名比对一下就知道证书有没有被掉包了。

 公开密钥**加密**（**英语**：public-key cryptography，又译为公开密钥**加密**），也称为**非对称加密**（asymmetric cryptography），一种密码学算法类型，在这种密码学方法中，需要一对密钥\(其实这里密钥说法不好，就是“钥”\)，一个是私人密钥，另一个则是公开密钥。

