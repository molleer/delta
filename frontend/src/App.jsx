import { BrowserRouter, Switch } from "react-router-dom";
import { Route } from "react-router";
import NewPassword from "./NewPassword";
import Callback from "./Callback";
import { UserContextProvider } from "./UserContext";

const App = () => (
    <UserContextProvider>
        <BrowserRouter>
            <Switch>
                <Route exact path="/" component={NewPassword} />
                <Route path="/callback" component={Callback} />
                <Route path="/" component={() => <h2>404 Page not found</h2>} />
            </Switch>
        </BrowserRouter>
    </UserContextProvider>
);
export default App;
