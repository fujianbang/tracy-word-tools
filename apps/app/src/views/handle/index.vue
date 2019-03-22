<template>
    <div class="handle">
        <div class="result-area">
            <textarea v-model="result">
            </textarea>
        </div>
        <div class="upload-area">
            <Upload
                multiple
                type="drag"
                action="http://tracy-word.0ooops.com/upload"
                :on-error="uploadError"
                :on-success="uploadSuccess">
                <div style="padding: 20px 0">
                    <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                    <p>Click or drag files here to upload</p>
                </div>
            </Upload>

            <div class="footer">
                <Button type="error" long @click="clearHistory">清除历史</Button>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            header: {
                'Access-Control-Allow-Origin': '*'
            },
            credentials: true,
            result: ''
        }
    },
    methods: {
        uploadSuccess(response, file, fileList) {
            this.result = response.data
            // response.filename
        },
        uploadError(error, file, fileList) {
            this.result = "出错啦"
            this.$Message.error({
                content: "出错啦，赶紧联系我！！！",
                duration: 10
            })
        },
        clearHistory() {
            this.result = "已清除..."
        }
    },
    mounted() {
    }
}
</script>
<style lang="less" scoped>
@import './style.less';
</style>
