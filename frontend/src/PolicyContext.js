import React,{ useState, createContext } from 'react';

export const PolicyContext = createContext();


export const PolicyProvider = (props) =>{
    
    const [policies, setPolicies] = useState([])


    return (
        <PolicyContext.Provider value={[policies, setPolicies]}>
            {props.children}
        </PolicyContext.Provider>
    );
} 

