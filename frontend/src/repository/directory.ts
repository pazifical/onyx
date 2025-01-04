import type { DirectoryContent } from '@/types'

export class DirectoryContentRepository {
  async getByPath(path: string): Promise<DirectoryContent> {
    const url = `/api/directory${path}`
    console.log("get", url)
    const response = await fetch(url)

    const data = await response.json()
    if (response.status > 400) {
      throw new Error(data.error_message)
    }

    return data
  }

  async createNew(path: string): Promise<DirectoryContent> {
    const url = `/api/directory${path}`
    const response = await fetch(url,{
      method: "POST",
    })

    const data = await response.json()
    if (response.status > 400) {
      throw new Error(data.error_message)
    }

    return data
  }
}
