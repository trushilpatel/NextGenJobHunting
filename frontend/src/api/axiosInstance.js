import axios from "axios";

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL;

// Create an Axios instance with default configuration
const axiosInstance = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Optional: Add request and response interceptors (e.g., to handle token authorization)
axiosInstance.interceptors.request.use(
  (config) => {
    // Example: Attach an authorization token (if you have it)
    const token = localStorage.getItem("token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // Handle errors (e.g., token expiration, network errors)
    return Promise.reject(error);
  }
);

export default axiosInstance;
