import React, {Component, useContext, useState} from 'react';
import '../App.css';
import { Link } from 'react-router-dom'
import Client from '../Client';
import {ClientContext} from '../ClientContext';
import ClientOptions from './clientoptions';
//import useVisibilityCard from './useVisibilityCard';
 

// function deleteClient(id){
//     console.log(`http://localhost:8000/api/clients/delete/${id}`)
//     fetch(`http://localhost:8000/api/clients/delete/${id}`, {method:"DELETE"}).then(response => {
//         console.log(response)
//     })
//   //  Client.set

// }
class ClientItem extends React.Component{ 
    constructor(props){
        super(props);
        this.props = props;
        this.state = {hide: false};
    }   
    handleClick(){
        this.setState({hide: true});
    }
    static contextType = ClientContext
    
    render(){
        const {hide} = this.state;
        if (hide){
            return null;
        }
        const [clients, setClients] = this.context

      //  const [clients, setClients] = useContext(ClientContext);
        const deleteClient = () =>{
       // console.log(`http://localhost:8000/api/clients/delete/${props.client.id}`)
        const index = clients.indexOf(this.props.client);
        // if (index > -1){
        //     console.log(clients);
            
        //   //  setClients.splice(index, 1);
        //     console.log(clients);
        // }
        console.log(index);
        // fetch(`http://localhost:8000/api/clients/delete/${props.client.id}`, {method:"DELETE"}).then(response => {
        //     console.log(response)
        // })
    }
        var ID = this.props.client.id;
        var policycount = 0;
        var backupcount = 0;
        for (const c in this.props.client.policies){
            policycount = policycount + 1;
        }
        if (this.props.client.backups != null){
            for (const c in this.props.client.backups){
                backupcount = backupcount + 1;
            }
        }
        console.log(ID);
        return(
            <div className="backup-item-grid-container">
                <p className="backup-label-1">Inc</p>
                <div className="backup-item-name-1">
                <p>{this.props.client.clientname}</p>
                </div>
                <ClientOptions id={this.props.client.id} onClick={this.handleClick.bind(this)}/>
                <p className="backup-item-policy">Policy: {policycount}</p>
                <p className="backup-item-backup">Backups: {backupcount}</p>
            </div>  
        );
    }
}
// const ClientItem = (props) =>{
    
//     const [clients, setClients] = useContext(ClientContext);
//     const deleteClient = () =>{
//        // console.log(`http://localhost:8000/api/clients/delete/${props.client.id}`)
//         const index = clients.indexOf(props.client);
//         if (index > -1){
//             console.log(clients);
            
//           //  setClients.splice(index, 1);
//             console.log(clients);
//         }
//         console.log(index);
//         // fetch(`http://localhost:8000/api/clients/delete/${props.client.id}`, {method:"DELETE"}).then(response => {
//         //     console.log(response)
//         // })
//     }
//     var ID = props.client.id;
//     var policycount = 0;
//     var backupcount = 0;
//     for (const c in props.client.policies){
//         policycount = policycount + 1;
//     }
//     if (props.client.backups != null){
//         for (const c in props.client.backups){
//             backupcount = backupcount + 1;
//         }
//     }
//     console.log(ID);
  
//     return(
//         <div className="backup-item-grid-container">
//         <p className="backup-label-1">Inc</p>
//         <div className="backup-item-name-1">
//     <p>{props.client.clientname}</p>
//         </div>
//         <div className="dropdown">
//         <button className="backup-options">...</button>
//         <div className="dropdowncontent">
//             <Link to={`/policy/${props.client.id}`} className="dropdown-item">View Policies</Link>
//             <Link to={`/backup/${props.client.id}`} className="dropdown-item">View Backups</Link>
//             <li onClick={() => deleteClient(props.client.id)} className="dropdown-item">Delete Client</li>
//             {/* <li>Edit Name</li>
//             <li>Edit Directories</li>
//             <li>Delete Client</li> */}
//         </div>
        
//         </div>
//     <p className="backup-item-policy">Policy: {policycount}</p>
//     <p className="backup-item-backup">Backups: {backupcount}</p>
//     </div>  
//     );
// };

export default ClientItem;