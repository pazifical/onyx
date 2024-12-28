import type { Note } from '@/types'

export class NoteRepository {
  async getAllNotes(): Promise<Array<Note>> {
    const response = await fetch('/api/notes')
    return response.json()
  }
}
