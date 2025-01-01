import type { DirectoryContent } from '@/types'

export class DirectoryContentRepository {
  async getByPath(path: string): Promise<DirectoryContent> {
    const url = `/api/directory${path}`
    console.log("get", url)
    const response = await fetch(url)
    return response.json()
  }

  async createNew(path: string): Promise<DirectoryContent> {
    const url = `/api/directory${path}`
    const response = await fetch(url,{
      method: "POST",
    })
    return response.json()
  }
}
