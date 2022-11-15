import { http } from './http'
import { ResponseData } from './common'

// uuid: group uuid
async function listGroupUsers(uuid: string) {
  return await http.get<ResponseData>(`/groups/${uuid}/users`)
}

export { listGroupUsers }
