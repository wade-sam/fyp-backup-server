import React,{ useState, createContext } from 'react';

export const ClientContext = createContext();


export const ClientProvider = (props) =>{
    
    const [clients, setClients] = useState([])
    const [policies, setPolicies] = useState([])
    
    const value = {
        clients,
        setClients,
        policies,
        setPolicies
    };

    return (
        <ClientContext.Provider value={value} >
            {props.children}
        </ClientContext.Provider>
    );
} 
