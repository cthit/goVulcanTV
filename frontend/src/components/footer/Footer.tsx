import { Link, Typography } from "@material-ui/core";
import React from "react";

const Footer = ({ version }: { version: any }) => {
    return (
        <Typography
            variant="body2"
            align="center"
            color="textPrimary"
            style={{ alignSelf: "end", width: "100%" }}
        >
            <p>
                <b>goVulcanTV Admin</b> version {version}
                <br />
                Made with ❤️ by{" "}
                <Link href="https://github.com/viddem">Vidde</Link>{" "}
                (digIT&apos;18) &{" "}
                <Link href="https://github.com/swexbe">Swexbe</Link>{" "}
                (digit&apos;20)
                <br />
                <Link href="https://github.com/swexbe/goVulcanTV">
                    Source on Github
                </Link>
            </p>
        </Typography>
    );
};
export default Footer;
