
import './App.css';
import PolicyItem from './components/policy-item'
import PolicyList from './PolicyList'
import React, {useState, useEffect, useContext}from 'react';
import {ClientContext, fetchAll} from './ClientContext';

function Policy({ match }) {
  console.log("HELLO")
  const {policies, setPolicies} = useContext(ClientContext);
  //const [policies, setPolicies] = useContext(ClientContext);
  useEffect(() =>{
    fetchAll();
  },[]);
 
  const fetchAll = async() =>{
    const data = await fetch('http://localhost:8000/api/policies/list');
    const items = await data.json();
    
    console.log(items)

   
   setPolicies(items)
  }

  if (match.params.id == "all"){
    console.log("ALL")

    return (
      <div className="grid-container">
          <input type="text" className="searchbox"></input>
          <PolicyList policies={policies} />
      </div>
    );
  }
  else{
    var policy = []
    for (const i in policies){
      console.log("hello", policies[i])
      for (const j in policies[i].clients){
        if (policies[i].clients[j][1] == match.params.id){
          policy.push(policies[i])
          console.log("FOUND", policies[i].clients[j][1])
          break;
          
        }
      }
    }
    return(
      <div className="grid-container">
        <input type="text" className="searchbox"></input>
        <PolicyList policies={policy} />
  </div>
    );
  }
  
}

export default Policy;
