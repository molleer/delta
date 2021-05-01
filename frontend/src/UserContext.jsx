import { createContext, useState } from "react";

const UserContext = createContext([{ logged_in: false, name: "" }, () => {}]);

export const UserContextProvider = ({ children }) => {
    const value = useState({ logged_in: false, name: "" });
    return (
        <UserContext.Provider value={value}>{children}</UserContext.Provider>
    );
};

export default UserContext;
