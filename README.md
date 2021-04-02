# Beego Tutorial

 1. Basic Template https://github.com/nik-hil/beego-apiproject/releases/tag/v0.1
 1. QuerySet usage  https://github.com/nik-hil/beego-apiproject/releases/tag/v0.3
 1. Test case usage https://github.com/nik-hil/beego-apiproject/releases/tag/v0.5
 1. Validation usage https://github.com/nik-hil/beego-apiproject/releases/tag/v0.6
 1. Default routing not working https://github.com/nik-hil/beego-apiproject/releases/tag/v0.7

 ## Routing issue with Beego

 1. We have to be careful with routing
 1. in [v0.7](https://github.com/nik-hil/beego-apiproject/releases/tag/v0.7) GET  request with URL `/v1/user/login` will map to `r:/v1/user/:uid`
 1. in [v0.7](https://github.com/nik-hil/beego-apiproject/releases/tag/v0.7) POST  request with URL `/v1/user/login` will map to `r:/v1/user/login`