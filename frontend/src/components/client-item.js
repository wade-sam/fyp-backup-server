import React from 'react';
import '../App.css';
import { Link } from 'react-router-dom'

const ClientItem = ({id, name, policies, backups}) =>{
    var ID = id;
    var policycount = 0;
    var backupcount = 0;
    for (const c in policies){
        policycount = policycount + 1;
    }
    if (backups != null){
        for (const c in backups){
            backupcount = backupcount + 1;
        }
    }
    console.log(ID);
  
    return(
        <div className="backup-item-grid-container">
        <p className="backup-label-1">Inc</p>
        <div className="backup-item-name-1">
    <p>{name}</p>
        </div>
        <div className="dropdown">
        <button className="backup-options">...</button>
        <div className="dropdowncontent">
            <Link to={`/policy/${id}`} className="dropdown-item">View Policies</Link>
            <Link to={`/backup/${id}`} className="dropdown-item">View Backups</Link>
            
            {/* <li>Edit Name</li>
            <li>Edit Directories</li>
            <li>Delete Client</li> */}
        </div>
        
        </div>
    <p className="backup-item-policy">Policy: {policycount}</p>
    <p className="backup-item-backup">Backups: {backupcount}</p>
    </div>  
    );
};

export default ClientItem;