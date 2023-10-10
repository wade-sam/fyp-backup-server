import React, {useEffect, useState} from 'react';
import '../App.css';
import {ClientContext} from '../ClientContext';
const NewClient = (props) =>{
    const [editClient, setEditClient] = useState(false);
    const [newClientView, setNewClientView] = useState(true);
    const [editView, setEditView] = useState(false);
    const [empty, setEmpty] = useState(true);
    const editClientName = () =>{
        console.log("new client selected");
        
        setNewClientView(!newClientView);
        setEditView(!editView);
    }
    const createClientSubmitted = (e) =>{
        e.preventDefault();
        var bodymesage = {"consumerID":client, "clientname":change.clientname};
        console.log(bodymesage, JSON.stringify(bodymesage));
        createClient(bodymesage);
        
    }

    const createClient = async(bodymesage) =>{
        const data = await fetch('http://localhost:8000/api/clients/create', {method:"POST", body: JSON.stringify(bodymesage)});
        const items = await data.json();
        console.log(items)
    }

    const [change, setChange] = useState({
        "clientname":"",
       
    });

    const form = e =>{
        if (change.clientname.length == 0){
            if (empty == false){
                setEmpty(!empty);
            }
        }
        if(change.clientname.length > 0){
            if (empty == true){
                setEmpty(!empty);
            }
        }
        setChange({[e.target.name]:e.target.value});
        console.log(change, change.clientname.length);
    }

    const toggleNewEditPopup = (id) =>{
      
        setEditClient(!editClient);
        
      }


    const [client, setClient] = useState([]);
    const fetchNewClient = async() =>{
        const data = await fetch('http://localhost:8000/api/clients/search', {method:"GET"});
        const items = await data.json();
        console.log(items); 
        setClient(items)
       
      }
    return(
        <div className="popup-box">
            <div className="box">
                <span className="close-icon" onClick={props.toggle}>x</span>
                {newClientView && <p className="newclient-row1">New Clients ...</p>}
                {newClientView && <div className="list-new-clients-box">
                    <li onClick={editClientName}>{client}</li>
                    {/* {editClient && <EditClientName toggle={toggleNewEditPopup} id={client}/>} */}
                </div>}
                {editView && <form>
                    <label>Enter Client Name: </label>
                    <input name="clientname" type="text" id="username" value={change.clientname} onChange={e => form(e)}></input>
                    {/* <input type="submit" value="Add" onClick={createClient("hello")}/> */}
                <button onClick={e => createClientSubmitted(e)} disabled={empty}>Create</button>
                </form>}
                {newClientView && <button onClick={fetchNewClient}>Search</button>}
                {/* {editView && <button>Add</button>} */}
                
               
            </div>
        </div>
    );
};

export default NewClient