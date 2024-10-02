// src/utils/errorHandler.js

export const handleError = (error) => {
  if (error.response) {
    // Server responded with a status other than 2xx
    console.error("Error Response:", error.response.data);
    return (
      error.response.data.message || "An error occurred. Please try again."
    );
  } else if (error.request) {
    // Request was made but no response received
    console.error("No Response:", error.request);
    return "No response from the server. Please check your network.";
  } else {
    // Something happened in setting up the request
    console.error("Error", error.message);
    return "Request setup error. Please try again.";
  }
};
