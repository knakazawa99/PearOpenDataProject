import axios from 'axios';
import { Container, Paper, Table, TableBody, TableCell, TableHead, TableRow, Typography } from '@mui/material';
import Grid from '@mui/material/Grid';
import React, { useEffect, useState } from 'react';

import { AdminAuth } from 'components/features/admin_pear/Type';
import Add from 'components/features/admin/Add';
import { BASE_URL } from 'config/config';
import router from 'Rotuer';
import Delete from './Delete';
import admin from '../../pages/Admin';

type APIAdminAuth = {
  id: number
  email: string
  created_at: string
}

const Admin = () => {
  const [adminAuths, setAdminAuths] = useState<AdminAuth[]>([])

  useEffect(() => {
    const fetchData = async () => {
      const jwtKey = localStorage.getItem("jwtKey")
      const jwtToken = localStorage.getItem("jwtToken")
      if (jwtKey == "" || jwtToken == "") {
        await router.navigate("/admin/signin")
      }

      const path = BASE_URL + "/v1/admin/admins/";
      await axios.get(path, {headers: {
          'authorization': `Bearer ${jwtToken}`,
          'x-jwtKey': `${jwtKey}`,
        }}).then((response) => {
        const adminAuths = response.data?.map((info: APIAdminAuth) => {
          return {
            id: info.id,
            email: info.email,
            createdAt: new Date(info.created_at)
          }
        })
        setAdminAuths(adminAuths)
      })
    }
    fetchData()
  }, [])

  const deleteAdminAuth = (adminAuth: AdminAuth) => {
    if (!adminAuths) {
      return
    }
    const afterAdminAuths = adminAuths.filter((beforeAdminAuth) => {
      return beforeAdminAuth.id != adminAuth.id
    })
    setAdminAuths(afterAdminAuths)
  }

  const addAdminAuth = (adminAuth: AdminAuth) => {
    if (!adminAuths) {
      return
    }
    const afterAdminAuths = [...adminAuths]
    afterAdminAuths.push(adminAuth)
    setAdminAuths(afterAdminAuths)
  }

  return <div>
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Grid item xs={12}>
        <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
          <Typography align="right">
            <Add updateFunc={addAdminAuth}/>
          </Typography>
          <Typography
            component="h4"
            variant="h5"
            align="left"
            color="text.primary"
            gutterBottom
          >
            管理者一覧
          </Typography>
          <Table size="small">
            <TableHead>
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell>メールアドレス</TableCell>
                <TableCell align="right">作成日</TableCell>
                <TableCell align="right">削除</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {adminAuths?.map((adminAuth) => (
                <TableRow key={adminAuth.id}>
                  <TableCell>{adminAuth.id}</TableCell>
                  <TableCell>{adminAuth.email}</TableCell>
                  <TableCell align="right">{adminAuth.createdAt.getFullYear() + "." + adminAuth.createdAt.getMonth() + "." + adminAuth.createdAt.getDate()}</TableCell>
                  <TableCell　align="right">
                    <Delete key={adminAuth.id}  adminAuth={adminAuth} updateFunc={deleteAdminAuth} />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Paper>
      </Grid>
    </Container>
  </div>
}

export default Admin
