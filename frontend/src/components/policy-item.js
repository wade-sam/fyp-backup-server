import React from 'react';
import '../App.css';
import { Link } from 'react-router-dom'

const PolicyItem = (props) =>{
    var clientcount = 0;
    var backupcount = 0;
    for (const c in props.policy.clients){
        clientcount = clientcount + 1;
    }
    console.log(props.policy.clients)
    if (props.policy.backuprun != null){
        for (const c in props.policy.backuprun){
            backupcount = backupcount + 1;
        }
    }
    return(
        <div className="backup-item-grid-container">
        <p className="backup-label-1">{props.policy.type}</p>
        <div className="backup-item-name-1">
            <p>{props.policy.policyname}</p>
        </div>
        <div className="dropdown">
        <button className="backup-options">...</button>
        <div className="dropdowncontent">
            <Link to={`/clients/${props.policy.id}`} className="dropdown-item">View Clients</Link>
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