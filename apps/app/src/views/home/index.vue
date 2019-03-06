<template>
    <div class="welcome">
        <div class="left">
            <div class="logo-area">
                <img :src="TracyImg" />
            </div>
            <div class="info-area">
                <h2>陈镜璇的Word检查工具</h2>
                <p>{{'Ver.' + version}}</p>
            </div>
            <p class="time">
            last update:
            <Time :time="lastUpdateDate"></Time>
            </p>
        </div>
        <div class="right">
            <div class="logo-area">
                <img :src="StefanieImg" />
            </div>
            <div class="button-area">
                <Button type="error" ghost @click="enterHandle">立即进入</Button>
            </div>
        </div>
    </div>
</template>
<script>
import TracyImg from '@/assets/tracy.jpg'
import StefanieImg from '@/assets/stefanie.jpg'
export default {
    data() {
        return {
            TracyImg,
            StefanieImg,
            version: '20190306',
            lastUpdateDate: new Date()
        }
    },
    methods: {
        /*
         * 进入处理页面
         */
        enterHandle() {
            this.$router.push({
                name: 'handle'
            })
        }
    },
    mounted() {
    }
}
</script>
<style lang="less" scoped>
@import './style.less';
</style>
