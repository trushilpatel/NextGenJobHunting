// src/pages/SignUp/SignUp.jsx

import React, { useState } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as Yup from 'yup';
import { signUp } from '../../api/authAPI'; // Import signUp API call
import { useNavigate } from 'react-router-dom'; // For navigation after successful sign-up

// Validation schema using Yup
const validationSchema = Yup.object().shape({
  firstName: Yup.string().required('First Name is required'),
  lastName: Yup.string().required('Last Name is required'),
  username: Yup.string().required('Username is required'),
  email: Yup.string().email('Invalid email format').required('Email is required'),
  password: Yup.string()
    .min(6, 'Password must be at least 6 characters long')
    .required('Password is required'),
  confirmPassword: Yup.string()
    .oneOf([Yup.ref('password'), null], 'Passwords must match')
    .required('Confirm Password is required'),
});

const SignUp = () => {
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(validationSchema),
  });

  const [apiError, setApiError] = useState(null); // State to handle API errors
  const navigate = useNavigate(); // Use navigate for redirecting after signup

  const onSubmit = async (data) => {
    try {
      // Call the signUp API
      const response = await signUp({
        firstName: data.firstName,
        lastName: data.lastName,
        username: data.username,
        email: data.email,
        password: data.password,
      });

      console.log('User registered:', response);

      // After successful sign-up, redirect to sign-in page
      navigate('/sign-in');
    } catch (error) {
      // Handle errors returned by the API
      setApiError(error.response?.data?.message || 'An error occurred during sign-up.');
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
          Sign Up
        </Typography>

        {apiError && (
          <Typography variant="body2" color="error" gutterBottom>
            {apiError}
          </Typography>
        )}

        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ width: '100%' }}>
          {/* First Name Field */}
          <TextField
            label="First Name"
            variant="outlined"
            fullWidth
            margin="normal"
            {...register('firstName')}
            error={!!errors.firstName}
            helperText={errors.firstName?.message}
          />

          {/* Last Name Field */}
          <TextField
            label="Last Name"
            variant="outlined"
            fullWidth
            margin="normal"
            {...register('lastName')}
            error={!!errors.lastName}
            helperText={errors.lastName?.message}
          />

          {/* Username Field */}
          <TextField
            label="Username"
            variant="outlined"
            fullWidth
            margin="normal"
            {...register('username')}
            error={!!errors.username}
            helperText={errors.username?.message}
          />

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

          {/* Confirm Password Field */}
          <TextField
            label="Confirm Password"
            variant="outlined"
            type="password"
            fullWidth
            margin="normal"
            {...register('confirmPassword')}
            error={!!errors.confirmPassword}
            helperText={errors.confirmPassword?.message}
          />

          <Button
            type="submit"
            variant="contained"
            fullWidth
            sx={{ marginTop: 2 }}
          >
            Sign Up
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default SignUp;
