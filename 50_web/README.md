# Web 框架

一个 Web  程序的编写往往要涉及更多的方面，目前各种各样的中间件能够完成一些任务。但许多时候，我们总是希望他人帮我们完成更多的事情，于是就产生了许多的 Web 框架。根据架构的不同，这些框架大致可分为两大类：

- 微架构型框架：其核心框架只提供很少的功能，而更多的功能则需要组合各种中间件来提供，因此这种框架也可称为混搭型框架。它相当灵活，但相对来说需要使用者在组合使用各种中间件时花费更大的力气。像 [Echo](https://github.com/labstack/echo)、[Goji](https://github.com/goji/goji)、[Gin](https://github.com/gin-gonic/gin) 等都属于微架构型框架。
- 全能型架构：它基本上提供了编写 Web 应用时需要的所有功能，因此更加重型，多数使用 MVC 架构模式设计。在使用这类框架时感觉更轻省，但其做事风格一般不同于 Go 语言惯用的风格，也较难弄明白这些框架是如何工作的。像 [Beego](http://beego.me/)、[Revel](http://revel.github.io/) 等就属于全能型架构。

对于究竟该选择微架构还是全能型架构，仍有较多的争议。像 [The Case for Go Web Frameworks](https://medium.com/@richardeng/the-case-for-go-web-frameworks-a791fcd79d47#.7qe9n08aw) 一文就力挺全能型架构，并且其副标题就是“Idiomatic Go is not a religion”，但该文也收到了较多的反对意见，见[这里](https://groups.google.com/forum/#!searchin/golang-nuts/framework/golang-nuts/vX086U_49Qo/KLXcyKwVil4J)和[这里](https://www.reddit.com/r/programming/comments/2jsrsq/the_case_for_go_web_frameworks_idiomatic_go_is/)。总体上来说，Go 语言社区已越来越偏向使用微架构型框架，当将来 `context` 包进入标准库后，`http.Handler` 本身就定义了较完善的中间件编写规范，这种使用微架构的趋势可能更加明显，并且各种微架构的实现方式有望进一步走向统一，这样其实 `http` 包就是一个具有庞大生态系统的微架构框架。