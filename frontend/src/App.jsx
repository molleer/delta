import { BrowserRouter, Switch } from "react-router-dom";
import { Route } from "react-router";
import NewPassword from "./Use-cases/NewPassword";
import Callback from "./Use-cases/Callback";
import { UserContextProvider } from "./UserContext";
import { Home } from "./Use-cases";
import Success from "./Use-cases/Success";

const App = () => (
    <UserContextProvider>
        <BrowserRouter>
            <Switch>
                <Route exact path="/" component={Home} />
                <Route exact path="/new_password" component={NewPassword} />
                <Route exact path="/success" component={Success} />
                <Route path="/callback" component={Callback} />
                <Route path="/" component={() => <h2>404 Page not found</h2>} />
            </Switch>
        </BrowserRouter>
    </UserContextProvider>
);
export default App;
