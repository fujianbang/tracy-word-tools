const BASE_URL = ''

module.exports = {
    publicPath: '/',
    devServer: {
        proxy: 'http://127.0.0.1:8800'
    }
}