import axios from 'axios';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { LoadingButton } from '@mui/lab';
import { Alert, Box, Checkbox, FormControlLabel } from '@mui/material';
import TextField from '@mui/material/TextField';

import { useState } from 'react';
import * as React from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';

import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { AdminVersion, AdminVersionFormValues } from 'components/features/admin_pear/Type';
import { BASE_URL } from 'config/config';
import { AlertColor } from '@mui/material/Alert/Alert';

const EditFormDialog = (prop: {adminVersion: AdminVersion, updateFunc: (adminVersion: AdminVersion) => void}) => {
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [isAlertMessage, setIsAlertMessage] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>()
  const [alertMessageSeverity, setAlertMessageSeverity] = useState<AlertColor>("success")
  const [open, setOpen] = useState(false)

  const { register, handleSubmit } = useForm<AdminVersionFormValues>({})

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
    const path = BASE_URL + "/v1/admin/versions/" + prop.adminVersion.id
    await axios.put(
      path,
      submitData
    ).then((res) => {

      prop.updateFunc({
        id: prop.adminVersion.id,
        version: prop.adminVersion.version,
        filePath: prop.adminVersion.filePath,
        releaseNote: submitData.releaseNote,
        releaseComment: submitData.releaseComment,
        releaseFlag: submitData.releaseFlag,
        createdAt: new Date()
      })
      unblocking();
      setIsLoading(false)
      handleClose()
    }).catch(() => {
      setIsAlertMessage(true)
      setAlertMessage("更新に失敗しました。")
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
        編集
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>編集({prop.adminVersion.id})</DialogTitle>
        <DialogContent>
          <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>
            <TextField
              autoFocus
              disabled
              margin="dense"
              label="ファイルパス"
              fullWidth
              variant="standard"
              value={prop.adminVersion.filePath}
            />
            <TextField
              autoFocus
              margin="dense"
              label="リリースノート"
              fullWidth
              variant="standard"
              defaultValue={prop.adminVersion.releaseNote}
              {...register("release_note")}
            />
            <TextField
              autoFocus
              margin="dense"
              label="リリースコメント"
              fullWidth
              variant="standard"
              defaultValue={prop.adminVersion.releaseComment}
              {...register("release_comment")}
            />
            <FormControlLabel control={<Checkbox defaultChecked={prop.adminVersion.releaseFlag} />} label="リリースフラグ" {...register("release_flag")} />
            <LoadingButton
              loading={isLoading}
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              更新
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

export default EditFormDialog
