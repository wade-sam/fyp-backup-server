import React from 'react';
import '../App.css';

const BackupItem = () =>{
    return(

        <div className="backup-item-grid-container">
            <p className="backup-label-1">Inc</p>
            <div className="backup-item-name-2">
                <p>08/04/2021</p>
            </div>
            <p className="backup-item-row-2">Friday's Backup</p>
            <div className="dropdown">
            <button className="backup-options">...</button>
            <div className="dropdowncontent">
                <li>View</li>
                <li>Restore</li>
                <li>Delete</li>
                <li>Edit Retention</li>
            </div>
            
            </div>
            <p className="backup-item-policy">Retention: 5 days</p>
            <p className="backup-item-backup">Failures: 6</p>
        </div>
    );
};

export default BackupItem;