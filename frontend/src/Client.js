import React, {useState, useEffect, useContext}from 'react';
import './App.css';
import ClientItem from './components/client-item'
import ClientList from './ClientList';
import {ClientContext} from './ClientContext';
import ClientDirectoryScanPage from './components/ClientDirectoryScanPage';


function Client() {
  const {clients} = useContext(ClientContext);
  const {setClients} = useContext(ClientContext);
  useEffect(() =>{
    fetchItems();
  },[]);

  const fetchItems = async() =>{
    const data = await fetch('http://localhost:8000/api/clients/list', {method:"GET"});
    const items = await data.json();
    console.log(typeof items); 
    setClients(items)
  }
  var test = {}
  const [isOpen, setIsOpen] = useState(false);
  const [client, setClient] = useState();
  const toggleNewPopup = (id) => {
      for (var item in clients){
        console.log(clients[item]);
        if (clients[item].id == id){
          setClient(clients[item]);
          break;
         
        }
      }
      setIsOpen(!isOpen);
  }
  console.log("FOUND", client);

  return (
    <div className="grid-container">
        <input type="text" className="searchbox"></input>
        <ClientList toggle={toggleNewPopup}clients={clients}/>
        {isOpen && <ClientDirectoryScanPage toggle={toggleNewPopup} client={client}/>}

          
    </div>
  );
}

export default Client;
