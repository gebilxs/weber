module.exports = {
    assetsDir: "static",
    devServer: {
        proxy: {
            // '/api/v1': {
            //   target: 'http://127.0.0.1:8081',
            //   changeOrigin: true,
            // }
            '/api':{
              target:'http://127.0.0.1:8081',
              ws:true,
              changeOrigin:true,//允许跨域
              pathRewrite:{
                  '^/api':''   //请求的时候使用这个api就可以
              }
          }
        }
        
    }
  }
