import * as React from 'react';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';

interface VersionProps {
  version: {
    version: string
    releaseNote: string
    createdAt: Date
  };
}

const VersionItem = (props: VersionProps) => {
  const { version } = props;

  return (
    <Grid item xs={12} md={6}>
      <Card sx={{ display: 'flex' }}>
        <CardContent sx={{ flex: 1 }}>
          <Typography component="h2" variant="h5">
            バージョン {version.version}
          </Typography>
          <Typography variant="subtitle1" color="text.secondary">
            リリース日 {version.createdAt.getFullYear()}.{version.createdAt.getMonth()}.{version.createdAt.getDate()}
          </Typography>
          <Typography variant="subtitle2" paragraph>
            リリースノート: <br/>
            {version.releaseNote}
          </Typography>
        </CardContent>
      </Card>
    </Grid>
  );
}

export default VersionItem
