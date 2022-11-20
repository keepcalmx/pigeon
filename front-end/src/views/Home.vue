<script setup lang="ts">
import { onMounted, onUnmounted, ref, computed, onUpdated } from 'vue'
import { useRouter } from 'vue-router'
import { POSITION, useToast } from 'vue-toastification'
import moment from 'moment'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { baseURL } from '../api/common'
import { getUserContactList, getUserInfo } from '../api/user'
import { listGroupUsers } from '../api/group'

moment.locale('zh-cn', {
  relativeTime: {
    future: '%s内',
    past: '%s前',
    s: '几秒',
    ss: '%d秒',
    m: '1分钟',
    mm: '%d分钟',
    h: '1小时',
    hh: '%d小时',
    d: '1天',
    dd: '%d天',
    M: '1个月',
    MM: '%d个月',
    y: '1年',
    yy: '%d年'
  }
})

const toast = useToast()
const router = useRouter()
const msgWebSocket = ref()

// 用户信息
const userUUID = localStorage.getItem('userUUID') || ''
const displayName = localStorage.getItem('displayName')
const displayAvatar = localStorage.getItem('displayAvatar') || '../assets/avatar.png'

// 联系人状态
const activedChatUUID = ref('')
const activedChatType = ref('')
const recentContactInfo = ref(new Map())
const userInfo = ref(new Map())

// 用户群组所属关系
const affiliation = ref(new Map())

// 输入内容
const chatText = ref()

const currentChatName = computed(() => {
  return recentContactInfo.value.get(activedChatUUID.value)?.displayName
})

const currentMsgList = computed(() => {
  return recentContactInfo.value.get(activedChatUUID.value)?.recentMsg || []
})

const firstThreeUser = computed(() => {
  if (!affiliation.value.has(activedChatUUID.value)) {
    pullGroupInfo(activedChatUUID.value)
  }
  return affiliation.value.get(activedChatUUID.value)?.slice(0, 3) || []
})

const otherUserCount = computed(() => {
  const count = affiliation.value.get(activedChatUUID.value)?.length - 3
  if (count < 0) {
    return 0
  }
  return count
})

function avatar(uuid: string) {
  if (uuid == userUUID) {
    return displayAvatar
  }
  if (recentContactInfo.value.has(uuid)) {
    return recentContactInfo.value.get(uuid).avatar
  }
  if (userInfo.value.has(uuid)) {
    return userInfo.value.get(uuid).avatar
  }
  getUserInfo(uuid)
    .then((res: any) => {
      userInfo.value.set(uuid, res.data.data)
      return res.data.data.avatar
    })
    .catch((err: any) => {
      console.log(`获取用户信息失败，${err}`)
    })
}

// 注册websocket客户端
async function registerMsgWebSocket() {
  // 注册需要校验token
  const params = `token=${localStorage.getItem('token') || ''}`

  msgWebSocket.value = new WebSocket(`ws://${baseURL}/msg?${params}`)

  msgWebSocket.value.addEventListener('open', (event: any) => {
    console.log('websocket register success.', event)
  })

  msgWebSocket.value.addEventListener('close', (event: any) => {
    console.log('websocket closed.', event)
  })

  msgWebSocket.value.addEventListener('error', (event: any) => {
    console.log('websocket register error: ', event)
  })

  msgWebSocket.value.addEventListener('message', (event: any) => {
    if (!recentContactInfo.value) {
      return
    }
    const data = JSON.parse(event.data)
    switch (data.type) {
      case 'message':
        if (data.data.toType === 'group') {
          recentContactInfo.value.get(data.data.to).recentMsg.push(data.data)
        } else if (data.data.toType === 'user') {
          recentContactInfo.value.get(data.data.from).recentMsg.push(data.data)
        }
        break
      case 'status':
        if (data.data.target == 'online') {
          // 更新在线状态
          recentContactInfo.value.get(data.data.uuid).online = data.data.value
        }
        break
      default:
        console.error('unknown type: ', data.type)
    }
  })
}

async function unregisterMsgWebSocket() {
  if (msgWebSocket) {
    msgWebSocket.value.close()
  }
}

async function doSendMsg() {
  if (chatText.value == '') {
    return
  }
  const newMsg = {
    type: 'text',
    content: chatText.value,
    from: userUUID,
    toType: activedChatType.value,
    to: activedChatUUID.value
  }

  try {
    msgWebSocket.value.send(JSON.stringify(newMsg))
    recentContactInfo.value.get(activedChatUUID.value).recentMsg.push(newMsg)
    chatText.value = ''
  } catch (error) {
    console.error('send msg error: ', error)
    toast.error(`发送消息失败，错误：${error}`, {
      position: POSITION.TOP_RIGHT
    })
  }
}

async function doLogout() {
  unregisterMsgWebSocket()

  localStorage.removeItem('displayName')
  localStorage.removeItem('displayAvatar')
  localStorage.removeItem('token')

  router.push({
    name: 'Login'
  })
}

// 获取联系人列表
async function getContactList() {
  getUserContactList(userUUID)
    .then((res) => {
      // 转为 uuid -> info 的字典
      res.data.data.forEach((item: any) => {
        recentContactInfo.value.set(item.uuid, item)
      })
      // 默认启用第一个联系人的聊天窗
      // TODO 优化为上次聊天的联系人
      if (res.data.data.length > 0) {
        activedChatUUID.value = res.data.data[0].uuid
        activedChatType.value = res.data.data[0].type
      }
    })
    .catch((err) => {
      toast.error(`获取联系列表失败，错误：${err}`, {
        position: POSITION.TOP_RIGHT
      })
    })
}

function pullGroupInfo(uuid: string) {
  listGroupUsers(uuid).then((res) => {
    affiliation.value.set(uuid, [])
    console.log('listGroupUsers', res)
    res.data.data.forEach((item: any) => {
      affiliation.value.get(uuid).push(item.uuid)
      userInfo.value.set(item.uuid, {
        displayName: item.nickname || item.username,
        avatar: item.avatar
      })
    })
  })
}

function changeActivedChatStatus(uuid: any) {
  activedChatUUID.value = uuid
  activedChatType.value = recentContactInfo.value.get(uuid).type
  // TODO 更新未读消息数
}

function initThemeChange() {
  const toggleButton = document.querySelector('.dark-light')
  const colors = document.querySelectorAll('.color')

  colors.forEach((color) => {
    color.addEventListener('click', (e) => {
      colors.forEach((c) => c.classList.remove('selected'))
      const theme = color.getAttribute('data-color') || ''
      document.body.setAttribute('data-theme', theme)
      color.classList.add('selected')
    })
  })

  toggleButton?.addEventListener('click', () => {
    document.body.classList.toggle('dark-mode')
  })
}

function showAddExtend() {
  toast.info(`在做了在做了`, {
    position: POSITION.TOP_RIGHT
  })
}

function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

onUpdated(() => {
  // auto scroll to bottom
  let chatContent = document.querySelector('#chat-area')
  if (chatContent) {
    chatContent.scrollTop = chatContent.scrollHeight
  }
})

onMounted(() => {
  initThemeChange()
  getContactList()
  sleep(1000).then(() => {
    registerMsgWebSocket()
  })
})

onUnmounted(() => {
  unregisterMsgWebSocket()
    .then(() => {})
    .catch(() => {})
})
</script>
<template>
  <div class="app">
    <!-- 顶栏 -->
    <div class="header">
      <!-- 搜索框 -->
      <div class="search-bar">
        <input type="text" placeholder="Search..." />
      </div>
      <div class="user-settings">
        <div class="dark-light">
          <svg
            viewBox="0 0 24 24"
            stroke="currentColor"
            stroke-width="1.5"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z" />
          </svg>
        </div>
        <div class="settings">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <circle cx="12" cy="12" r="3" />
            <path
              d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 010 2.83 2 2 0 01-2.83 0l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-2 2 2 2 0 01-2-2v-.09A1.65 1.65 0 009 19.4a1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83 0 2 2 0 010-2.83l.06-.06a1.65 1.65 0 00.33-1.82 1.65 1.65 0 00-1.51-1H3a2 2 0 01-2-2 2 2 0 012-2h.09A1.65 1.65 0 004.6 9a1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 010-2.83 2 2 0 012.83 0l.06.06a1.65 1.65 0 001.82.33H9a1.65 1.65 0 001-1.51V3a2 2 0 012-2 2 2 0 012 2v.09a1.65 1.65 0 001 1.51 1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 0 2 2 0 010 2.83l-.06.06a1.65 1.65 0 00-.33 1.82V9a1.65 1.65 0 001.51 1H21a2 2 0 012 2 2 2 0 01-2 2h-.09a1.65 1.65 0 00-1.51 1z"
            />
          </svg>
        </div>
        <div class="logout" @click="doLogout">
          <font-awesome-icon icon="fa-solid fa-arrow-right-from-bracket" />
        </div>
        <img class="user-profile" :src="displayAvatar" alt="" />
        <span class="profile-name">{{ displayName }}</span>
      </div>
    </div>
    <div class="wrapper">
      <!-- 左侧联系人栏 -->
      <div class="conversation-area">
        <div
          v-for="([key, obj], idx) in recentContactInfo"
          :key="idx"
          @click="changeActivedChatStatus(obj.uuid)"
          :class="[{ online: obj.online, offline: !obj.online, active: obj.uuid === activedChatUUID }, 'msg']"
        >
          <img class="msg-profile" :src="obj.avatar" />
          <div class="msg-detail">
            <div class="msg-username">
              {{ obj.displayName }}
              <!-- 未读消息提醒 -->
              <!-- <span class="c-number">4</span> -->
            </div>
            <div class="msg-content">
              <span v-if="obj.recentMsg.length == 0" class="msg-message">暂无消息</span>
              <span v-else class="msg-message">{{
                obj.recentMsg[obj.recentMsg.length - 1].content
              }}</span>
              <span class="msg-date" v-if="obj.recentMsg.length > 1">
                {{ moment(obj.recentMsg[obj.recentMsg.length - 1].created_at).fromNow() }}
              </span>
            </div>
          </div>
        </div>
        <button class="add" @click="showAddExtend"></button>
        <div class="overlay"></div>
      </div>
      <!-- 右侧聊天框 -->
      <div class="chat-area" id="chat-area">
        <div class="chat-area-header">
          <div class="chat-area-title">{{ currentChatName }}</div>
          <div v-if="activedChatType == 'group'" class="chat-area-group">
            <img
              v-for="uuid in firstThreeUser"
              class="chat-area-profile"
              :src="avatar(uuid)"
              alt=""
            />
            <span v-if="otherUserCount > 0">+{{ otherUserCount }}</span>
          </div>
        </div>
        <div class="chat-area-main">
          <div
            v-for="(msg, idx) in currentMsgList"
            :key="idx"
            :class="[{ owner: msg.from == userUUID }, 'chat-msg']"
          >
            <div class="chat-msg-profile">
              <img class="chat-msg-img" :src="avatar(msg.from)" alt="" />
              <div class="chat-msg-date">{{ moment(msg.created_at).fromNow() }}</div>
            </div>
            <div class="chat-msg-content">
              <div class="chat-msg-text">{{ msg.content }}</div>
            </div>
          </div>
        </div>
        <div class="chat-area-footer">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="feather feather-video"
          >
            <path d="M23 7l-7 5 7 5V7z" />
            <rect x="1" y="5" width="15" height="14" rx="2" ry="2" />
          </svg>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="feather feather-image"
          >
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2" />
            <circle cx="8.5" cy="8.5" r="1.5" />
            <path d="M21 15l-5-5L5 21" />
          </svg>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="feather feather-paperclip"
          >
            <path
              d="M21.44 11.05l-9.19 9.19a6 6 0 01-8.49-8.49l9.19-9.19a4 4 0 015.66 5.66l-9.2 9.19a2 2 0 01-2.83-2.83l8.49-8.48"
            />
          </svg>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="feather feather-smile"
          >
            <circle cx="12" cy="12" r="10" />
            <path d="M8 14s1.5 2 4 2 4-2 4-2M9 9h.01M15 9h.01" />
          </svg>
          <input
            type="text"
            v-model.trim="chatText"
            @keyup.enter="doSendMsg"
            placeholder="Type something here..."
          />
          <button class="chat-send-btn" @click="doSendMsg">发送</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Manrope:300,400,500,600,700&display=swap&subset=latin-ext');

:root {
  --body-bg-color: #e5ecef;
  --theme-bg-color: #fff;
  --settings-icon-hover: #9fa7ac;
  --developer-color: #f9fafb;
  --input-bg: #f8f8fa;
  --input-chat-color: #a2a2a2;
  --border-color: #eef2f4;
  --body-font: 'Manrope', sans-serif;
  --body-color: #273346;
  --settings-icon-color: #c1c7cd;
  --msg-message: #969eaa;
  --chat-chatText-bg: #f1f2f6;
  --theme-color: #0086ff;
  --msg-date: #c0c7d2;
  --button-bg-color: #f0f7ff;
  --button-color: var(--theme-color);
  --detail-font-color: #919ca2;
  --msg-hover-bg: rgba(238, 242, 244, 0.4);
  --active-conversation-bg: linear-gradient(
    to right,
    rgba(238, 242, 244, 0.4) 0%,
    rgba(238, 242, 244, 0) 100%
  );
  --overlay-bg: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 0) 0%,
    rgba(255, 255, 255, 1) 65%,
    rgba(255, 255, 255, 1) 100%
  );
  --chat-header-bg: linear-gradient(
    to bottom,
    rgba(255, 255, 255, 1) 0%,
    rgba(255, 255, 255, 1) 78%,
    rgba(255, 255, 255, 0) 100%
  );
  --scrollbar-thumb: linear-gradient(to top, rgba(131, 164, 212, 0.5), rgb(182, 188, 255, 0.4));
}

::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  border-radius: 10px;
}

::-webkit-scrollbar-thumb {
  border-radius: 10px;
  background: var(--scrollbar-thumb);
}

[data-theme='purple'] {
  --theme-color: #9f7aea;
  --button-color: #9f7aea;
  --button-bg-color: rgba(159, 122, 234, 0.12);
}

[data-theme='green'] {
  --theme-color: #38b2ac;
  --button-color: #38b2ac;
  --button-bg-color: rgba(56, 178, 171, 0.15);
}

[data-theme='orange'] {
  --theme-color: #ed8936;
  --button-color: #ed8936;
  --button-bg-color: rgba(237, 137, 54, 0.12);
}

.dark-mode {
  --body-bg-color: #1d1d1d;
  --theme-bg-color: #27292d;
  --border-color: #323336;
  --body-color: #d1d1d2;
  --active-conversation-bg: linear-gradient(
    to right,
    rgba(47, 50, 56, 0.54),
    rgba(238, 242, 244, 0) 100%
  );
  --msg-hover-bg: rgba(47, 50, 56, 0.54);
  --chat-text-bg: #383b40;
  --chat-text-color: #b5b7ba;
  --msg-date: #626466;
  --msg-message: var(--msg-date);
  --overlay-bg: linear-gradient(to bottom, rgba(0, 0, 0, 0) 0%, #27292d 65%, #27292d 100%);
  --input-bg: #2f3236;
  --chat-header-bg: linear-gradient(
    to bottom,
    #27292d 0%,
    #27292d 78%,
    rgba(255, 255, 255, 0) 100%
  );
  --settings-icon-color: #7c7e80;
  --developer-color: var(--border-color);
  --button-bg-color: #393b40;
  --button-color: var(--body-color);
  --input-chat-color: #6f7073;
  --detail-font-color: var(--input-chat-color);
}

.blue {
  background-color: #0086ff;
}

.purple {
  background-color: #9f7aea;
}

.green {
  background-color: #38b2ac;
}

.orange {
  background-color: #ed8936;
}

* {
  outline: none;
  box-sizing: border-box;
}

img {
  max-width: 100%;
}

body {
  background-color: var(--body-bg-color);
  font-family: var(--body-font);
  color: var(--body-color);
}

html {
  box-sizing: border-box;
  -webkit-font-smoothing: antialiased;
}

.app {
  display: flex;
  flex-direction: column;
  background-color: var(--theme-bg-color);
  width: 1000px;
  // height: calc(100vh - 100px);
  height: 90vh;
  margin: 0 auto;
  overflow: hidden;
  border-radius: 5px;
}

.header {
  height: 80px;
  width: 100%;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.wrapper {
  width: 100%;
  display: flex;
  flex-grow: 1;
  overflow: hidden;
}

.conversation-area,
.detail-area {
  width: 340px;
  flex-shrink: 0;
}

.detail-area {
  border-left: 1px solid var(--border-color);
  margin-left: auto;
  padding: 30px 30px 0 30px;
  display: flex;
  flex-direction: column;
  overflow: auto;
}

.search-bar {
  height: 80px;
  z-index: 3;
  position: relative;
  margin-left: 0px;

  input {
    height: 100%;
    width: 100%;
    display: block;
    background-color: transparent;
    border: none;
    color: var(--body-color);
    padding: 0 54px;
    background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 56.966 56.966' fill='%23c1c7cd'%3e%3cpath d='M55.146 51.887L41.588 37.786A22.926 22.926 0 0046.984 23c0-12.682-10.318-23-23-23s-23 10.318-23 23 10.318 23 23 23c4.761 0 9.298-1.436 13.177-4.162l13.661 14.208c.571.593 1.339.92 2.162.92.779 0 1.518-.297 2.079-.837a3.004 3.004 0 00.083-4.242zM23.984 6c9.374 0 17 7.626 17 17s-7.626 17-17 17-17-7.626-17-17 7.626-17 17-17z'/%3e%3c/svg%3e");
    background-repeat: no-repeat;
    background-size: 16px;
    background-position: 25px 48%;
    font-family: var(--body-font);
    font-weight: 600;
    font-size: 15px;

    &::placeholder {
      color: var(--input-chat-color);
    }
  }
}

.user-settings {
  display: flex;
  align-items: center;
  cursor: pointer;
  margin-left: auto;
  flex-shrink: 0;

  & > * + * {
    margin-left: 14px;
  }
}

.dark-light {
  width: 22px;
  height: 22px;
  color: var(--settings-icon-color);
  flex-shrink: 0;

  svg {
    width: 100%;
    fill: transparent;
    transition: 0.5s;
  }
}

.user-profile {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.settings {
  color: var(--settings-icon-color);
  width: 22px;
  height: 22px;
  flex-shrink: 0;
}

.logout {
  color: var(--settings-icon-color);
  width: 22px;
  height: 22px;
  flex-shrink: 0;
}

.profile-name {
  margin-bottom: 4px;
  font-weight: 600;
  font-size: 25px;
}

.conversation-area {
  border-right: 1px solid var(--border-color);
  overflow-y: auto;
  overflow-x: hidden;
  display: flex;
  flex-direction: column;
  position: relative;
}

.msg-profile {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 15px;

  &.group {
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: var(--border-color);

    svg {
      width: 60%;
    }
  }
}

.msg {
  display: flex;
  align-items: center;
  padding: 20px;
  cursor: pointer;
  transition: 0.2s;
  position: relative;

  &:hover {
    background-color: var(--msg-hover-bg);
  }

  &.active {
    background: var(--active-conversation-bg);
    border-left: 4px solid var(--theme-color);
  }

  &.online:before {
    content: '';
    position: absolute;
    background-color: #23be7e;
    width: 9px;
    height: 9px;
    border-radius: 50%;
    border: 2px solid var(--theme-bg-color);
    left: 50px;
    bottom: 19px;
  }
}

.msg-detail {
  overflow: hidden;
}

.msg-username {
  margin-bottom: 4px;
  font-weight: 600;
  font-size: 15px;
}

.c-number {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background-color: #c4142c;
  color: #fff;
  font-weight: 500;
  font-size: 12px;
  width: 15px;
  height: 15px;
  border-radius: 50%;
  margin-left: 15px;
}

.msg-content {
  font-weight: 500;
  font-size: 13px;
  display: flex;
}

.msg-message {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  color: var(--msg-message);
}

.msg-date {
  font-size: 14px;
  color: var(--msg-date);
  margin-left: 3px;

  &:before {
    content: '•';
    margin-right: 2px;
  }
}

.add {
  position: sticky;
  bottom: 25px;
  background-color: var(--theme-color);
  width: 40px;
  height: 40px;
  border: 0;
  border-radius: 50%;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' width='24' height='24' fill='none' stroke='white' stroke-width='2' stroke-linecap='round' stroke-linejoin='round' class='feather feather-plus'%3e%3cpath d='M12 5v14M5 12h14'/%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: 50%;
  background-size: 28px;
  box-shadow: 0 0 12px var(--theme-color);
  margin: auto auto -55px;
  flex-shrink: 0;
  z-index: 1;
  cursor: pointer;
}

.overlay {
  position: sticky;
  bottom: 0;
  left: 0;
  width: 340px;
  flex-shrink: 0;
  background: var(--overlay-bg);
  height: 80px;
}

.chat-area {
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  overflow: auto;
  min-width: 500px;

  &-header {
    display: flex;
    position: sticky;
    top: 0;
    left: 0;
    z-index: 2;
    width: 100%;
    align-items: center;
    justify-content: space-between;
    padding: 20px;
    background: var(--chat-header-bg);
  }

  &-profile {
    width: 32px;
    border-radius: 50%;
    object-fit: cover;
  }

  &-title {
    font-size: 18px;
    font-weight: 600;
  }

  &-main {
    flex-grow: 1;
  }
}

.chat-msg-img {
  height: 40px;
  width: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.chat-msg-profile {
  flex-shrink: 0;
  margin-top: auto;
  margin-bottom: -20px;
  position: relative;
}

.chat-msg-date {
  position: absolute;
  left: calc(100% + 12px);
  bottom: 0;
  font-size: 12px;
  font-weight: 300;
  color: var(--msg-date);
  white-space: nowrap;
}

.chat-msg {
  display: flex;
  padding: 0 20px 45px;

  &-content {
    margin-left: 12px;
    max-width: 70%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  &-text {
    background-color: #eeeef1;
    padding: 15px;
    border-radius: 20px 20px 20px 0;
    line-height: 1.5;
    font-size: 14px;
    font-weight: 500;

    & + & {
      margin-top: 10px;
    }
  }
}

.chat-msg-text {
  color: var(--chat-text-color);
}

.owner {
  flex-direction: row-reverse;

  .chat-msg-content {
    margin-left: 0;
    margin-right: 12px;
    align-items: flex-end;
  }

  .chat-msg-text {
    background-color: var(--theme-color);
    color: #fff;
    border-radius: 20px 20px 0 20px;
  }

  .chat-msg-date {
    left: auto;
    right: calc(100% + 12px);
  }
}

.chat-msg-text img {
  max-width: 300px;
  width: 100%;
}

.chat-area-footer {
  display: flex;
  border-top: 1px solid var(--border-color);
  width: 100%;
  padding: 10px 20px;
  align-items: center;
  background-color: var(--theme-bg-color);
  position: sticky;
  bottom: 0;
  left: 0;
}

.chat-area-footer svg {
  color: var(--settings-icon-color);
  width: 20px;
  flex-shrink: 0;
  cursor: pointer;

  &:hover {
    color: var(--settings-icon-hover);
  }

  & + svg {
    margin-left: 12px;
  }
}

.chat-area-footer input {
  border: none;
  color: var(--body-color);
  background-color: var(--input-bg);
  padding: 14px;
  border-radius: 6px;
  font-size: 15px;
  margin: 0 12px;
  width: 100%;

  &::placeholder {
    color: var(--input-chat-color);
  }
}

.chat-area-footer button {
  height: 100%;
  width: 80px;
  color: #fff;
  text-justify: center;
  background-color: var(--theme-color);
  border: none;
  border-radius: 4px;
  padding: 0 10px;
  font-size: 14px;
}

.developer {
  position: absolute;
  color: var(--detail-font-color);
  font-weight: 600;
  left: 0;
  top: -100%;
  display: flex;
  transition: 0.3s;
  padding: 0 20px;
  align-items: center;
  justify-content: center;
  background-color: var(--developer-color);
  width: 100%;
  height: 100%;
}

.developer img {
  border-radius: 50%;
  width: 26px;
  height: 26px;
  object-fit: cover;
  margin-right: 10px;
}

.dark-mode {
  .search-bar input,
  .detail-area input {
    background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 56.966 56.966' fill='%236f7073'%3e%3cpath d='M55.146 51.887L41.588 37.786A22.926 22.926 0 0046.984 23c0-12.682-10.318-23-23-23s-23 10.318-23 23 10.318 23 23 23c4.761 0 9.298-1.436 13.177-4.162l13.661 14.208c.571.593 1.339.92 2.162.92.779 0 1.518-.297 2.079-.837a3.004 3.004 0 00.083-4.242zM23.984 6c9.374 0 17 7.626 17 17s-7.626 17-17 17-17-7.626-17-17 7.626-17 17-17z'/%3e%3c/svg%3e");
  }

  .dark-light svg {
    fill: #ffce45;
    stroke: #ffce45;
  }

  .chat-area-group span {
    color: #d1d1d2;
  }
}

.chat-area-group {
  flex-shrink: 0;
  display: flex;

  * {
    border: 2px solid var(--theme-bg-color);
  }

  * + * {
    margin-left: -5px;
  }

  span {
    width: 32px;
    height: 32px;
    background-color: var(--button-bg-color);
    color: var(--theme-color);
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 14px;
    font-weight: 500;
  }
}

@media (max-width: 1120px) {
  .detail-area {
    display: none;
  }
}

@media (max-width: 780px) {
  .conversation-area {
    display: none;
  }

  .search-bar {
    margin-left: 0;
    flex-grow: 1;

    input {
      padding-right: 10px;
    }
  }
}
</style>
