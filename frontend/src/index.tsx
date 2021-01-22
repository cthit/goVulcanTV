import * as React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { ThemeProvider, createMuiTheme } from "@material-ui/core/styles";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import FrontPage from "./FrontPage";

const theme = createMuiTheme({
    palette: {
        type: "dark",
    },
});

ReactDOM.render(
    <ThemeProvider theme={theme}>
        <Router>
            <Route path="/admin">
                <App />
            </Route>
            <Route path="/">
                <FrontPage />
            </Route>
        </Router>
    </ThemeProvider>,
    document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
