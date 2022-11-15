<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast, POSITION } from 'vue-toastification'

import { createUser } from '../api/user'
import FloatingAnimation from '../components/FloatingAnimation.vue'

const toast = useToast()
const router = useRouter()

const registerForm = ref({
  nickname: '',
  username: '',
  password: '',
  comfirmPassword: ''
})

function doRegister() {
  if (registerForm.value.username.length < 2) {
    toast.error('用户名不能少于2个字符哦。', {
      position: POSITION.TOP_RIGHT
    })
    return
  }
  if (registerForm.value.password == '') {
    toast.error('密码不能为空。', {
      position: POSITION.TOP_RIGHT
    })
    return
  }
  if (registerForm.value.password.length < 6) {
    toast.error('密码不能少于6个字符哦。', {
      position: POSITION.TOP_RIGHT
    })
    return
  }
  if (registerForm.value.comfirmPassword == '') {
    toast.error('确认密码不能为空。', {
      position: POSITION.TOP_RIGHT
    })
    return
  }
  if (registerForm.value.password != registerForm.value.comfirmPassword) {
    toast.error('两次密码不一致，请检查后重试。', {
      position: POSITION.TOP_RIGHT
    })
    return
  }

  createUser(registerForm.value)
    .then((res) => {
      if (res.data.code != '200') {
        toast.error(`注册失败，${res.data.msg}`, {
          position: POSITION.TOP_RIGHT
        })
        return
      }
      toast.success(`用户${registerForm.value.username}注册成功。`, {
        position: POSITION.TOP_RIGHT
      })
      router.push({
        name: 'Login'
      })
    })
    .catch((err) => {
      console.log('aaaaaaaa', err.data)
      toast.error(`用户${registerForm.value.username}注册失败，错误${err}`)
    })
}
</script>
<template>
  <div class="container">
    <div class="tit">注册Pigeon账号</div>
    <input type="text" v-model="registerForm.nickname" placeholder="昵称" />
    <input type="text" v-model="registerForm.username" placeholder="账号（不少于2个字符）" />
    <input type="password" v-model="registerForm.password" placeholder="密码（不少于6个字符）" />
    <input type="password" v-model="registerForm.comfirmPassword" placeholder="确认密码" />
    <button @click="doRegister">注册</button>
    <span>已有账号？马上<RouterLink to="/login">登录</RouterLink></span>
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
  width: 400px;
  height: 560px;
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
  height: 50px;
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

.container span {
  float: start;
  font-size: 14px;
  color: #999;
}

.container button:hover {
  background: linear-gradient(-200deg, #aac2ee, #fac0e7);
  background-position-x: -280px;
}

.container span {
  font-size: 14px;
  margin-bottom: 6px;
}

.container a {
  color: plum;
  text-decoration: none;
}
</style>
