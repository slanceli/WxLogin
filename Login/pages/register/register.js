// pages/register/register.js
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },
  onShow: function () {
    var regUrl = app.globalData.baseUrl + "register";
    wx.request({
      url: regUrl,
      method: 'GET',
      success(rsp){
        console.log(rsp)
      }
    })
  },
//进入登录页面
login: function(e) {
  wx.redirectTo({
    url: '../index/index',
    fail :rsp=>{
      console.log(rsp)
    },
    success :res =>{
      console.log(res)
    }
  })
},
//注册
submitBindingre(e){
  var body = e.detail.value
  var registerUrl = app.globalData.baseUrl + 'register'
  wx.request({
    url: registerUrl,
    data:{
      name: body['Name'],
      passwd: body['Password'],
    },
    method: 'POST',
    header: {
      'content-type': 'application/x-www-form-urlencode'
    },
    dataType: String,
    success :res =>{
      console.log(res.data)
      if(res.data == "注册成功"){
        wx.showModal({
          cancelColor: '#000000',
          title: "注册成功",
          content: "注册成功，即将跳转到登录页面",
          showCancel: false,
          success :res=>{
            wx.redirectTo({
              url: '../index/index',
            })
          }
        })
      }
      else if(res.data == "注册失败"){
        wx.showModal({
          cancelColor: '#000000',
          title: res.data,
          content: "注册失败，请联系管理员",
          showCancel: false,
        })
      }
      else if(res.data == "用户名或密码不能为空"){
        wx.showModal({
          cancelColor: '#000000',
          title: "注册失败",
          content: res.data,
          showCancel: false
        })
      }
    },
  })
}
})