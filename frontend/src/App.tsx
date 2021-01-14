import React from "react";
import logo from "./logo.svg";
import "./App.css";
import {
  AppBar,
  Button,
  IconButton,
  Paper,
  Typography,
} from "@material-ui/core";
import { Toolbar } from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";
import CurrentlyPlaying from "./components/currently-playing/CurrentlyPlaying";

function App() {
  return (
    <div className="App">
      <AppBar position="static" className="App-header">
        <Toolbar>
          <Typography variant="h6" color="inherit">
            goVulcanTV Admin
          </Typography>
        </Toolbar>
      </AppBar>
      <div
        style={{
          flex: 1,
          alignItems: "center",
          display: "flex",
          flexDirection: "column",
          marginLeft: "100px",
          marginRight: "100px",
          marginTop: "20px",
        }}
      >
        <CurrentlyPlaying />
      </div>
    </div>
  );
}

export default App;
