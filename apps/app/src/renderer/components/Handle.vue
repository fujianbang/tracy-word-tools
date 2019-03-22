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
.handle {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: space-between;
    .result-area {
        width: 60vw;
        height: 100%;
        padding: 20px;
        background-image: linear-gradient(to right, #43e97b 0%, #38f9d7 100%);
        textarea {
            width: 100%;
            height: 100%;
            overflow:auto;
        }
    }
    .upload-area {
        width: 40vw;
        border-left: 1px solid green;
        padding: 20px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        .footer {
            button {
            }
        }
    }
}
</style>
