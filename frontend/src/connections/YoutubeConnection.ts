import axios from "axios";

export const getTitleFromID = async (id: string) => {
    try {
        return (
            await axios.get(
                "https://www.youtube.com/oembed?url=https://www.youtube.com/watch?v=" +
                    id
            )
        ).data.title;
    } catch {
        return "";
    }
};
