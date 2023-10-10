import React, {useState} from 'react';
import ClientItem from './components/client-item';
import ClientDirectoryScanPage from './components/ClientDirectoryScanPage';
const ClientList = (props) => {

    const [isOpen, setIsOpen] = useState(false);
    //check to see whether client ignore list exists. will change the view
    const [isEmpty, setIsEmpty] = useState(true);
    const [popupid, setpopupid] = useState();
    const toggleNewPopup = (client) => {
       // props.toggle(client);
       console.log("ignorelist",client.ignorepath);
       setIsEmpty(!isEmpty);
       setIsOpen(!isOpen);
       setpopupid(client);
       console.log("popupclient",popupid)
        // setIsOpen(!isOpen);
    }
    return(
        <div className="grid-container-list" >
            {props.clients.map(client =>(
                <ClientItem  toggle={toggleNewPopup} key={client.id} client={client}/>
            ))}
             {isOpen && <ClientDirectoryScanPage toggle={props.toggle} client={popupid} ignorelist={isEmpty}/>}
       </div>
    );

}

export default ClientList