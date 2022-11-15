<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast, POSITION } from 'vue-toastification'

import { getToken } from '../api/auth'
import FloatingAnimation from '../components/FloatingAnimation.vue'

const toast = useToast()
const router = useRouter()

const loginForm = ref({
  username: '',
  password: ''
  // rememberMe: false,
})

function doLogin() {
  getToken(loginForm.value)
    .then((res) => {
      // set localstorage
      localStorage.setItem('userUUID', res.data.data.userUUID)
      localStorage.setItem('displayName', res.data.data.displayName)
      localStorage.setItem('displayAvatar', res.data.data.displayAvatar)
      localStorage.setItem('token', res.data.data.token)

      router.push({
        name: 'Home'
      })
    })
    .catch((err) => {
      toast.error('登录失败，服务异常或网络错误。', {
        position: POSITION.TOP_RIGHT
      })
    })
}
</script>
<template>
  <div class="container">
    <div class="tit">登录到Pigeon</div>
    <input type="text" v-model="loginForm.username" placeholder="账号" />
    <input type="password" v-model="loginForm.password" placeholder="密码" />
    <button @click="doLogin">登录</button>
    <span>没有账号？前往<RouterLink to="/register">注册</RouterLink></span>
  </div>
  <FloatingAnimation></FloatingAnimation>
</template>

<style lang="css" scoped>
.container {
  /* 相对定位 */
  position: relative;
  z-index: 1;
  background-color: #fff;
  border-radius: 15px;
  /* 弹性布局 垂直排列 */
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 350px;
  height: 500px;
  /* 阴影 */
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
}

.container .tit {
  font-size: 26px;
  margin: 65px auto 70px auto;
}

.container input {
  width: 280px;
  height: 30px;
  text-indent: 8px;
  border: none;
  border-bottom: 1px solid #ddd;
  outline: none;
  margin: 12px auto;
}

.container button {
  width: 280px;
  height: 40px;
  margin: 35px auto 40px auto;
  border: none;
  background: linear-gradient(-200deg, #fac0e7, #aac2ee);
  color: #fff;
  font-weight: bold;
  letter-spacing: 8px;
  border-radius: 10px;
  cursor: pointer;
  /* 动画过渡 */
  transition: 0.5s;
}

.container button:hover {
  background: linear-gradient(-200deg, #aac2ee, #fac0e7);
  background-position-x: -280px;
}

.container span {
  font-size: 14px;
}

.container a {
  color: plum;
  text-decoration: none;
}
</style>
