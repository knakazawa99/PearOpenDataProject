import axios from 'axios';
import Box from '@mui/material/Box'
import { LoadingButton } from '@mui/lab';
import { Alert, Button} from '@mui/material';
import { AlertColor } from '@mui/material/Alert/Alert';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import React, { useRef, useState } from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';

import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { EmailValidPattern } from 'common/emailValidPattern';
import { AdminAuthFormValues } from 'components/features/admin/Type';
import { AdminAuth } from 'components/features/admin_pear/Type';
import { BASE_URL } from 'config/config';


const Add = (prop: {updateFunc: (adminAuth: AdminAuth) => void}) => {

  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [isAlertMessage, setIsAlertMessage] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>()
  const [alertMessageSeverity, setAlertMessageSeverity] = useState<AlertColor>("success")
  const [open, setOpen] = useState(false)

  const emailRef = useRef<HTMLInputElement>(null)
  const [emailError, setEmailError] = useState(false)
  const [emailErrorMessage, setEmailErrorMessage] = useState<string>()

  const passwordRef = useRef<HTMLInputElement>(null)
  const [passwordError, setPasswordError] = useState(false)
  const [passwordErrorMessage, setPasswordErrorMessage] = useState<string>()


  const { register, handleSubmit } = useForm<AdminAuthFormValues>({})

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
    if (!formValidation()) {
      unblocking();
      return
    }
    event?.preventDefault();
    setIsLoading(true)
    const path = BASE_URL + "/v1/admin/admins/"
    await axios.post(
      path,
      submitData,
      {
        headers: {
          'authorization': `Bearer ${localStorage.getItem("jwtToken")}`,
          'x-jwtKey': `${localStorage.getItem("jwtKey")}`,
        }
      }
    ).then((res) => {
      prop.updateFunc({
        id: res.data.id,
        email: res.data.email,
        createdAt: new Date(res.data.createdAt)
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

  const formValidation = (): boolean => {
    let valid = true;

    const email = emailRef?.current;
    if (email) {
      const ok = email.validity.valid;
      setEmailError(!ok);
      valid &&= ok;
      if (!ok) {
        setEmailErrorMessage("メールアドレスを正しい形式で入力してください。")
      } else {
        setEmailErrorMessage("")
      }
    }

    const password = passwordRef?.current;
    if (password) {
      const ok = password.validity.valid;
      setPasswordError(!ok);
      valid &&= ok;
      if (!ok) {
        setPasswordErrorMessage("パスワードを入力してください。")
      } else {
        setPasswordErrorMessage("")
      }
    }
    return valid
  }
  // TODO: https://chusotsu-program.com/js-base64-encode/
  return (

    <div>
      <Button  variant="outlined" onClick={handleClickOpen}>
        管理者を追加
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>管理者を追加</DialogTitle>
        <DialogContent>
          <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              type="email"
              id="email"
              label="メールアドレス"
              autoComplete="email"
              variant="outlined"
              {...register('email')}
              inputRef={emailRef}
              inputProps={ {required: true, pattern: EmailValidPattern} }
              helperText={emailError ? emailErrorMessage: ""}
            />
            <TextField
              margin="normal"
              required
              label="パスワード"
              fullWidth
              variant="outlined"
              {...register("password")}
              inputRef={passwordRef}
              inputProps={ {required: true}}
              helperText={ passwordError? passwordErrorMessage: ""}
            />

            <LoadingButton
              loading={isLoading}
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              追加
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

export default Add