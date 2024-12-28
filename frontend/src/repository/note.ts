import type { Note } from '@/types'

export class NoteRepository {
  async getAll(): Promise<Array<Note>> {
    const response = await fetch('/api/notes')
    return response.json()
  }

  async getByID(id: string): Promise<Note> {
    const response = await fetch(`/api/notes/${id}`)
    return response.json()
  }

  async getByPath(path: string): Promise<Note> {
    const url = `/api/notes/${path}`
    const response = await fetch(url)
    return response.json()
  }
}
