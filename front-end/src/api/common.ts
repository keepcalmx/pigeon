const baseURL = import.meta.env.PROD ? import.meta.env.VITE_PROD_URL : import.meta.env.VITE_DEV_URL

interface ResponseData {
  code: string
  msg: string
  data: any
}

export { baseURL }
export type { ResponseData }
