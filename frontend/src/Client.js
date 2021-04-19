import logo from './logo.svg';
import React, {useState, useEffect}from 'react';
import './App.css';
import ClientItem from './components/client-item'


function Client() {

  useEffect(() =>{
    fetchItems();
  },[]);


  const [items, setItems] = useState([]);
  const fetchItems = async() =>{
    const data = await fetch('http://localhost:8000/api/clients/list');
    const items = await data.json();
    console.log(items);

   // console.log(items.policies.length)
 
    setItems(items)
  }


  // items.map(client => (
  //   c = client.policies.length
  // ));
 
  return (
    <div className="grid-container">
        <input type="text" className="searchbox"></input>
        
        {items.map(client =>(
          <ClientItem  key={client.id} id = {client.id} name={client.clientname} policies={client.policies} backups={client.backups}/>
        ))}
          
    </div>
  );
}

export default Client;
