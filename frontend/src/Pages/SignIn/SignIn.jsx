// src/pages/SignIn/SignIn.jsx

import React, { useState } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as Yup from 'yup';
import { signIn } from '../../api/authAPI'; // Import signIn API call
import { useNavigate } from 'react-router-dom'; // For navigation after successful sign-in

// Validation schema using Yup
const validationSchema = Yup.object().shape({
  email: Yup.string().email('Invalid email format').required('Email is required'),
  password: Yup.string().required('Password is required'),
});

const SignIn = () => {
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(validationSchema),
  });

  const [apiError, setApiError] = useState(null); // State to handle API errors
  const navigate = useNavigate(); // Use navigate for redirecting after sign-in

  const onSubmit = async (data) => {
    try {
      // Call the signIn API
      const response = await signIn({
        email: data.email,
        password: data.password,
      });

      console.log('User signed in:', response);

      // Store the token (if returned) in localStorage or cookies
      localStorage.setItem('token', response.token);

      // After successful sign-in, redirect to the dashboard or home page
      navigate('/dashboard');
    } catch (error) {
      // Handle errors returned by the API
      setApiError(error.response?.data?.message || 'An error occurred during sign-in.');
    }
  };

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          height: '80vh',
          padding: 4,
          background: 'linear-gradient(135deg, #2E3B55 30%, #4F5D75 100%)',
          borderRadius: 2,
          boxShadow: 2,
          marginTop: 4,
        }}
      >
        <Typography variant="h4" gutterBottom>
          Sign In
        </Typography>

        {apiError && (
          <Typography variant="body2" color="error" gutterBottom>
            {apiError}
          </Typography>
        )}

        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ width: '100%' }}>
          {/* Email Field */}
          <TextField
            label="Email"
            variant="outlined"
            fullWidth
            margin="normal"
            {...register('email')}
            error={!!errors.email}
            helperText={errors.email?.message}
          />

          {/* Password Field */}
          <TextField
            label="Password"
            variant="outlined"
            type="password"
            fullWidth
            margin="normal"
            {...register('password')}
            error={!!errors.password}
            helperText={errors.password?.message}
          />

          <Button
            type="submit"
            variant="contained"
            fullWidth
            sx={{ marginTop: 2 }}
          >
            Sign In
          </Button>
        </Box>

        <Typography variant="body2" sx={{ marginTop: 2 }}>
          Don’t have an account?{' '}
          <a href="/sign-up" style={{ color: '#FF8A65', textDecoration: 'none' }}>
            Sign Up
          </a>
        </Typography>
      </Box>
    </Container>
  );
};

export default SignIn;
