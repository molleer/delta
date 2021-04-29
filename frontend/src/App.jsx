import {
    Button,
    Card,
    CardContent,
    makeStyles,
    Typography
} from "@material-ui/core";
import { useState } from "react";
import ReactPasswordStrength from "react-password-strength";

const useStyles = makeStyles({
    root: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center"
    },
    card: {
        width: "32rem",
        marginTop: "2rem"
    },
    form: {
        display: "flex",
        justifyContent: "space-between"
    },
    inputField: {
        width: "25rem"
    }
});

const App = () => {
    const classes = useStyles();
    const [password, setPassword] = useState();
    return (
        <div className={classes.root}>
            <Card className={classes.card}>
                <CardContent>
                    <Typography variant="h6">
                        Enter your new password
                    </Typography>
                    <div className={classes.form}>
                        <ReactPasswordStrength
                            className={classes.inputField}
                            minLength={5}
                            minScore={2}
                            scoreWords={[
                                "weak",
                                "okay",
                                "good",
                                "strong",
                                "stronger"
                            ]}
                            inputProps={{
                                name: "password_input",
                                autoComplete: "off",
                                placeholder: "New password"
                            }}
                            changeCallback={e => setPassword(e.password)}
                        />
                        <Button onClick={() => console.log(password)}>
                            <Typography variant="button">Submit</Typography>
                        </Button>
                    </div>
                </CardContent>
            </Card>
        </div>
    );
};

export default App;
