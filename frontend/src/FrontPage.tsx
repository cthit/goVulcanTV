import Iframe from "react-iframe";
import * as React from "react";
import "./FrontPage.css";
import { useEffect, useRef } from "react";
import { getNext } from "./connections/BackendConnection";

export default function FrontPage() {
    const [video, setVideo] = React.useState("");
    useEffect(() => {
        getNext().then(vid => setVideo(vid));
    }, []);
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
