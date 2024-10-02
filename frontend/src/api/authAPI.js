import axiosInstance from "./axiosInstance";

// SignUp API call
export const signUp = async (userData) => {
  try {
    const response = await axiosInstance.post("auth/signup", userData);
    return response.data; // Return the API response data
  } catch (error) {
    throw error;
  }
};

// SignIn API call
export const signIn = async (credentials) => {
  try {
    const response = await axiosInstance.post("auth/signin", credentials);
    return response.data; // Return the API response data
  } catch (error) {
    console.error("Sign In failed:", error.response || error.message);
    throw error;
  }
};
