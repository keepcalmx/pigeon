import { http } from './http'
import { ResponseData } from './common'

async function getUserInfo(uuid: string) {
  return await http.get<ResponseData>(`/users/${uuid}`)
}

async function createUser(data: any) {
  return await http.post<ResponseData>('/users', data)
}

async function getUserContactList(uuid: string) {
  return await http.get<ResponseData>(`/users/${uuid}/contacts`)
}

export { createUser, getUserInfo, getUserContactList }
