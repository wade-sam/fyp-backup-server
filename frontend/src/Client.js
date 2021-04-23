import logo from './logo.svg';
import React, {useState, useEffect, useContext}from 'react';
import './App.css';
import ClientItem from './components/client-item'
import ClientList from './ClientList';
import {ClientContext} from './ClientContext';


function Client() {
  const [clients, setClients] = useContext(ClientContext);
  useEffect(() =>{
    fetchItems();
  },[]);

  const fetchItems = async() =>{
    const data = await fetch('http://localhost:8000/api/clients/list', {method:"GET"});
    const items = await data.json();
    console.log(items);


 
    setClients(items)
  }

  return (
    <div className="grid-container">
        <input type="text" className="searchbox"></input>
        <ClientList clients={clients}/>

          
    </div>
  );
}

export default Client;
