module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:7002/',
                changeOrigin: true,
                onProxyReq: function (proxyReq, req, res) {
                    console.log(proxyReq.url, req.url, res)
             
                    // if (req.headers['authorization']) {
                    //     proxyReq.setHeader('authorization', req.headers['authorization']);
                    // }

                    // // 复制Cookie
                    // if (req.headers['cookie']) {
                    //     proxyReq.setHeader('cookie', req.headers['cookie']);
                    // }
                }
            }
        }
    },
    publicPath: '/web/dist',
    outputDir: "dist",
};
