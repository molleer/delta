import {
    Button,
    Card,
    CardContent,
    makeStyles,
    Typography
} from "@material-ui/core";
import { Link } from "react-router-dom";

const useStyle = makeStyles({
    root: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center"
    },
    card: {
        width: "29rem",
        marginTop: "2rem"
    }
});

const Success = () => {
    const classes = useStyle();
    return (
        <div className={classes.root}>
            <Card className={classes.card}>
                <CardContent>
                    <Typography variant="h6">Success!</Typography>
                    <Typography variant="subtitle1">
                        You have successfully changed your password and have
                        been logged out!
                    </Typography>
                    <Link to="/">
                        <Button color="primary" variant="contained">
                            Go back to home
                        </Button>
                    </Link>
                </CardContent>
            </Card>
        </div>
    );
};

export default Success;
