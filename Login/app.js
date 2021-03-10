// app.js
App({
  globalData: {
    baseUrl:'http://127.0.0.1:2359/',
 },
  onShow(){
    wx.login({
      success :res => {
        wx.request({
          url: getApp().globalData.baseUrl + "login",
          method: 'GET',
          success(rsp){
            console.log(rsp)
          }
        })
      }
    })
  }
})
