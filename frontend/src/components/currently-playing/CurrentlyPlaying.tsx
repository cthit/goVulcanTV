import { Link } from "@material-ui/core";
import React from "react";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardMedia from "@material-ui/core/CardMedia";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import "./CurrentlyPlaying.css";

const CurrentlyPlaying = () => {
    return (
        <Card className="currently-playing">
            <CardHeader title="Currently Playing" />
            <CardMedia
                className="thumbnail"
                image="https://img.youtube.com/vi/OCHDsn6K8qw/hqdefault.jpg"
                title="Thumbnail"
            />
            <CardContent>
                <Link href="https://google.com" variant="h4">
                    Very cool video
                </Link>
                <Typography variant="body2" color="textPrimary" component="p">
                    This impressive paella is a perfect party dish and a fun
                    meal to cook together with your guests. Add 1 cup of frozen
                    peas along with the mussels, if you like.
                </Typography>
            </CardContent>
        </Card>
    );
};

export default CurrentlyPlaying;
