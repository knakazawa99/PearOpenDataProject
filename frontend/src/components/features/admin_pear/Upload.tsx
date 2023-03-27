import axios from 'axios';
import Box from '@mui/material/Box'
import { LoadingButton } from '@mui/lab';
import { Alert, Button, Checkbox, FormControlLabel, FormGroup, FormHelperText, Typography } from '@mui/material';
import { AlertColor } from '@mui/material/Alert/Alert';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import TextField from '@mui/material/TextField';
import React, { ChangeEvent, useRef, useState } from 'react';
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';

import { useBlockDoubleClick } from 'common/useBlockDoubleClick';
import { VersionValidPattern } from 'common/versionValidPattern';
import { AdminVersionFormValues } from 'components/features/admin_pear/Type';
import { BASE_URL } from 'config/config';


const Upload = () => {
  const [fileName, setFileName] = useState<string>()

  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [isAlertMessage, setIsAlertMessage] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>()
  const [alertMessageSeverity, setAlertMessageSeverity] = useState<AlertColor>("success")
  const [open, setOpen] = useState(false)

  const versionRef = useRef<HTMLInputElement>(null);
  const [versionError, setVersionError] = useState(false);
  const [versionErrorMessage, setVersionErrorMessage] = useState<string>()

  const releaseNoteRef = useRef<HTMLInputElement>(null);
  const [releaseNoteError, setReleaseNoteError] = useState(false);
  const [releaseNoteErrorMessage, setReleaseNoteErrorMessage] = useState<string>()

  const [fileError, setFileError] = useState(false);
  const [fileErrorMessage, setFileErrorMessage] = useState<string>()

  const { register, handleSubmit, setValue, watch } = useForm<AdminVersionFormValues>({})

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
    let valid = true
    valid = formValidation()

    if (submitData.files.length == 0) {
      setFileError(false)
      setFileErrorMessage("ファイルは必須項目です。")
    }

    if (!valid) {
      unblocking()
      return
    }

    event?.preventDefault();
    setIsLoading(true)
    const path = BASE_URL + "/v1/admin/versions/"

    const fileUpload = {
      "version": submitData.version,
      "release_note": submitData.release_note,
      "release_comment": submitData.release_comment,
      "release_flag": submitData.release_flag,
      "file": submitData.files[0]
    }
    await axios.post(
      path,
      fileUpload,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
          'authorization': `Bearer ${localStorage.getItem("jwtToken")}`,
          'x-jwtKey': `${localStorage.getItem("jwtKey")}`,
        }
      }
    ).then((res) => {
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

  const onFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    // @ts-ignore
    setValue("files", e.target.files)
    setFileName(e.target?.files?.[0].name)
  }


  const formValidation = (): boolean => {
    let valid = true;

    const version = versionRef?.current;
    if (version) {
      const ok = version?.validity.valid;
      setVersionError(!ok);
      valid &&= ok;
      if (!ok) {
        setVersionErrorMessage("バージョンを正しい形式で入力してください。(例) 1.0.0 や 1.10.11")
      } else {
        setVersionErrorMessage("")
      }
    }

    const releaseNote = releaseNoteRef?.current;
    if (releaseNote) {
      const ok = releaseNote?.validity.valid;
      setReleaseNoteError(!ok);
      valid &&= ok;
      if (!ok) {
        setReleaseNoteErrorMessage("リリースノートは入力してください。")
      } else {
        setReleaseNoteErrorMessage("")
      }
    }

    return valid
  }
  // TODO: https://chusotsu-program.com/js-base64-encode/
  return (

    <div>
      <Button  variant="outlined" onClick={handleClickOpen}>
        データ追加
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>データ追加</DialogTitle>
        <DialogContent>
          <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{ mt: 1 }}>
            <TextField
              autoFocus
              margin="dense"
              label="バージョン"
              fullWidth
              variant="standard"
              required
              {...register("version")}
              inputRef={versionRef}
              inputProps={ {required: true, pattern: VersionValidPattern} }
              helperText={versionError ? versionErrorMessage: ""}
            />
            <TextField
              autoFocus
              margin="dense"
              label="リリースノート"
              fullWidth
              variant="standard"
              required
              {...register("release_note")}
              inputRef={releaseNoteRef}
              helperText={releaseNoteError ? releaseNoteErrorMessage: ""}
            />
            <TextField
              autoFocus
              margin="dense"
              label="リリースコメント"
              fullWidth
              variant="standard"
              {...register("release_comment")}
            />
            <FormGroup>
              <FormControlLabel control={<Checkbox defaultChecked />} label="リリースフラグ" {...register("release_flag")}/>
            </FormGroup>

            <FormGroup>
              <Button
                variant="contained"
                component="label"
              >
                Upload File
                <input
                  type="file"
                  hidden
                  accept=".zip"
                  onChange={onFileChange}
                />
              </Button>
              <Typography variant="subtitle1" align="left" color="text.secondary">
                {fileName}
              </Typography>
              <FormHelperText>
                {fileError ? fileErrorMessage: ""}
              </FormHelperText>
            </FormGroup>

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

export default Upload