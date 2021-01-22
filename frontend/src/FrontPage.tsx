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
            setVideo(vid.id);
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
                url={"https://www.youtube.com/embed/" + video}
                width="100%"
                height="100%"
            ></Iframe>
        </div>
    );
}
