import { http } from './http'
import { ResponseData } from './common'

async function getToken(data: any) {
  return await http.post<ResponseData>('/login', data)
}

export { getToken }
