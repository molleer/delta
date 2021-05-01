import { BrowserRouter, Switch } from "react-router-dom";
import { Route } from "react-router";
import NewPassword from "./NewPassword";

const App = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={NewPassword} />
            <Route path="/callback" component={() => <div>Callback</div>} />
            <Route path="/" component={() => <h2>404 Page not found</h2>} />
        </Switch>
    </BrowserRouter>
);
export default App;
