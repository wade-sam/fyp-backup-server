import logo from './logo.svg';
import './App.css';
import PolicyItem from './components/policy-item'
import React, {useState, useEffect}from 'react';
function Policy({ match }) {
  useEffect(() =>{
    fetchAll();
  },[]);

  const [items, setItems] = useState([]);
  const [policy, setPolicy] = useState();
  const fetchAll = async() =>{
    const data = await fetch('http://localhost:8000/api/policies/list');
    const items = await data.json();
    console.log(items);
   // console.log(items.policies.length)
    setItems(items)
  }

  // const getItem = async() =>{
  //   const data = await fetch (`http://localhost:8000/api/policies/get/${id}`);
  //   const policy = await data.json();
  //   setPolicy = policy;
  // }

  if (match.params.id == "all"){
    console.log("ALL")
    return (
      <div className="grid-container">
          <input type="text" className="searchbox"></input>
          {items.map(policy =>(
            <PolicyItem key={policy.id} id={policy.id} name={policy.policyname} clients={policy.clients} type={policy.type} runs={policy.BackupRun}/>
      ))}
      </div>
    );
  }
  else{
    var policies = []
    for (const i in items){
      console.log("hello", items[i])
      for (const j in items[i].clients){
        if (items[i].clients[j][1] == match.params.id){
          policies.push(items[i])
          console.log("FOUND", items[i].clients[j][1])
          break;
          
        }
      }
    }
    return(
      <div className="grid-container">
      <input type="text" className="searchbox"></input>
      {policies.map(policy =>(
        <PolicyItem key={policy.id} id={policy.id} name={policy.policyname} clients={policy.clients} type={policy.type} runs={policy.BackupRun}/>
      ))}
  </div>
    );
  }
  
}

export default Policy;
