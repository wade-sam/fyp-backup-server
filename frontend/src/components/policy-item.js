import React from 'react';
import '../App.css';
import { Link } from 'react-router-dom'

const PolicyItem = ({id, name, clients, type, runs}) =>{
    var clientcount = 0;
    var backupcount = 0;
    for (const c in clients){
        clientcount = clientcount + 1;
    }
    console.log(clients)
    if (runs != null){
        for (const c in runs){
            backupcount = backupcount + 1;
        }
    }
    return(
        <div className="backup-item-grid-container">
        <p className="backup-label-1">{type}</p>
        <div className="backup-item-name-1">
    <p>{name}</p>
        </div>
        <div className="dropdown">
        <button className="backup-options">...</button>
        <div className="dropdowncontent">
            <Link to={`/clients/${id}`} className="dropdown-item">View Clients</Link>
            <li>Run Policy</li>
            <li>Delete Policy</li>
            <li>Edit Policy</li>
            <li>Edit Clients</li>
        </div>
        
        </div>
    <p className="backup-item-policy">Clients: {clientcount}</p>
    <p className="backup-item-backup">Runs: {backupcount}</p>
    </div>  
    );
};

export default PolicyItem;