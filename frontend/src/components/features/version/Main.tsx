
import axios from 'axios';
import { ChangeEvent, useEffect, useState } from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';
import { BASE_URL } from 'config/config';
import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import  { saveAs } from 'file-saver';
import Loading from 'components/ui/Loading';
import { getFileNameFromContentDisposition } from '../../../common/file';
import { Button, FormControl, FormGroup, InputLabel, MenuItem, Select, TextField } from '@mui/material';
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
      ).then((response) => {
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

      }).catch((reason) => {
        unblocking();
        setIsLoading(false)
      })
    }
  }


  const [onSubmit, processing, unblocking] = useBlockDoubleClick(
    onSubmitInner,
  );

  return <div>
    <div>
      {isLoading ? <Loading/> : <div/>}
      <form onSubmit={handleSubmit(onSubmit)}>

        {downloadStage == 1 &&
          <div>
            <FormControl sx={{ m: 3 }} component="fieldset" variant="standard">
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
              </FormGroup>
            </FormControl>

            <FormControl sx={{ m: 3 }} component="fieldset" variant="standard">
              <FormGroup>
                <TextField id="email" label="メールアドレス" variant="outlined" {...register('email')} />
              </FormGroup>
            </FormControl>
          </div>
        }

        {downloadStage == 2 &&
          <div>
            <div>
              <TextField id="token" label="トークン" variant="outlined" {...register('token')} />
            </div>
          </div>
        }
        <FormControl sx={{ m: 3 }} component="fieldset" variant="standard">
          <Button
            variant="contained"
            color="primary"
            type="submit"
          >送信</Button>
        </FormControl>
      </form>
    </div>
  </div>
}

export default Version