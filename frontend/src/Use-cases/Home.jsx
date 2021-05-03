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

const Home = () => {
    const classes = useStyle();
    return (
        <div className={classes.root}>
            <Card className={classes.card}>
                <CardContent>
                    <Typography variant="h6">Welcome to delta!</Typography>
                    <Typography variant="subtitle1">
                        This application allows you to change your password in
                        the legacy account system at the IT student division.
                        The old account system is still used for managing access
                        committee wikis and SSH access to the servers of digIT.
                    </Typography>
                    <Link to="/new_password">
                        <Button color="primary" variant="contained">
                            Login and change password
                        </Button>
                    </Link>
                </CardContent>
            </Card>
        </div>
    );
};

export default Home;
