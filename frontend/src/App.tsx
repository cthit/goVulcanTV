import * as React from "react";
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
import CurrentlyPlaying from "./components/currently-playing/CurrentlyPlaying";
import VideoTable from "./components/video-table/VideoTable";

function App() {
    return (
        <div className="app">
            <AppBar position="static" className="header">
                <Toolbar>
                    <Typography variant="h6" color="inherit">
                        goVulcanTV Admin
                    </Typography>
                </Toolbar>
            </AppBar>
            <div className="center-box">
                <CurrentlyPlaying />
                <VideoTable />
            </div>
        </div>
    );
}

export default App;
