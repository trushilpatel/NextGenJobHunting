// src/pages/Home/Home.jsx

import React from 'react';
import { Box, Button, Container, Typography, Grid } from '@mui/material';


const Home = () => {
  return (
    <Container maxWidth="lg">
      {/* Hero Section */}
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          height: '80vh',
          textAlign: 'center',
          padding: 4,
          background: 'linear-gradient(135deg, #2E3B55 30%, #4F5D75 100%)',
          color: 'white',
          borderRadius: 2,
marginTop: 8,       }}

      >
        <Typography variant="h2" component="h1" gutterBottom>
          Welcome to NextGenJobHunting
        </Typography>
        <Typography variant="h5" component="p" gutterBottom>
          Your one-stop platform to find, manage, and land your dream job in the tech industry.
        </Typography>
        <Button 
          variant="contained" 
          color="secondary" 
          size="large" 
          sx={{ marginTop: 4 }}
        >
          Get Started
        </Button>
      </Box>

      {/* Features Section */}
      <Box sx={{ marginTop: 8, marginBottom: 8,  }}>
        <Typography variant="h4" align="center" 
       sx={{backgroundColor:'background.card', color: 'white',
          borderRadius: 2, padding: 4}} gutterBottom>
          Why Choose NextGenJobHunting?
        </Typography>
        <Grid container spacing={4} sx={{ marginTop: 4 }}>
          <Grid item xs={12} md={4}>
            <Box
              sx={{
                textAlign: 'center',
                padding: 4,
                borderRadius: 2,
                backgroundColor: 'background.card',
                boxShadow: 3,
              }}
            >
              <Typography variant="h6" component="h3" gutterBottom>
                Tailored Job Recommendations
              </Typography>
              <Typography>
                We analyze your skills and preferences to recommend the best job opportunities for you.
              </Typography>
            </Box>
          </Grid>
          <Grid item xs={12} md={4}>
            <Box
              sx={{
                textAlign: 'center',
                padding: 4,
                borderRadius: 2,
                backgroundColor: 'background.card',
                boxShadow: 3,
              }}
            >
              <Typography variant="h6" component="h3" gutterBottom>
                Seamless Application Process
              </Typography>
              <Typography>
                Apply to multiple job openings with a single click and track the status in real-time.
              </Typography>
            </Box>
          </Grid>
          <Grid item xs={12} md={4}>
            <Box
              sx={{
                textAlign: 'center',
                padding: 4,
                borderRadius: 2,
                backgroundColor: 'background.card',
                boxShadow: 3,
              }}
            >
              <Typography variant="h6" component="h3" gutterBottom>
                Personalized Career Dashboard
              </Typography>
              <Typography>
                Manage your job applications, track your progress, and stay organized with our custom dashboard.
              </Typography>
            </Box>
          </Grid>
        </Grid>
      </Box>

      {/* CTA Section */}
      <Box
        sx={{
          textAlign: 'center',
          padding: 6,
          backgroundColor: 'secondary.main',
          color: 'white',
          borderRadius: 2,
          marginBottom:10,
        }}
      >
        <Typography variant="h4" component="h2" gutterBottom>
          Ready to take the next step in your career?
        </Typography>
        <Button variant="contained" color="primary" size="large">
          Join Now
        </Button>
      </Box>
    </Container>
  );
};

export default Home;
