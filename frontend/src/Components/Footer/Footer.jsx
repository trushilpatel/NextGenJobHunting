// src/components/Footer.jsx

import React from 'react';
import { Box, Typography } from '@mui/material';
import { COLORS } from '../../Theme/Constants';
const Footer = () => {
  return (
    <Box
      sx={{
        width: '100%',
        background: COLORS.Gradient1,
        color: 'white',
        textAlign: 'center',
        padding: 2,
        position: 'fixed',
        bottom: 0,
      }}
    >
      <Typography variant="body1">
        © 2024 NextGenJobHunting. All Rights Reserved.
      </Typography>
    </Box>
  );
};

export default Footer;
