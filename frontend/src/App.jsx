// src/App.jsx
import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import {Header }from './Components/Header';
import {Footer} from './Components/Footer'; // Updated import path (ensure correct case)
import {AppRouter} from './Router';         // Import the AppRouter component
import { ThemeProvider } from '@mui/material/styles'; // Import ThemeProvider for theming
import CssBaseline from '@mui/material/CssBaseline';  // Import CssBaseline for reset
import { theme } from './Theme/Constants';

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline /> 

      <BrowserRouter>
        <Header />
        <AppRouter />
        <Footer />
      </BrowserRouter>

    </ThemeProvider>
  );
}

export default App;
