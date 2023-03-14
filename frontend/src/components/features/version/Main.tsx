import axios from 'axios';
import  { saveAs } from 'file-saver';
import { LoadingButton } from '@mui/lab';
import {
  Alert,
  Box,
  Button, Checkbox,
  Container,
  createTheme,
  CssBaseline,
  FormControl, FormControlLabel,
  FormGroup,
  FormHelperText,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  ThemeProvider,
  Typography
} from '@mui/material';
import React, { useEffect, useRef, useState } from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';

import { getFileNameFromContentDisposition } from 'common/file';
import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { BASE_URL } from 'config/config';
import VersionItem from './VersionItem';
import Grid from '@mui/material/Grid';
import { AlertColor } from '@mui/material/Alert/Alert';
import ScrollPlayground from './Term';
import Term from './Term';

type PearInformation = {
  version: string
  releaseNote: string
  createdAt: Date
}

type APIPearInformation = {
  version: string
  release_note: string
  created_at: string
}

type PearDownloadFormValues = {
  agreement: boolean
  organization: string
  name: string
  email: string
  version: string
  token: string
}

const theme = createTheme();
const EmailValidPattern = "^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:.[a-zA-Z0-9-]+)*$";

const Version = () => {
  const [pearInformation, setPearInformation] = useState<PearInformation[]>()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [downloadStage, setDownloadStage] = useState<number>(1)

  const [agreement, setAgreement] = useState(false)

  const organizationRef = useRef<HTMLInputElement>(null);
  const [organizationError, setOrganizationError] = useState(false);
  const [organizationErrorMessage, setOrganizationErrorMessage] = useState<string>()


  const nameRef = useRef<HTMLInputElement>(null);
  const [nameError, setNameError] = useState(false);
  const [nameErrorMessage, setNameErrorMessage] = useState<string>()


  const emailRef = useRef<HTMLInputElement>(null);
  const [emailError, setEmailError] = useState(false);
  const [emailErrorMessage, setEmailErrorMessage] = useState<string>()
  const versionRef = useRef<HTMLInputElement>(null);

  const [versionErrorMessage, setVersionErrorMessage] = useState<string>()

  const tokenRef = useRef<HTMLInputElement>(null);
  const [tokenError, setTokenError] = useState(false);
  const [tokenErrorMessage, setTokenErrorMessage] = useState<string>()

  const [isAlertMessage, setIsAlertMessage] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>()
  const [alertMessageSeverity, setAlertMessageSeverity] = useState<AlertColor>("success")

  const { register, handleSubmit } = useForm<PearDownloadFormValues>({});
  useEffect(() => {
    const fetchData = async () => {
      const path = BASE_URL + "/v1/pears/";
      await axios.get(path).then((response) => {
        const pearInformation = response.data?.map((info: APIPearInformation) => {
          return {
            version: info.version,
            releaseNote: info.release_note,
            createdAt: new Date(info.created_at)
          }
        })
        setPearInformation(pearInformation)
      })
    }
    fetchData()
  }, [])

  const onSubmitInner: SubmitHandler<FieldValues> = async (
    submitData,
    event,
  ) => {
    if (!formValidation()) {
      unblocking();
      return
    }
    setIsLoading(true)
    event?.preventDefault();
    let path = BASE_URL
    if (downloadStage == 1) {
      path += "/v1/auth/notify/request"
      await axios.post(
        path,
        submitData
      ).then(() => {
        unblocking();
        setDownloadStage(downloadStage ? downloadStage + 1 : 1)
        setIsLoading(false)
        setAlertMessageSeverity("success")
        setIsAlertMessage(true)
        setAlertMessage("メールを送信しました！")
      }).catch(() => {
        setIsLoading(false)
        setAlertMessageSeverity("error")
        setIsAlertMessage(true)
        setAlertMessage("メールを送信に失敗しました。")
      })
    } else if (downloadStage == 2) {
      path += "/v1/auth/download?"
      const keys = Object.keys(submitData);
      for (let i = 0; i<keys.length; i++) {
        path += keys[i] + "=" + submitData[keys[i]] + "&"
      }
      await axios.get(
        path,{
          responseType: 'arraybuffer',
          headers: { Accept: 'application/zip' },
        }
      ).then((response) => {
        unblocking();
        const blob = new Blob([response.data], {
          type: response.data.type,
        });
        const contentDisposition = response.headers["content-disposition"];
        const fileName = getFileNameFromContentDisposition(contentDisposition)
        saveAs(blob, fileName);
        setDownloadStage(1)
        setIsLoading(false)
        setIsAlertMessage(true)
        setAlertMessageSeverity("success")
        setAlertMessage("ダウンロードに成功しました！")
      }).catch(() => {
        unblocking();
        setIsLoading(false)
        setAlertMessageSeverity("error")
        setIsAlertMessage(true)
        setAlertMessage("ダウンロードに失敗しました。")
      })
    }
  }

  const [onSubmit, processing, unblocking] = useBlockDoubleClick(
    onSubmitInner,
  );

  const formValidation = (): boolean => {
    let valid = true;

    if (downloadStage == 2) {
      const token = tokenRef?.current
      if (token) {
        const ok = token?.validity.valid;
        setTokenError(!ok)
        valid &&= ok;
        if (!ok) {
          setTokenErrorMessage("トークンを入力してください")
        } else {
          setTokenErrorMessage("")
        }
      }
      return valid
    }

    const organization = organizationRef?.current;
    if (organization) {
      const ok = organization?.validity.valid;
      setOrganizationError(!ok);
      valid &&= ok;
      if (!ok) {
        setOrganizationErrorMessage("組織名を入力してください")
      } else {
        setOrganizationErrorMessage("")
      }
    }

    const name = nameRef?.current;
    if (name) {
      const ok = name?.validity.valid;
      setNameError(!ok);
      valid &&= ok;
      if (!ok) {
        setNameErrorMessage("氏名を入力してください")
      } else {
        setNameErrorMessage("")
      }
    }

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

    const version = versionRef?.current
    if (version) {
      const ok = version?.validity.valid;
      valid &&= ok;
      if (!ok) {
        setVersionErrorMessage("取得したいバージョンを選択してください")
      } else {
        setVersionErrorMessage("")
      }
    } else {
      setVersionErrorMessage("取得したいバージョンを選択してください")
    }

    return valid
  }

  return <div>
    <Container maxWidth="md">
      <Typography
        component="h2"
        variant="h5"
        align="center"
        color="text.primary"
        gutterBottom
      >
        公開済みのバージョン一覧
      </Typography>
      <Grid container spacing={4}>
      {pearInformation?.map((info, index) => (
        <VersionItem version={info} key={index}/>
      ))}
      </Grid>
    </Container>
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="md">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Typography component="h1" variant="h4">
            データ取得リクエスト
          </Typography>
          {downloadStage == 1 &&
            <Typography paragraph={true}>
              取得したいデータのバージョンとメールアドレスを入力してください。<br/>
              入力して送信すると、メールアドレス宛にトークンが送信されます。
            </Typography>
          }
          {downloadStage == 2 &&
            <Typography paragraph={true}>
              メールアドレスに送信されたトークンを入力してください。
            </Typography>
          }
          <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>

            {downloadStage == 1 &&
              <div>

                <TextField
                  margin="normal"
                  required
                  fullWidth
                  id="organization"
                  label="組織名"
                  autoFocus
                  variant="outlined"
                  {...register('organization')}
                  inputRef={organizationRef}
                  helperText={organizationError ? organizationErrorMessage: ""}
                />

                <TextField
                  margin="normal"
                  required
                  fullWidth
                  id="name"
                  label="氏名"
                  autoFocus
                  variant="outlined"
                  {...register('name')}
                  inputRef={nameRef}
                  helperText={nameError ? nameErrorMessage: ""}
                />

                <TextField
                  margin="normal"
                  required
                  fullWidth
                  type="email"
                  id="email"
                  label="メールアドレス"
                  autoComplete="email"
                  autoFocus
                  variant="outlined"
                  {...register('email')}
                  inputRef={emailRef}
                  inputProps={ {required: true, pattern: EmailValidPattern} }
                  helperText={emailError ? emailErrorMessage: ""}
                />

                <FormControl
                  margin="normal"
                  required
                  fullWidth
                >
                  <FormGroup>
                    <InputLabel id="version-label">取得するバージョン</InputLabel>
                    <Select
                      required
                      labelId="version-label"
                      id="version"
                      label="取得するバージョン"
                      {...register('version')}
                      variant="outlined"
                      defaultValue=""
                    >
                      {pearInformation?.map((info, index) => (
                        <MenuItem value={info.version} key={index}>{info.version}</MenuItem>
                      ))}
                    </Select>
                    <FormHelperText>{versionErrorMessage}</FormHelperText>
                  </FormGroup>
                </FormControl>
              <Term/>

              <FormControlLabel control={<Checkbox checked={agreement} onChange={() => setAgreement(!agreement)} />} label="利用規約に同意する" {...register('agreement')} />
              </div>
            }

            {downloadStage == 2 &&
              <div>
                <div>
                  <TextField
                    margin="normal"
                    required
                    fullWidth
                    id="token"
                    label="トークン"
                    autoFocus
                    variant="outlined"
                    {...register('token')}
                    inputRef={tokenRef}
                    helperText={tokenError ? tokenErrorMessage: ""}
                  />
                </div>
                <Button
                  onClick={() => {setDownloadStage(1)}}
                  >メールを再送する
                </Button>
              </div>
            }

            <LoadingButton
              loading={isLoading}
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              disabled={downloadStage==1 && !agreement}
            >
              送信
            </LoadingButton>

            {
              isAlertMessage ? <Alert variant="outlined" onClose={() => {setIsAlertMessage(false)}} severity={alertMessageSeverity}>{alertMessage}</Alert>
                : <div/>
            }
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  </div>
}

export default Version