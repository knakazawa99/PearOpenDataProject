import axios from 'axios';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { LoadingButton } from '@mui/lab';
import { Alert, Box, Typography } from '@mui/material';

import { useState } from 'react';
import * as React from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';

import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { AdminAuth} from 'components/features/admin_pear/Type';
import { BASE_URL } from 'config/config';
import { AlertColor } from '@mui/material/Alert/Alert';
import { AdminAuthFormValues } from './Type';

const Delete = (prop: {adminAuth: AdminAuth, updateFunc: (adminAuth: AdminAuth) => void}) => {
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [isAlertMessage, setIsAlertMessage] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>()
  const [alertMessageSeverity, setAlertMessageSeverity] = useState<AlertColor>("success")
  const [open, setOpen] = useState(false)

  const { handleSubmit } = useForm<AdminAuthFormValues>({})

  const handleClickOpen = () => {
    setOpen(true);
  }

  const handleClose = () => {
    setOpen(false);
  }

  const onSubmitInner: SubmitHandler<FieldValues> = async (
    submitData,
    event,
  ) => {
    event?.preventDefault();
    setIsLoading(true)
    const path = BASE_URL + "/v1/admin/admins/" + prop.adminAuth.id
    await axios.delete(
      path,
      {
        headers: {
          'authorization': `Bearer ${localStorage.getItem("jwtToken")}`,
          'x-jwtKey': `${localStorage.getItem("jwtKey")}`,
        }
      }
    ).then((res) => {

      prop.updateFunc({
        id: prop.adminAuth.id,
        email: "",
        createdAt: new Date(),
      })
      unblocking();
      setIsLoading(false)
      handleClose()
    }).catch(() => {
      setIsAlertMessage(true)
      setAlertMessage("削除に失敗しました。")
      setAlertMessageSeverity("error")
      unblocking()
      setIsLoading(false)
    })
  }

  const [onSubmit, processing, unblocking] = useBlockDoubleClick(
    onSubmitInner,
  )

  return (
    <div>
      <Button variant="outlined" onClick={handleClickOpen}>
        削除
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>削除({prop.adminAuth.id})</DialogTitle>
        <DialogContent>
          <Typography variant="subtitle1" align="left" color="text.secondary">
            本当に管理者を削除しますか？
          </Typography>
          <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>
            <LoadingButton
              loading={isLoading}
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              削除
            </LoadingButton>

            {
              isAlertMessage ? <Alert variant="outlined" onClose={() => {setIsAlertMessage(false)}} severity={alertMessageSeverity}>{alertMessage}</Alert>
                : <div/>
            }
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>閉じる</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default Delete
