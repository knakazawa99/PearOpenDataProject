import axios from 'axios';
import Avatar from '@mui/material/Avatar';
import { LoadingButton } from '@mui/lab';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import { AlertColor } from '@mui/material/Alert/Alert';
import { Alert } from '@mui/material';
import * as React from 'react';
import { useRef, useState } from 'react';

import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';
import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { EmailValidPattern} from 'common/emailValidPattern';
import { BASE_URL } from 'config/config';
import router from 'Rotuer';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';


type AdminAuthFormValues = {
  email: string
  password: string
}

const AdminAuth = () => {

  const [isLoading, setIsLoading] = useState<boolean>(false)

  const emailRef = useRef<HTMLInputElement>(null);
  const [emailError, setEmailError] = useState(false)
  const [emailErrorMessage, setEmailErrorMessage] = useState<string>()

  const passwordRef = useRef<HTMLInputElement>(null);
  const [passwordError, setPasswordError] = useState(false)
  const [passwordErrorMessage, setPasswordErrorMessage] = useState<string>()

  const [isAlertMessage, setIsAlertMessage] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>()
  const [alertMessageSeverity, setAlertMessageSeverity] = useState<AlertColor>("success")

  const { register, handleSubmit } = useForm<AdminAuthFormValues>({});

  const onSubmitInner: SubmitHandler<FieldValues> = async (
    submitData,
    event,
  ) => {
    setIsLoading(true)
    event?.preventDefault();

    if (!formValidation()) {
      setIsLoading(false)
      unblocking();
      return

    }

    const path = BASE_URL + "/v1/admin/signup"
    await axios.post(
      path,
      submitData
    ).then((res) => {
      localStorage.setItem("jwtToken", res.data.token)
      localStorage.setItem("jwtKey", submitData.email)
      unblocking();
      setIsLoading(false)
      router.navigate("/admin/versions")
    }).catch(() => {
      unblocking()
      setAlertMessageSeverity("error")
      setIsAlertMessage(true)
      setAlertMessage("サインアップに失敗しました。")
      setIsLoading(false)
    })
  }

  const [onSubmit, processing, unblocking] = useBlockDoubleClick(
    onSubmitInner,
  )

  const formValidation = (): boolean => {
    let valid = true

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

  return <div>
    <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
      <LockOutlinedIcon />
    </Avatar>
    <Typography component="h1" variant="h5">
      サインアップ
    </Typography>
    <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>

      <TextField
        margin="normal"
        required
        fullWidth
        id="email"
        label="メールアドレス"
        autoComplete="email"
        {...register('email')}
        inputRef={emailRef}
        inputProps={ {required: true, pattern: EmailValidPattern} }
        helperText={emailError ? emailErrorMessage: ""}
      />

      <TextField
        margin="normal"
        required
        fullWidth
        label="パスワード"
        type="password"
        id="password"
        autoComplete="current-password"
        {...register('password')}
        inputRef={passwordRef}
        helperText={passwordError ? passwordErrorMessage: ""}
      />

      <LoadingButton
        loading={isLoading}
        type="submit"
        fullWidth
        variant="contained"
        sx={{ mt: 3, mb: 2 }}
      >
        Sign In
      </LoadingButton>

      {
        isAlertMessage ? <Alert variant="outlined" onClose={() => {setIsAlertMessage(false)}} severity={alertMessageSeverity}>{alertMessage}</Alert>
          : <div/>
      }
    </Box>
  </div>
}

export default AdminAuth
