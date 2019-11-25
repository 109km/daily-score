const { override, fixBabelImports, setWebpackPublicPath } = require('customize-cra');


let webpackPublicPath = ""
if (process.env.NODE_ENV === 'production') {
  webpackPublicPath = "https://ad-1251636704.file.myqcloud.com/admin-iqiyi-1029/";
}

module.exports = override(
  fixBabelImports('import', {
    libraryName: 'antd',
    libraryDirectory: 'es',
    style: 'css',
  }),
  webpackPublicPath && setWebpackPublicPath("https://ad-1251636704.file.myqcloud.com/admin-iqiyi-1029/")
);