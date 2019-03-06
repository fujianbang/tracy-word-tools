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
                action="http://127.0.0.1:9090/upload"
                :on-success="uploadSuccess">
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
        }
    },
    mounted() {
    }
}
</script>
<style lang="less" scoped>
@import './style.less';
</style>
