import React from "react";
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";

// Layouts
import Login from "./components/login/login";
import Studies from "./components/studies/studies";
import Viewer from "./components/viewer/viewer";

// Styles
import { createMuiTheme } from "@material-ui/core/styles";
import { ThemeProvider } from "@material-ui/styles";
import "./App.css";
import { green } from "@material-ui/core/colors";

function App() {
  // dark theme with CHC primary color
  const theme = createMuiTheme({
    palette: {
      type: "dark",
      primary: {
        main: "#EF3755",
      },
      secondary: {
        main: "#007228",
      },
    },
  });

  return (
    <ThemeProvider theme={theme}>
      <BrowserRouter>
        {/* Redirect to Login if user landed on Home page */}
        <Switch>
          <Route path={["/"]} exact render={() => <Redirect to="/login" />} />
        </Switch>

        {/* Main App */}
        <div className="App">
          <header>Change Healthcare - Viewer</header>
          <Switch>
            <Route path="/login" component={() => <Login />} />
            <Route path="/studies" component={() => <Studies />} />
            <Route path="/viewer" component={() => <Viewer />} />
          </Switch>
        </div>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
