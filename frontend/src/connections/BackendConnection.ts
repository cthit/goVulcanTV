import axios from "axios";
import { getTitleFromID } from "./YoutubeConnection";

export const getVideos = async () => {
    const videos: [any] = (await axios.get("/api/page_contents")).data.data;
    const promises = videos.map((video: any) => {
        return new Promise((resolve, reject) => {
            getTitleFromID(video.youtubeID)
                .then((title: any) => {
                    resolve({ ...video, title });
                })
                .catch(err => reject(err));
        });
    });
    const videosWithTitle = await Promise.all(promises);
    console.log(videosWithTitle);
    return videosWithTitle;
};

export const addVideo = async (video: any) => {
    return (await axios.post("/api/page_contents", video)).data;
};

export const getCurrent = async () => {
    return (await axios.get("/api/videos/current")).data.data;
};

export const override = async (id: number) => {
    return (await axios.put("/api/videos/override/" + id)).data.success;
};
