import { Table, TableBody, TableCell, TableHead, TableRow } from "@mui/material"
import React, { useEffect, useState } from 'react';
import { BASE_URL } from 'config/config';
import axios from 'axios';
import EditFormDialog from './EditForm';
import { AdminVersion } from 'components/features/admin_pear/Type';

type APIAdminVersionInformation = {
  id: number
  file_path: string
  version: string
  release_note: string
  release_comment: string
  release_flag: boolean
  created_at: string
}

const AdminPear = () => {
  const [adminVersionInformation, setAdminVersionInformation] = useState<AdminVersion[]>()

  useEffect(() => {
    const fetchData = async () => {
      const path = BASE_URL + "/v1/admin/versions";
      await axios.get(path).then((response) => {
        const adminVersionInformation = response.data?.map((info: APIAdminVersionInformation) => {
          return {
            id: info.id,
            filePath: info.file_path,
            version: info.version,
            releaseNote: info.release_note,
            releaseComment: info.release_comment,
            releaseFlag: info.release_flag,
            createdAt: new Date(info.created_at)
          }
        })
        setAdminVersionInformation(adminVersionInformation)
      })
    }
    fetchData()
  }, [])

  const updateAdminVersionInformation = (adminVersion: AdminVersion) => {
    if (!adminVersionInformation) {
      return
    }
    const afterAdminVersionInformation = [...adminVersionInformation]
    for (let i =0; i<adminVersionInformation?.length; i++) {
      if (adminVersion.id == adminVersionInformation[i].id) {
        afterAdminVersionInformation[i].releaseComment = adminVersion.releaseComment
        afterAdminVersionInformation[i].releaseNote = adminVersion.releaseNote
        afterAdminVersionInformation[i].releaseFlag = adminVersion.releaseFlag
        break
      }
    }
    setAdminVersionInformation(afterAdminVersionInformation)
  }

  return <div>
    <Table size="small">
      <TableHead>
        <TableRow>
          <TableCell>バージョン</TableCell>
          <TableCell>ファイルパス</TableCell>
          <TableCell>リリースノート</TableCell>
          <TableCell>リリースコメント</TableCell>
          <TableCell>リリースフラグ</TableCell>
          <TableCell align="right">リリース日</TableCell>
          <TableCell align="right">編集</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>

        {adminVersionInformation?.map((adminVersion) => (
          <TableRow key={adminVersion.id}>
            <TableCell>{adminVersion.version}</TableCell>
            <TableCell>{adminVersion.filePath}</TableCell>
            <TableCell>{adminVersion.releaseNote}</TableCell>
            <TableCell>{adminVersion.releaseComment}</TableCell>
            <TableCell>{adminVersion.releaseFlag? "ON": "OFF"}</TableCell>
            <TableCell align="right">{adminVersion.createdAt.getFullYear() + "." + adminVersion.createdAt.getMonth() + "." + adminVersion.createdAt.getDate()}</TableCell>
            <TableCell　align="right">
              <EditFormDialog key={adminVersion.id}  adminVersion={adminVersion} updateFunc={updateAdminVersionInformation} />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </div>
}

export default AdminPear
