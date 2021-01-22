import axios from "axios";
import { getTitleFromID } from "./YoutubeConnection";

export const getVideos = async () => {
    const videos: [any] = (await axios.get("/api/pageContent")).data.data;
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
    return (await axios.post("/api/pageContent", video)).data;
};

export const getCurrent = async () => {
    return await (await axios.get("/api/current")).data.data;
};
