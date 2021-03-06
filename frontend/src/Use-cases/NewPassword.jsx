import {
    Button,
    Card,
    CardContent,
    makeStyles,
    Typography
} from "@material-ui/core";
import { useContext, useEffect, useState } from "react";
import ReactPasswordStrength from "react-password-strength";
import CheckCircleIcon from "@material-ui/icons/CheckCircle";
import CancelIcon from "@material-ui/icons/Cancel";
import Axios from "axios";
import UserContext from "../UserContext";
import { Redirect } from "react-router";

const useStyles = makeStyles({
    root: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center"
    },
    card: {
        width: "29rem",
        marginTop: "2rem"
    },
    again: {
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center"
    },
    inputField: {
        width: "25rem"
    },
    inputFieldAgain: {
        width: "25rem",
        marginTop: "1rem"
    },
    submitButton: {
        flow: "right",
        marginTop: "1rem"
    },
    matchIcon: {
        marginTop: "1rem"
    }
});

const scoreWords = ["weak", "okay", "good", "strong", "stronger"];
const passDefault = { isValid: false, password: "" };

const PasswordField = ({ className, placeholder, onChange, logged_in }) => (
    <ReactPasswordStrength
        className={className}
        minLength={5}
        minScore={2}
        scoreWords={scoreWords}
        inputProps={{
            name: "password_input",
            autoComplete: "off",
            placeholder: placeholder,
            disabled: !logged_in
        }}
        changeCallback={onChange}
    />
);

const NewPassword = () => {
    const [user, setUser] = useContext(UserContext);
    const classes = useStyles();
    const [password, setPassword] = useState(passDefault);
    const [passwordAgain, setPasswordAgain] = useState(passDefault);
    const [error, setError] = useState();
    const [matchingIcon, setMatchingIcon] = useState(null);
    const [success, setSuccess] = useState(false);

    const handleSubmit = () => {
        if (!password.password) {
            setError("You must enter a new password!");
            return;
        }
        if (!password.isValid) {
            setError("Your password is too weak!");
            return;
        }
        if (password.password !== passwordAgain.password) {
            setError("You need to enter the same password twice!");
            return;
        }
        setError("");

        Axios.post("/api/admin/setPassword", {
            password: password.password
        })
            .then(async () => {
                await Axios.post("/api/logout");
                setSuccess(true);
                setUser({ logged_in: false, name: "" });
            })
            .catch(err => {
                setError(err.response.data);
            });
    };

    useEffect(() => {
        if (!password.password && !passwordAgain.password) {
            setMatchingIcon(null);
            return;
        }
        if (password.password === passwordAgain.password) {
            setMatchingIcon(
                <CheckCircleIcon
                    className={classes.matchIcon}
                    style={{ color: "green" }}
                />
            );
        } else {
            setMatchingIcon(
                <CancelIcon
                    className={classes.matchIcon}
                    style={{ color: "red" }}
                />
            );
        }
    }, [password, passwordAgain, classes.matchIcon]);

    useEffect(() => {
        if (!user.logged_in) {
            Axios.get("/api/admin/checkLogin")
                .then(res => {
                    setUser(res.data);
                })
                .catch(err => {
                    if (
                        err.response.data &&
                        err.response.data.match(/^http/g)
                    ) {
                        window.location.href = err.response.data;
                    }
                });
        }
    }, [user, setUser]);

    return (
        <div className={classes.root}>
            {success ? <Redirect to="/success" /> : null}
            <Card className={classes.card}>
                <CardContent>
                    <Typography variant="h5">Hi {user.name}!</Typography>
                    <Typography variant="h6">
                        Enter your new password
                    </Typography>
                    <PasswordField
                        className={classes.inputField}
                        placeholder={"New password"}
                        onChange={e => setPassword(e)}
                        logged_in={user.logged_in}
                    />
                    <div className={classes.again}>
                        <PasswordField
                            className={classes.inputFieldAgain}
                            placeholder={"New password again"}
                            onChange={e => setPasswordAgain(e)}
                            logged_in={user.logged_in}
                        />
                        {matchingIcon}
                    </div>
                    <Typography
                        variant="subtitle1"
                        style={{
                            visibility: error ? "visible" : "hidden",
                            color: "#CC0000"
                        }}
                    >
                        {error}
                    </Typography>
                    <Button
                        className={classes.submitButton}
                        variant="contained"
                        color="primary"
                        onClick={handleSubmit}
                    >
                        <Typography variant="button">Submit</Typography>
                    </Button>
                </CardContent>
            </Card>
        </div>
    );
};

export default NewPassword;
