import { Button, Paper, TextField, Typography } from "@material-ui/core";
import * as React from "react";
import { useState } from "react";
import "./AddVideo.css";
import axios from "axios";

const AddVideo = () => {
    const [videoTitle, setVideoTitle] = useState("");
    const [id, setID] = useState("");
    const [addedBy, setAddedBy] = useState("");
    const [description, setDescription] = useState("");

    const idError = () => {
        if (id.length === 0) {
            return false;
        }
        return !id.match(/[a-zA-Z0-9_-]{11}/);
    };

    React.useEffect(() => {
        if (id.match(/[a-zA-Z0-9_-]{11}/)) {
            axios
                .get(
                    "https://www.youtube.com/oembed?url=https://www.youtube.com/watch?v=" +
                        id
                )
                .then(res => setVideoTitle(res.data.title))
                .catch(err => setVideoTitle(""));
        } else {
            setVideoTitle("");
        }
    }, [id]);
    return (
        <Paper className="add-video">
            <form noValidate autoComplete="off">
                <div
                    style={{
                        display: "flex",
                        flexDirection: "row",
                    }}
                >
                    <div className="add-video-input-container">
                        <TextField
                            className="add-video-input"
                            id="id"
                            label="ID"
                            error={idError()}
                            value={id}
                            onChange={event => setID(event.target.value)}
                            variant="outlined"
                        />
                    </div>
                    <div className="add-video-input-container">
                        <TextField
                            className="add-video-input"
                            id="addedBy"
                            label="Added By"
                            value={addedBy}
                            onChange={event => setAddedBy(event.target.value)}
                            variant="outlined"
                        />
                    </div>
                </div>
                <div className="add-video-input-container">
                    <TextField
                        rows={4}
                        className="add-video-input"
                        id="description"
                        label="Description"
                        value={description}
                        onChange={event => setDescription(event.target.value)}
                        variant="outlined"
                        multiline
                    />
                </div>
            </form>
            <div className="add-video-button-container">
                <div>
                    <Typography>{videoTitle}</Typography>
                </div>
                <Button
                    variant="contained"
                    color="primary"
                    disabled={!videoTitle}
                >
                    Add Video
                </Button>
            </div>
        </Paper>
    );
};

export default AddVideo;
