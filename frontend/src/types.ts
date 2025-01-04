export interface Note {
  path: string
  text: string
}

export interface DirectoryContent {
  files: Array<string>
  directories: Array<string>
}

export interface Reminder {
  date: string
  todo: string
  source: string
}

export interface OnyxError {
	error_message: string
}
