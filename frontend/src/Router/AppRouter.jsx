import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { Home } from '../Pages/Home';
import { SignInPage } from '../Pages/SignIn';
import { SignUpPage } from '../Pages/SignUp';
import { Dashboard } from '../Pages/Dashboard';
import { PAGE_URLS } from '../Router/paths';

const AppRouter = () => {
 const isSignedIn = false;

  return (
    <Routes>
      {/* Public Routes */}
      <Route 
        path={PAGE_URLS.HOME} 
        element={isSignedIn ? <Navigate to={PAGE_URLS.DASHBOARD} /> : <Home />} 
      />
      <Route 
        path={PAGE_URLS.SIGN_IN} 
        element={isSignedIn ? <Navigate to={PAGE_URLS.DASHBOARD} /> : <SignInPage />} 
      />
      <Route 
        path={PAGE_URLS.SIGN_UP} 
        element={isSignedIn ? <Navigate to={PAGE_URLS.DASHBOARD} /> : <SignUpPage />} 
      />

      {/* Protected Routes */}
      <Route 
        path={PAGE_URLS.DASHBOARD} 
        element={isSignedIn ? <Dashboard /> : <Navigate to={PAGE_URLS.SIGN_IN} />} 
      />
     

      {/* Catch-all Redirect */}
      <Route path="*" element={<Navigate to={PAGE_URLS.HOME} />} />
    </Routes>
  );
};

export default AppRouter;
