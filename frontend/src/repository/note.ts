import type { Note } from '@/types'

export class NoteRepository {
  async getAll(): Promise<Array<Note>> {
    const response = await fetch('/api/notes')
    return response.json()
  }

  async getByPath(path: string): Promise<Note> {
    const url = `/api/notes/${path}`
    const response = await fetch(url)
    return response.json()
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
    return response.json()
  }
}
