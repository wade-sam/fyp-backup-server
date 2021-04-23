import React, {useState} from 'react';
import ClientItem from './components/client-item';
import useVisibilityCard from './components/useVisibilityCard';
const ClientList = (props) => {
    const [] = useVisibilityCard();
    
    return(
        <div className="grid-container-list" >
            {props.clients.map(client =>(
                <ClientItem  key={client.id} client={client}/>
            ))}
       </div>
    );

}

export default ClientList