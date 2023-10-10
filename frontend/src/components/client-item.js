import React, {Component, useContext, useState} from 'react';
import '../App.css';
import { Link } from 'react-router-dom'
import Client from '../Client';

import ClientOptions from './clientoptions';

class ClientItem extends React.Component{ 
    constructor(props){
        super(props);
        this.props = props;
        this.state = {hide: false};
    }   

    handleNewClientClick(){
        this.setState({hide: true});
    }


   
    
    render(){
        const {hide} = this.state;
        if (hide){
            return null;
        }
      //  const [clients, setClients] = this.context;
        const {clients, setClients} = this.context;
        
        const deleteClient = () =>{

        const index = clients.indexOf(this.props.client);

        console.log(index);
  
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
                <ClientOptions client={this.props.client} popup = {this.props.toggle} deleteOnClick={this.handleNewClientClick.bind(this)}/>
                <p className="backup-item-policy">Policy: {policycount}</p>
                <p className="backup-item-backup">Backups: {backupcount}</p>
            </div>  
        );
    }
}


export default ClientItem;