import { Link } from "@material-ui/core";
import React, { useEffect, useState } from "react";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardMedia from "@material-ui/core/CardMedia";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import "./CurrentlyPlaying.css";
import useInterval from "../../utils/useInterval";
import { getCurrent } from "../../connections/BackendConnection";
import { getTitleFromID } from "../../connections/YoutubeConnection";
import Iframe from "react-iframe";

const CurrentlyPlaying = () => {
    const [video, setVideo] = useState("");
    const [title, setTitle] = useState("");
    useInterval(() => {
        getCurrent().then(vid => {
            console.log(vid.id);
            setVideo(vid.id);
        });
    }, 1000);
    useEffect(() => {
        getTitleFromID(video).then(_title => setTitle(_title));
    }, [video]);
    return (
        <Card className="currently-playing">
            <CardHeader title="Currently Playing" />
            {/* <CardMedia
                className="thumbnail"
                image="https://img.youtube.com/vi/OCHDsn6K8qw/hqdefault.jpg"
                title="Thumbnail"
            /> */}
            <div className="thumbnail">
                <Iframe
                    className="iframe-player"
                    url={`https://www.youtube.com/embed/${video}?playlist=${video}&autoplay=1&mute=1&hd=0&loop=1&cc_load_policy=1&cc_lang_pref=en`}
                />
            </div>

            <CardContent>
                <Typography variant="h4">{title}</Typography>
            </CardContent>
        </Card>
    );
};

export default CurrentlyPlaying;
