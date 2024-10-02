// src/index.js

import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import CssBaseline from "@mui/material/CssBaseline"; // Import MUI CssBaseline

ReactDOM.render(
  <React.StrictMode>
    {/* Add CssBaseline here */}
    <CssBaseline />
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
