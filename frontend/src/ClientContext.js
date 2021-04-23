import React,{ useState, createContext } from 'react';

export const ClientContext = createContext();


export const ClientProvider = (props) =>{
    
    const [clients, setClients] = useState([])


    return (
        <ClientContext.Provider value={[clients, setClients]}>
            {props.children}
        </ClientContext.Provider>
    );
} 
