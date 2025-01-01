import type { Reminder } from '@/types'

export class ReminderRepository {
  async getAll(): Promise<Array<Reminder>> {
    const url = `/api/reminders`
    const response = await fetch(url)
    return response.json()
  }
}
