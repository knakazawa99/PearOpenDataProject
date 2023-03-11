import axios from 'axios';
import  { saveAs } from 'file-saver';
import {
  Box,
  Button,
  Container,
  createTheme,
  CssBaseline,
  FormControl,
  FormGroup,
  FormHelperText,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  ThemeProvider,
  Typography
} from '@mui/material';
import React, { useEffect, useState } from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';

import Loading from 'components/ui/Loading';
import { getFileNameFromContentDisposition } from 'common/file';
import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { BASE_URL } from 'config/config';
import VersionItem from './VersionItem';
import Grid from '@mui/material/Grid';

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
  version: string
  email: string
  token: string
}

const theme = createTheme();

const Version = () => {
  const [pearInformation, setPearInformation] = useState<PearInformation[]>()

  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [downloadStage, setDownloadStage] = useState<number>(1)


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
    setIsLoading(true)
    event?.preventDefault();
    let path = BASE_URL
    if (downloadStage == 1) {
      path += "/v1/auth/notify/request"
      console.log("submitData: ", submitData)
      await axios.post(
        path,
        submitData
      ).then(() => {
        unblocking();
      })
      setDownloadStage(downloadStage ? downloadStage + 1 : 1)
      setIsLoading(false)
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

      }).catch(() => {
        unblocking();
        setIsLoading(false)
      })
    }
  }


  const [onSubmit, processing, unblocking] = useBlockDoubleClick(
    onSubmitInner,
  );

  return <div>
    {isLoading ? <Loading/> : <div/>}

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
        <VersionItem version={info}/>
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
                <FormControl
                  margin="normal"
                  required
                  fullWidth
                >
                  <FormGroup>
                    <InputLabel id="version-label">取得するバージョン</InputLabel>
                    <Select
                      labelId="version-label"
                      id="version"

                      label="取得するバージョン"
                      {...register('version')}
                      variant="outlined"
                    >
                      <MenuItem value="">
                        <em>選択してください</em>
                      </MenuItem>
                      {pearInformation?.map((info, index) => (
                        <MenuItem value={info.version} key={index}>{info.version}</MenuItem>
                      ))}
                    </Select>
                  <FormHelperText>取得したいデータのバージョンを選択してください。</FormHelperText>
                  </FormGroup>
                </FormControl>

                <TextField
                  margin="normal"
                  required
                  fullWidth
                  id="email"
                  label="メールアドレス"
                  autoComplete="email"
                  autoFocus
                  variant="outlined"
                  {...register('email')}
                />
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
                  />
                </div>
              </div>
            }

            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              送信
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  </div>
}

export default Version