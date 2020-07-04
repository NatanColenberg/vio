import React, { useState } from "react";
import { Redirect } from "react-router-dom";
import { TextField, Button } from "@material-ui/core";

import "./login.css";

export default function Login() {
  const [redirectToStudies, SetRedirectToPatients] = useState(false);
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [credentialsError, serCredentialsError] = useState(false);
  const login = () => {
    if (userName === "Admin" && password === "Admin") {
      SetRedirectToPatients(true);
    } else {
      setUserName("");
      setPassword("");
      serCredentialsError(true);
    }
  };
  return (
    <div className="login-wrap">
      <h1 className="login-title">Login</h1>
      <TextField
        className="login-textField"
        required
        label="User Name"
        value={userName}
        onChange={(e) => {
          setUserName(e.target.value);
          serCredentialsError(false);
        }}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            login();
          }
        }}
      />
      <TextField
        className="login-textField"
        required
        label="Password"
        type="password"
        value={password}
        onChange={(e) => {
          setPassword(e.target.value);
          serCredentialsError(false);
        }}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            login();
          }
        }}
      />
      <p hidden={!credentialsError} className="login-credentialsError">
        Wrong User Name or Password, Please try again.
      </p>
      <Button
        className="login-button"
        variant="contained"
        color="primary"
        onClick={login}
      >
        Login
      </Button>

      {/* Redirect to Patients on tenant Changed and user is in Patient Studies page */}
      {redirectToStudies && <Redirect to="/studies" />}
    </div>
  );
}
