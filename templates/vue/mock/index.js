const Mock = require('mockjs');//mockjs 导入依赖模块

const UserInfo = {
    "error": 0,
    "data": {
        "userid": "@id()",
        "username": "@cname()",
        "date": "@date()",
        "avatar": "@image('200x200','red','#fff','avatar')",
        "description": "@paragraph()",
        "ip": "@ip()",
        "email": "@email()"
    }
}

//返回一个函数
module.exports = function (app) {
    //监听http请求
    app.get('/user/userinfo', function (rep, res) {
        res.json(Mock.mock(UserInfo));
    });
}