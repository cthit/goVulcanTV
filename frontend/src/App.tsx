import * as React from "react";
import "./App.css";
import { AppBar, Typography } from "@material-ui/core";
import { Toolbar } from "@material-ui/core";
import CurrentlyPlaying from "./components/currently-playing/CurrentlyPlaying";
import VideoTable from "./components/video-table/VideoTable";
import packageJson from "../package.json";
import Footer from "./components/footer/Footer";
import AddVideo from "./components/add-video/AddVideo";
import { getVideos } from "./connections/BackendConnection";

function App() {
    const [videos, setVideos] = React.useState([]);
    const reloadVideos = () => {
        getVideos().then((vids: any) => setVideos(vids));
    };
    React.useEffect(() => {
        reloadVideos();
    }, []);
    return (
        <div className="app">
            <AppBar position="static" className="header">
                <Toolbar>
                    <Typography variant="h6" color="inherit">
                        📺goVulcanTV Admin
                    </Typography>
                </Toolbar>
            </AppBar>
            <div className="center-box">
                <CurrentlyPlaying />
                <AddVideo reloadVideos={reloadVideos} />
                <VideoTable videos={videos} reloadVideos={reloadVideos} />
            </div>
            <Footer version={packageJson.version} />
        </div>
    );
}

export default App;
