import Iframe from "react-iframe";
import * as React from "react";
import "./FrontPage.css";
import { useEffect, useRef } from "react";
import { getCurrent } from "./connections/BackendConnection";
import useInterval from "./utils/useInterval";

export default function FrontPage() {
    const [video, setVideo] = React.useState("");
    useInterval(() => {
        getCurrent().then(vid => {
            console.log(vid);
            setVideo(vid.Video.id);
        });
    }, 1000);
    return (
        <div
            style={{
                height: "100%",
                width: "100%",
                border: "10px",
                boxSizing: "border-box",
            }}
        >
            <Iframe
                url={`https://www.youtube.com/embed/${video}?playlist=${video}&autoplay=1&mute=1&hd=0&loop=1&cc_load_policy=1&cc_lang_pref=en`}
                width="100%"
                height="100%"
            ></Iframe>
        </div>
    );
}
