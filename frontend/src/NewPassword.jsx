import {
    Button,
    Card,
    CardContent,
    makeStyles,
    Typography
} from "@material-ui/core";
import { useEffect, useState } from "react";
import ReactPasswordStrength from "react-password-strength";
import CheckCircleIcon from "@material-ui/icons/CheckCircle";
import CancelIcon from "@material-ui/icons/Cancel";

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

const PasswordField = ({ className, placeholder, onChange }) => (
    <ReactPasswordStrength
        className={className}
        minLength={5}
        minScore={2}
        scoreWords={scoreWords}
        inputProps={{
            name: "password_input",
            autoComplete: "off",
            placeholder: placeholder
        }}
        changeCallback={onChange}
    />
);

const NewPassword = () => {
    const classes = useStyles();
    const [password, setPassword] = useState({ isValid: false, password: "" });
    const [passwordAgain, setPasswordAgain] = useState({
        isValid: false,
        password: ""
    });
    const [error, setError] = useState();
    const [matchingIcon, setMatchingIcon] = useState(null);

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
        //TODO: Send request to backend
        console.log(password.password);
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

    return (
        <div className={classes.root}>
            <Card className={classes.card}>
                <CardContent>
                    <Typography variant="h5">Hi ITStud!</Typography>
                    <Typography variant="h6">
                        Enter your new password
                    </Typography>
                    <PasswordField
                        className={classes.inputField}
                        placeholder={"New password"}
                        onChange={e => setPassword(e)}
                    />
                    <div className={classes.again}>
                        <PasswordField
                            className={classes.inputFieldAgain}
                            placeholder={"New password again"}
                            onChange={e => setPasswordAgain(e)}
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
