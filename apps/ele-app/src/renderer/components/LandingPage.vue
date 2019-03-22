<template>
  <div class="welcome">
    <div class="left">
      <div class="logo-area">
        <img :src="TracyImg">
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
        <img :src="StefanieImg">
      </div>
      <div class="button-area">
        <Button type="primary" ghost @click="enterHandle">立即进入</Button>
      </div>
    </div>
  </div>
</template>

<script>
import TracyImg from "@/assets/tracy.jpg";
import StefanieImg from "@/assets/stefanie.jpg";
export default {
  data() {
    return {
      TracyImg,
      StefanieImg,
      version: "检查中",
      lastUpdateDate: new Date()
    };
  },
  methods: {
    /*
     * 进入处理页面
     */
    enterHandle() {
      this.$router.push({
        name: "handle"
      });
    }
  },
  mounted() {}
};
</script>

<style lang="less" scoped>
  .welcome {
    display: flex;
    width: 100vw;
    height: 100vh;
    .left {
        flex: 1 1 auto;
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        .logo-area img {
            width: 300px;
            height: 300px;
            display: inherit;
            margin: 0 auto;
        }
        .info-area {
            h2 {
                text-align: center;
            }
            p {
                text-align: center;
            }
        }
        p.time {
            text-align: center;
            margin-top: 20px;
        }
    }
    .right {
        flex: 0 0 400px;
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        background-image: linear-gradient(120deg, #84fab0 0%, #8fd3f4 100%);
        .logo-area img {
            width: 200px;
            height: 200px;
            display: inherit;
            margin: 0 auto;
            border-radius: 20px;
        }
        .button-area {
            button {
                width: 200px;
                margin: 0 auto;
                display: inherit;
            }
        }
    }
}
</style>
