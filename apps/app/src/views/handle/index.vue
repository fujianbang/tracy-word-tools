<template>
    <div class="handle">
        <div class="result-area">
            {{ result }}
        </div>
        <div class="upload-area">
            <Upload
                multiple
                type="drag"
                action="//jsonplaceholder.typicode.com/posts/">
                <div style="padding: 20px 0">
                    <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                    <p>Click or drag files here to upload</p>
                </div>
            </Upload>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            result: ''
        }
    },
    methods: {
    },
    mounted() {
        this.result = ''+
            '!!! 5' +
            'html' +
            '  include header' +
            '  body' +
            '    //if IE 6' +
            '        .alert.alert-error' +
            '            center 对不起，我们不支持IE6，请升级你的浏览器' +
            '                a(href="http://windows.microsoft.com/zh-CN/internet-explorer/download-ie") | IE8官方下载' +
            '                a(href="https://www.google.com/intl/en/chrome/browser/") | Chrome下载' +
            '    include head' +
            '    .container' +
            '        .row-fluid' +
            '            .span8' +
            '                block main' +
            '                include pagerbar' +
            '            .span4' +
            '                include sidebar' +
            '    include footer' +
            '    include script'
    }
}
</script>
<style lang="less" scoped>
@import './style.less';
</style>
