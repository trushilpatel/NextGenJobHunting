import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#F5F5F5", // Very light gray for primary background
      light: "#FFFFFF", // Pure white for lighter areas
    },
    secondary: {
      main: "#FF8A65", // A soft peach color for buttons and call-to-action
    },
    background: {
      default: "#F5F5F5", // Light gray background for the whole page
      paper: "#FFFFFF", // White background for cards and paper components
      card: "#3d3d3d",
      base: "linear-gradient(135deg, #4F5D75 30%, #2E3B55 100%)",
    },
    text: {
      primary: "#FFFFFF", // White text for headings and primary elements
      secondary: "#E0E0E0", // Lighter gray for secondary text
    },
  },
  typography: {
    h2: {
      fontWeight: 600,
      color: "#FFFFFF", // White for headings
    },
    h4: {
      fontWeight: 500,
      color: "#FFFFFF", // Consistent white for subheadings
    },
    h6: {
      fontWeight: 500,
      color: "#FFFFFF", // White for smaller headings
    },
    body1: {
      color: "#E0E0E0", // Lighter gray for body text
    },
  },
});
const COLORS = {
  primaryBackground: "#2E3B55", // Dark blue background
  secondaryBackground: "#4F5D75", // Lighter blue background
  textPrimary: "#FFFFFF", // White text for headings and primary elements
  textSecondary: "#B0BEC5", // Light gray text for secondary elements
  inputBackground: "#455A64", // Dark gray background for input fields
  inputBorder: "#B0BEC5", // Light gray for borders
  inputHoverBorder: "#FF8A65", // Peach color on hover for input borders
  buttonPrimary: "#FF8A65", // Peach color for buttons
  buttonHover: "#FF7043", // Slightly darker peach for hover state
  buttonText: "#FFFFFF", // White text for buttons
  Gradient1: "linear-gradient(135deg, #2E3B55 30%, #4F5D75 100%)", // Gradient
};

export { theme, COLORS };
