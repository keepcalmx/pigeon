import { http } from './http'

async function ping() {
  return await http.get<string>('/ping')
}

export { ping }
