import React, { useState } from 'react';
import './App.css';
import Backup from './Backup.js';
import Client from './Client.js';
import Policy from './Policy.js';
import Nav from './nav.js';
import {PolicyProvider} from './PolicyContext'
import {ClientProvider} from './ClientContext'
import { BrowserRouter as Router, Switch, Route } from'react-router-dom';
import NewClient from './components/NewClient';
import EditClientName from './components/EditClientName';
import ClientsPage from './ClientPage'
import PoliciesPage from './Policypage';
import BackupPage from './components/BackupPage';
function App() {

  const [isOpen, setIsOpen] = useState(false);
  const [editClient, setEditClient] = useState(false);
  const [clientid, setClientid] = useState();
  const toggleNewPopup = () => {
      setIsOpen(!isOpen);
  }
  const toggleNewEditPopup = (id) =>{
    setEditClient(!isOpen);
    setClientid(id);

  }
  return (
    
    <Router>
      {/* <div className="App">
       <Nav newClient={toggleNewPopup}/>
       {isOpen && <NewClient toggle={toggleNewPopup} editclient={toggleNewEditPopup}/>}
       {editClient && <EditClientName toggle={toggleNewEditPopup} id={clientid}/>}
       <Switch>
         <ClientProvider> <Route path="/client" component={Client} /> </ClientProvider>
       </Switch>
        {/* <Route path="/policy/" exact component={Policy}/> */}
        {/* <Switch>
        <PolicyProvider>
          <Route path="/policy/:id" component={Policy} />
        </PolicyProvider>
        
        <Route path="/backup" exact component={Backup} />
        <Route path="/backup/:id" component={Backup} />
       </Switch> */}
       <div className="App">
       <Nav newClient={toggleNewPopup}/>
       <Switch>
         <Route path="/client" component={ClientsPage}></Route>
         <Route path="/policy" component={PoliciesPage}></Route>
         <Route path="/policy/:id" component={PoliciesPage} />
        <Route path="/backup" component={BackupPage} />
       </Switch>
      
      </div>
        
    </Router>
  

  );
}

export default App;
