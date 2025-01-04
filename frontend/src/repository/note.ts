import type { Note } from '@/types'

export class NoteRepository {
  async getAll(): Promise<Array<Note>> {
    const response = await fetch('/api/notes')

    const data = await response.json()
    if (response.status > 400) {
      throw new Error(data.error_message)
    }

    return data
  }

  async getByPath(path: string): Promise<Note> {
    const url = `/api/notes/${path}`
    const response = await fetch(url)

    const data = await response.json()
    if (response.status > 400) {
      throw new Error(data.error_message)
    }

    return data
  }

  async saveNote(note: Note): Promise<Note> {
    const url = `/api/notes/${note.path}`
    const response = await fetch(url,{
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(note),
    })

    const data = await response.json()
    if (response.status > 400) {
      throw new Error(data.error_message)
    }

    return data
  }
}
