
import './App.css';
import Backup from './Backup.js';
import Client from './Client.js';
import Policy from './Policy.js';
import Nav from './nav.js';
import { BrowserRouter as Router, Switch, Route } from'react-router-dom';
function App() {
  return (
    <Router>
      <div className="App">
       <Nav />
       <Switch>
        <Route path="/client" component={Client} />
        {/* <Route path="/policy/" exact component={Policy}/> */}
        <Route path="/policy/:id" component={Policy} />
        <Route path="/backup" exact component={Backup} />
        <Route path="/backup/:id" component={Backup} />
       </Switch>
      </div>
        
    </Router>

  );
}

export default App;
