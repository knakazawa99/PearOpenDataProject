import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import CssBaseline from '@mui/material/CssBaseline';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import * as React from 'react';

import AdminAuth from 'components/features/admin_auth/Main';

const theme = createTheme();

const AdminAuthPage = () => {
  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <AdminAuth/>
        </Box>
      </Container>
    </ThemeProvider>
  );
}

export default AdminAuthPage
