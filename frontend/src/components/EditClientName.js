import React, {useEffect, useState} from 'react';
import '../App.css';

const EditClientName = (props) =>{

    const [client, setClient] = useState([]);
    const addNewClient = async(clientname) =>{
       // const data = await fetch('http://localhost:8000/api/clients/search', {method:"GET"});
       // const items = await data.json();
        console.log(props.clientid, clientname); 
        //setClient(items)
       
      }
      console.log("EDIT CLIENT")
    return(
        <div className="popup-box">
            <div className="box">
                <span className="close-icon" onClick={props.toggle}>x</span>
                <form>
                    <label>Enter Client Name</label>
                    <input type="text" id="username"></input>
                    <input type="submit" onClick={() => {
                         console.log(document.getElementById("#username")) 
                    }}/>
                    {/* <button onClick={() =>{
                       
                        // var id = document.querySelector('#username');
                        // addNewClient(id)
                        }}>Add</button> */}
                </form>
                {/* <p className="newclient-row1">Client Name</p> */}
               
               
            </div>
        </div>
    );
};

export default EditClientName