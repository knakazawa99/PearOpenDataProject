
import axios from 'axios';
import { ChangeEvent, useEffect, useState } from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';
import { BASE_URL } from 'config/config';
import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import  { saveAs } from 'file-saver';
import Loading from 'components/ui/Loading';
import { getFileNameFromContentDisposition } from '../../../common/file';
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
  const handleChangeVersion = (e: ChangeEvent<HTMLSelectElement>) => {
    // console.log(event.target.value)
  }

  return <div>
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        {isLoading ? <Loading/> : <div/>}

        {downloadStage == 1 &&
          <div>
          <div>
            <label>取得するバージョン</label>
            <select
              id="version"
              {...register('version')}
              defaultValue=""
              onChange={(e) => handleChangeVersion(e)}
            >
              <option value="">選択してください</option>
              {pearInformation?.map((info, index) => (
                <option value={info.version} key={index}>
                  {info.version}
                </option>
              ))}
            </select>
          </div>

            <div>
            <label>メールアドレス</label>
            <input
              type={'email'}
              id="email"
              {...register('email')}
            />
            </div>
          </div>
        }


        {downloadStage == 2 &&
          <div>
            <div>
              <label>トークン</label>
              <input
                id="token"
                {...register('token')}
              />
            </div>
          </div>
        }

        <input type="submit" value="送信" />
      </form>
    </div>
  </div>
}

export default Version