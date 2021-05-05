import React, {useState, useReducer, useContext, useEffect} from 'react';
import CheckboxTree from 'react-checkbox-tree';
//import "react-checkbox-tree/src/less/react-checkbox-tree.less";
import 'react-checkbox-tree/lib/react-checkbox-tree.css';
export function clientsReducer(state, action){
    switch(action.type){
        case 'hideClient':{
            const client = state.clients.find(item => item.ID === action.payload);
            console.log(client);
            client.visible = false;
            return {
                ...state,
                
            };
        }
        case 'isLoaded':{
            return{
                ... state,
                isLoading:false,
            }
        }
        case 'addClients':{
            
            return{
                ...state,
                clients:action.payload,
            }
        }
        case 'newClient':{
            return{
                ...state, 
                newClient:true,
            }
        }
        case 'popupfinished':{
            return{
                ...state,
                newClient:false,
                editScan: false,
            }
        }
        case 'directoryScanFinished':{
            return{
                ...state, 
                editScan: false,
            }
        }

        case 'selectClientForScan':{
            const clientFound = state.clients.find(item => item.ID === action.payload);
            console.log(clientFound);
            return{
                ...state,
                clientSelected: clientFound,
                editScan:true,
            }
        }
        
        default:
            break;
    }
    return state;
}


const fetchItems = fetch('api/clients/list', {method:"GET"}, {headers : { 
    'Content-Type': 'application/json',
    'Accept': 'application/json'
   }}).then((response) => response.json()).then((clients) =>{return clients})
export const ClientStateContext = React.createContext();
const ClientDispatchContext = React.createContext();
function setClientState (client){
    const returnClient = {
        ID: client.id,
        Name: client.clientname,
        ConsumerID: client.consumerID,
        Policies: client.policies,
        visible: true,
        ignoreList: client.ignorepath,
        backups: client.backups
    }
    return returnClient
}

const initialState = {
    isLoading: true,
    error: '',
    newClient:false,
    editScan:false,
    clientSelected:{},
    clients: []
};

export default function ClientsPage() {
    useEffect(() =>{
        fetchItems();

    },[]);


    const fetchItems = async() =>{
        const data = await fetch('/api/clients/list', {method:"GET"}, {headers : { 
            'Content-Type': 'application/json',
            'Accept': 'application/json'
           }});
        const items = await data.json();
        console.log(typeof items); 
        var clients = [];
        items.map(client =>{
            var returnClient = setClientState(client);
            clients.push(returnClient);
            
        })
        dispatch({type:'addClients', payload:clients})
       // initialState.clients.push(clients)
        console.log(initialState.clients);
      }
      
      //cancel newClient
      const toggleNewPopup = () => {
          dispatch({type:"popupfinished"});
          fetchItems();
      //  state.newClient = !state.newClient;
    }
    const [state, dispatch] = useReducer(clientsReducer, initialState)
    const {Name, ConsumerID, Policies, visible, ignoreList, backups} = state.clients;

    const newDirectoryScanTrigger = (id) =>{
        dispatch({type:"selectClientForScan", payload:id})
    }
    //dispatch({ type: 'fetch'});
    console.log(initialState);
    return(
        <ClientDispatchContext.Provider value ={dispatch}>
            <ClientStateContext.Provider value={state}>
                <div className="grid-container">
                    <div className="searchbox">
                  
                        <input type="text" ></input>
                        <button className="newClient" onClick={() =>{dispatch({type:"newClient"})}}>New</button>
                        
                    </div>
                   
                
                    
                    <div className="grid-container-list">
                    {state.newClient && <NewClient close={toggleNewPopup}/>}
                    {state.editScan && <EditDirectoryScan close={toggleNewPopup} client={state.clientSelected}/>}
                    {state.clients.map(item =>(

                        item.visible && <ClientBox key={item.ID} id={item.ID} consumerID={item.ConsumerID} name={item.Name} policies={item.Policies} runs={item.backups} visible={item.visible}></ClientBox>
                    ))}
                    </div>
                </div>
            </ClientStateContext.Provider>
        </ClientDispatchContext.Provider>
        
    );
}
//export default ClientsPage;

function ClientBox({id, consumerID, name, policies, runs, visible}){
    // console.log(id);
    var plength = 0;
    var blength = 0;
    if (policies != null){
        plength = policies.length;
    }
    if (runs != null){
        blength = runs.length;
    }
    if (visible == false){
        return null;
    }
    return(
        <div className="backup-item-grid-container">
            {/* <h1>Hello</h1> */}
             <p className="backup-label-1">Inc</p>
                <div className="backup-item-name-1">
                <p>{name}</p>
                </div>
                <ClientOptions id={id}/>
                <p className="backup-item-policy">Policy: {plength}</p>
                <p className="backup-item-backup">Backups: {blength}</p>

        </div>
    )
}



function newClientReducer(state, action){
    switch (action.type){
        case 'searching':{
            return{
                ...state,
                isSearching:true,
            }
        }
        case 'foundClient':{
            return{
                ...state,
                isFoundClient:true,
                isSearching: false,

                consumerID: action.payload,
            }
        }
        case 'noNewClient':{
            return{
                ...state,
                error:"No new client found",
                close: true,
                isSearching:false,
            }
        }
        case 'field':{
            if(action.value.length > 4){
                return{
                    ...state,
                    [action.field]: action.value,
                    isSubmit:false,
                }
            }else{
                return{
                    ...state,
                    [action.field]: action.value,
                    isSubmit:true,
                }
            }
            
        }
        case 'canceled':{
            return{
                ...state,
                isSearching: false,
                isFoundClient:false,
                isComplete:false,
                isSubmit:true,
                clientname:"",
                consumerID:"",
                
            }
        }
        case 'createSuccess':{
            return{
                ... state,
                isSearching: false,
                isFoundClient:false,
                isComplete:false,
                isSubmit:true,
                clientname:"",
                consumerID:"",
            }
        }
        case 'createFail':{
            return{
                ...state,
                error:"Couldn't Create Client. Please try again",
                clientname:"",
            }
        }
        default:break;
    }
}
const newClientInitialState = {
    close: false,
    error: "",
    isSearching: false,
    isFoundClient: false,
    isComplete: false,
    isSubmit:true,
    clientname: "",
    consumerID: ""
}



function NewClient(props){
   

    const [state, dispatch] = useReducer(newClientReducer, newClientInitialState);
    const {close, error, isSearching, isFoundClient, isComplete, isSubmit, clientname, consumerID} = state;
    
    useEffect(() =>{
        searchForNewClient();

    },[]);

    const searchForNewClient = async() =>{
        dispatch({type:"searching"})
        try{
            const data = await fetch('api/clients/search', {method:"GET"}, {headers : { 
                'Content-Type': 'application/json',
                'Accept': 'application/json'
               }});
            const items = await data.json();
            dispatch({type:"foundClient", payload:items})
            console.log(consumerID);
        }catch{
            dispatch({type:"noNewClient"})
        }
        
      }
      const onSubmit = async e =>{
        e.preventDefault();
        var bodymesage = {"consumerID":consumerID, "clientname":clientname};
        console.log(bodymesage, JSON.stringify(bodymesage));
        try{
            const data = await fetch('/api/clients/create', {method:"POST", body: JSON.stringify(bodymesage)});
            const items = await data.json();
            dispatch({type:"createSuccess"});
            props.close();
        }catch{
            dispatch({type:"createFail"});
        }
        
      }


    return(
        <div className="popup-box">
            <div className="box">
                <span className="close-icon" onClick={props.close}>x</span>
                 
                {isSearching ? (
                    <>
                        <h1>Searching</h1>
                    </>
                ):(
                    <>
                        <form>
                            <label>Enter Client Name</label>
                            <label>{consumerID}</label>
                            <h1>Hello</h1>
                            <input type="text" placeholder="clientname" value={clientname} onChange={e => dispatch({type: 'field', field:'clientname', value: e.currentTarget.value})}></input>
                            <button disabled={isSubmit} type="submit" onClick={onSubmit}>Create</button>
                            {<label>{error}</label>}
                        </form>
                    </>
                )}
            </div>
        </div>
        //<h1>New Client</h1>
    )
}

function directoryScanReducer(state, action){
    switch(action.type){
        case 'scanning':{
            return{
                ...state,
                isRequested: true,
            }
        }
        case 'scanComplete':{
            return{
                ...state,
                isRequested:false,
                isScanComplete: true,
                scan: [...state.scan, action.payload],
            }
        }
        case 'onCheck':{
                return{
                    ...state,
                    checked:action.payload,
                }
        }
        case 'onExpand':{
            
            return{
                ...state,
                    expanded:action.payload,
                
            }

        }
        case 'isComplete':{
            if (state.isComplete == true){
                console.log('isSave is true');
                return{
                    ...state, 
                    isComplete:false,
                }
            }else{
                console.log('isSave is false');
                return{
                    ...state, 
                    isComplete:true,
                }
            }
        }
       
        default:
            return
    }
}


const DscanInititalState = {
    close: false,
    error: "",
    isRequested: false,
    isScanComplete: false,
    isComplete: true,
    isSave:false,
    checked: [],
    expanded: [],
    scan:[],
    //ignoreList: [],
}

function EditDirectoryScan({close, client}){
    console.log(client);
    const consumerID = client.ConsumerID;
    const [state, dispatch] = useReducer(directoryScanReducer, DscanInititalState);
    const {error, isRequested, isScanComplete, isComplete, isSave,checked, expanded, scan} = state;
    useEffect(() =>{
        getScan();

    },[]);
    var dat = [];
    const getScan = async() =>{
        console.log("CLIENT", consumerID);
        dispatch({type:"scanning"})
        try{
            const data = await fetch(`/api/clients/scan/${consumerID}`, {method:"GET"}, {headers : { 
                'Content-Type': 'application/json',
                'Accept': 'application/json'
               }});
            const t = await data.json();
          
            dispatch({type:"scanComplete", payload:t})
            
            console.log(scan);
        }catch{
            console.log("ERROR");
        }
        

    }
    const updateClient = async(newList) =>{
        client.ignoreList = newList;
        var bodymesage = {"id":client.ID,"consumerID":client.ConsumerID, "clientname":client.Name, "policies":[], "tree": null, "ignorepath":client.ignoreList, "backups":client.backups};
        console.log(JSON.stringify(bodymesage));
        try{
            const data = await fetch(`/api/clients/update`, {method:"PUT", body:JSON.stringify(bodymesage)}, {headers : { 
                'Content-Type': 'application/json',
                'Accept': 'application/json'
               }});
            const t = await data.json();
            console.log(t);            
        }catch{
            
        }
        dispatch({type:"isComplete"});
       // close();
        
    }
    const handleSubmit  = (e) =>{
       // dispatch({type:"isSave"});
       var queue = []
        var ignoreList = [];
        e.preventDefault();    
        queue.push(scan[0]);
        while (!queue.length == 0){
            let visiting = queue.shift()
            if(visiting.children != null){
                for (let j = 0; j < visiting.children.length; j++){
                    if(checked.indexOf(visiting.children[j].value) != -1){
                        ignoreList.push(visiting.children[j].value)
                    }else if(checked.indexOf(visiting.children[j].value) == -1){
                        queue.push(visiting.children[j]);
                    }
                }
            }

        }
       // dispatch({type:"isSave"});
        updateClient(ignoreList);
    }


    //handle the selecting of checkbox, used to mverify whether the user can save. 
    const handleCheck = (checked) =>{
        dispatch({type:"onCheck", payload:checked});
      //  dispatch({type: "isSave"});
    }

    return(
        <div className="popup-box">
            <div className="box">
                <span className="close-icon" onClick={close}>x</span>
                {/* {isRequested && <h1>Requesting Scan</h1>} */}
                {isRequested ? (
                <>
                    <h1>Requesting Scan</h1>
                </>
                ):(
                    
                    <>
                    {isComplete ? (<h2>Select Directories</h2>):(
                        <h2>Success!!</h2>
                    )}
                    <p>Please select the directories {"\n"}you would like to be excluded</p>
                    <CheckboxTree
                        nodes={scan}
                     //   key={scan.value}
                        checked={checked}   
                        checkModel={'all'}
                       // noCascade={true}                      
                        expanded={expanded}
                        nativeCheckboxes={true}
                        onCheck={checked => handleCheck(checked)}
                        onExpand={expanded => dispatch({type:"onExpand", payload:expanded})}
                    />
                    <button onClick={handleSubmit}>{isSave ? 'Saving...': 'Save'}</button>
                        
                    </>
                   
                    
                    
                    
                )}
                
            </div>
        </div>
    )
}

function ClientOptions({id}){
    const dispatch = useContext(ClientDispatchContext);
    const deleteClient = () =>{
        fetch(`/api/clients/delete/${id}`, {method:"DELETE"}, {headers : { 
            'Content-Type': 'application/json',
            'Accept': 'application/json'
           }}).then(response => {
           console.log(response);
       
        })
    }
          //cancel newClient
          const toggleNewPopup = () => {
            dispatch({type:"directoryScanFinished"});
            fetchItems();
        //  state.newClient = !state.newClient;
      }
    //console.log(id);
    return(
        <div className="dropdown">
            <button className="backup-options">...</button>
            <div className="dropdowncontent">
            {/* <Link to={`/policy/${id}`} className="dropdown-item">View Policies</Link>
            <Link to={`/backup/${id}`} className="dropdown-item">View Backups</Link> */}
            <li className="dropdown-item" onClick={() =>{
                dispatch({type:"selectClientForScan", payload:id})
            }} >View/Edit Directories</li>
            <li className="dropdown-item">Edit Client Name</li>
            <li className="dropdown-item" onClick={() =>{
                deleteClient()
                dispatch({type: 'hideClient', payload:id})
            }}>Delete Client</li>
            {/* <li onClick={deleteClient} className="dropdown-item">Delete Client</li> */}
            
            </div>
    
        </div>
    );
}

//export default ClientBox;