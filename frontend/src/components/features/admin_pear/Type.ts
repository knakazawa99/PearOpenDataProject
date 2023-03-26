export type AdminVersion = {
  id: number
  version: string
  filePath: string
  releaseNote: string
  releaseComment: string
  releaseFlag: boolean
  createdAt: Date
}

export type AdminVersionFormValues = {
  files: []
  id: number
  version: string
  file_path: string
  release_note: string
  release_comment: string
  release_flag: boolean
  created_at: Date
}
