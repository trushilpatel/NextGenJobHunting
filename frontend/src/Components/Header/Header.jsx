// src/components/Header.jsx

import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import { useNavigate, useLocation, Link } from 'react-router-dom';
import { COLORS } from '../../Theme/Constants'; // Import the COLORS

const Header = () => {
  const location = useLocation(); // Get the current path
  const navigate = useNavigate(); // To navigate between pages
  const [isAuthenticated, setIsAuthenticated] = useState(false); // Simulate the auth state

  const handleLogout = () => {
    setIsAuthenticated(false);
    navigate('/'); // Redirect to home after logout
  };

  const shouldShowLoginButton = !['/sign-in', '/sign-up'].includes(location.pathname) && !isAuthenticated;

  return (
    <AppBar
      position="static"
      sx={{
        background: COLORS.Gradient1, // Using gradient from constants
        boxShadow: 'none', // No shadow for a clean look
        borderBottom: `1px solid ${COLORS.textSecondary}`, // Subtle border for separation
      }}
    >
      <Toolbar>
        <IconButton
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{
            color: COLORS.textPrimary, // White text from constants
          }}
        >
          <MenuIcon />
        </IconButton>
        <Typography
          variant="h6"
          component={Link}
          to="/"
          sx={{
            flexGrow: 1,
            color: COLORS.textPrimary, // White text for the logo text from constants
            textDecoration: 'none', // Remove underline from the logo link
          }}
        >
          NextGenJobHunting
        </Typography>
        {shouldShowLoginButton ? (
          <Button
            component={Link}
            to="/sign-in"
            sx={{
              backgroundColor: COLORS.buttonPrimary, // Using button primary color from constants
              border: `2px solid ${COLORS.buttonText}`, // White border from constants
              color: COLORS.buttonText, // White text from constants
              paddingLeft: 2,
              paddingRight: 2,
              '&:hover': {
                backgroundColor: COLORS.buttonHover, // Darker orange on hover from constants
                borderColor: COLORS.buttonText, // Ensure white border on hover
              },
            }}
          >
            Login
          </Button>
        ) : (
          isAuthenticated && (
            <Button
              onClick={handleLogout}
              sx={{
                backgroundColor: COLORS.buttonPrimary, // Using button primary color from constants
                border: `2px solid ${COLORS.buttonText}`, // White border from constants
                color: COLORS.buttonText, // White text from constants
                paddingLeft: 2,
                paddingRight: 2,
                '&:hover': {
                  backgroundColor: COLORS.buttonHover, // Darker orange on hover from constants
                  borderColor: COLORS.buttonText, // Ensure white border on hover
                },
              }}
            >
              Logout
            </Button>
          )
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Header;
