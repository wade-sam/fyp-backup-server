
import './App.css';
import Backup from './Backup.js';
import Client from './Client.js';
import Policy from './Policy.js';
import Nav from './nav.js';
import {PolicyProvider} from './PolicyContext'
import {ClientProvider} from './ClientContext'
import { BrowserRouter as Router, Switch, Route } from'react-router-dom';
import ClientItem from './components/client-item';
import ClientList from './ClientList';
function App() {
  return (
    
    <Router>
      <div className="App">
       <Nav />
       <Switch>
         <ClientProvider> <Route path="/client" component={Client} /> </ClientProvider>
       </Switch>
        {/* <Route path="/policy/" exact component={Policy}/> */}
        <Switch>
        <PolicyProvider>
          <Route path="/policy/:id" component={Policy} />
        </PolicyProvider>
        
        <Route path="/backup" exact component={Backup} />
        <Route path="/backup/:id" component={Backup} />
       </Switch>
      </div>
        
    </Router>
  

  );
}

export default App;
