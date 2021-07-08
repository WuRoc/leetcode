---
description: 面试高频指数：★★★★☆
---

# DNS 的作用和原理

## **DNS**

**DNS**（Domain Name System）是域名系统的英文缩写，是一种组织成**域层次结构**的计算机和网络服务命名系统，用于 **TCP/IP 网络**。

### DNS 的作用

通常我们有**两种方式识别主机**：**通过主机名或者 IP 地址**。人们喜欢便于记忆的主机名表示，而路由器则喜欢定长的、有着层次结构的 IP 地址。为了满足这些不同的偏好，我们就需要一种能够进行主机名到 IP 地址转换的目录服务，域名系统作为将域名和 IP 地址相互映射的一个**分布式数据库**，能够使人更方便地访问互联网。

### DNS 域名解析原理

DNS 采用了分布式的设计方案，其域名空间采用一种树形的层次结构：

![](.gitbook/assets/image%20%2814%29.png)

上图展示了 DNS 服务器的部分层次结构，从上到下依次为**根域名服务器**、**顶级域名服务器**和**权威域名服务器**。其实根域名服务器在因特网上有13个，大部分位于北美洲。第二层为顶级域服务器，这些服务器负责顶级域名（如 com、org、net、edu）和所有国家的顶级域名（如uk、fr、ca 和 jp）。在第三层为权威 DNS 服务器，因特网上具有公共可访问主机（例如 Web 服务器和邮件服务器）的每个组织机构必须提供公共可访问的 DNS 记录，这些记录由组织机构的权威 DNS 服务器负责保存，这**些记录将这些主机的名称映射为 IP 地址**。

除此之外，还有一类重要的 DNS 服务器，叫做**本地 DNS 服务器**。本地 DNS 服务器严格来说不在 DNS 服务器的层次结构中，但它对 DNS 层次结构是很重要的。一般来说，每个网络服务提供商（ISP） 都有一台本地 DNS 服务器。当主机与某个 ISP 相连时，该 ISP 提供一台主机的 IP 地址，**该主机具有一台或多台其本地 DNS 服务器的 IP 地址**。**主机的本地 DNS 服务器通常和主机距离较近**，当主机发起 DNS 请求时，该请求被发送到本地 DNS 服务器，它起着代理的作用，并将该请求转发到 DNS 服务器层次结构中。

我们以一个例子来了解 DNS 的工作原理，假设主机 A（IP 地址为 abc.xyz.edu） 想知道主机 B 的 IP 地址 （def.mn.edu），如下图所示，**主机 A 首先向它的本地 DNS 服务器发送一个 DNS 查询报文**。该查询报文含有被转换的**主机名 def.mn.edu**。本地 DNS 服务器将该报文转发到根 DNS 服务器，根 DNS 服务器注意到查询的 IP 地址前缀为 edu 后向本地 DNS 服务器返回负责 edu 的顶级域名服务器的 **IP 地址列表**。该本地 DNS 服务器则再次向这些 顶级域名服务器发送查询报文。该顶级域名服务器注意到 mn.edu 的前缀，并用权威域名服务器的 IP 地址进行响应。通常情况下，顶级域名服务器并不总是知道每台主机的权威 DNS 服务器的 IP 地址，而只知道中间的某个服务器，该中间 DNS 服务器依次能找到用于相应主机的 IP 地址，我们假设中间经历了权威服务器 ① 和 ②，最后找到了负责 def.mn.edu 的权威 DNS 服务器 ③，之后，本地 DNS 服务器直接向该服务器发送查询报文从而获得主机 B 的IP 地址。

![](.gitbook/assets/image%20%2822%29.png)

在上图中，IP 地址的查询其实经历了两种查询方式，分别是递归查询和迭代查询。

拓展：域名解析查询的两种方式

**递归查询**：如果主机所询问的本地域名服务器不知道被查询域名的 IP 地址，那么本地域名服务器就以 DNS 客户端的身份，向其他根域名服务器继续发出查询请求报文，即替主机继续查询，而不是让主机自己进行下一步查询，如上图步骤（1）和（10）。 

**迭代查询**：当根域名服务器收到本地域名服务器发出的迭代查询请求报文时，要么给出所要查询的 IP 地址，要么告诉本地服务器下一步应该找哪个域名服务器进行查询，然后让本地服务器进行后续的查询，如上图步骤（2）~（9）。
