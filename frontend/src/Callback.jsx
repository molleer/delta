import { useContext, useEffect, useState } from "react";
import { CircularProgress, makeStyles } from "@material-ui/core";
import { Redirect } from "react-router";
import Axios from "axios";
import UserContext from "./UserContext";

const useStyle = makeStyles({
    center: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: () => window.innerHeight
    }
});

const Callback = () => {
    const [, setUser] = useContext(UserContext);
    const [redirect, setRedirect] = useState(false);
    const classes = useStyle();

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        Axios.post(`/api/exchangeCode?grant=${params.get("code")}`)
            .then(res => {
                setUser(res.data);
                setRedirect(true);
            })
            .catch(err => console.log(err));
    }, [setUser]);

    return (
        <div className={classes.center}>
            {redirect ? <Redirect to="/" /> : <CircularProgress size="5rem" />}
        </div>
    );
};

export default Callback;
