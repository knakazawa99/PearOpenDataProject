import { Link, Typography } from '@mui/material';

const Copyright = () =>{
  return (
    <Typography variant="body2" color="text.secondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://www.eng.niigata-u.ac.jp/~yamazaki/">
        Yamazaki Lab.
      </Link>{' '}
      Niigata Univ, All rights reserved.
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

export default Copyright