import React, {useState, useReducer, useContext, useEffect} from 'react';

import TimePicker from "rc-time-picker";
import 'rc-time-picker/assets/index.css';
import moment from "moment";
import {ClientStateContext} from './ClientPage';
import { Multiselect } from 'multiselect-react-dropdown';
import { act } from 'react-dom/test-utils';

function policiesReducer(state, action){
    switch(action.type){

        case 'hidePolicy':{
            const policy = state.policies.find(item => item.ID === action.payload);
            console.log(policy);
            policy.visible = false;
            return{
                ...state,
            }
        }
        case 'addPolicies':{
            return{
                ...state,
                policies:action.payload,
            }
        }
        case 'newPolicy':{
            return{
                ...state,
                newPolicy:true,
            }
        }
        case 'selectPolicyForEditing':{
            const policyfound = state.policies.find(item => item.ID === action.payload);
            console.log(policyfound);
            return{
                ...state,
                policySelected: policyfound,
                editPolicy:true,
            }
        }
        case 'popupfinished':{
            return{
                ...state,
                newPolicy:false,
                editPolicy:false,
            }
        }

        default:
            break;
    }
}



//const fetchItems = fetch('http://localhost:8000/api/policies/list', {method:"GET"}).then((response) => response.json()).then((clients) =>{return clients})
const PolicyStateContext = React.createContext();
const PolicyDispatchContext = React.createContext();
function setPolicyState (policy){
    const returnPolicy = {
        ID: policy.id,
        Name: policy.policyname,
        RuntTime: policy.RuntTime,
        Clients: policy.clients,
        visible: true,
        retention:policy.retention,
        type:policy.type,
        fullbackup:policy.fullbackup,
        BackupRun:policy.BackupRun,
    }
    return returnPolicy
}

const initialState = {
    isLoading: true,
    error: '',
    newPolicy:false,
    editPolicy:false,
    policySelected:{},
    policies: []
};

export default function PoliciesPage({ match }){

    

    useEffect(() =>{
        fetchItems();

    },[]);

    const fetchItems = async() =>{
        const data = await fetch('/api/policies/list', {method:"GET"});
        const items = await data.json();
        console.log(typeof items); 
        var policies = [];
        items.map(policy =>{
            var returnClient = setPolicyState(policy);
            policies.push(returnClient);
            
        })
        dispatch({type:'addPolicies', payload:policies})
        console.log(initialState.policies);
      }
      //hide Popup
      const toggleNewPopup = () => {
        dispatch({type:"popupfinished"});
        fetchItems();
    //  state.newClient = !state.newClient;
    }
    

      const [state, dispatch] = useReducer(policiesReducer, initialState);
      return(
        <PolicyDispatchContext.Provider value={dispatch}>
            <PolicyStateContext.Provider value={state}>
            <div className="grid-container">
                <div className="searchbox">
                            <input type="text" ></input>
                            <button className="newPolicy" onClick={() =>{dispatch({type:"newPolicy"})}}>New</button>
                </div>
                <div className="grid-container-list">
                    
                    {state.newPolicy && <NewPolicy close={toggleNewPopup}/>}
                    {state.editPolicy && <EditPolicy close={toggleNewPopup} policy={state.policySelected}/>}
                    {state.policies.map(item =>(
                        item.visible && <PolicyBox key={item.ID} id={item.ID} name={item.Name} clients={item.Clients} runs={item.BackupRun} visible={item.visible}/>
                    ))}
                    
                    
                </div>
            </div>
            </PolicyStateContext.Provider>
        </PolicyDispatchContext.Provider>
      )

}

function PolicyBox({id, name, clients, type, runs, visible}){
    // console.log(id);
    var clength = 0;
    var blength = 0;
    if (clients != null){
        clength = clients.length;
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
             <p className="backup-label-1">{type}</p>
                <div className="backup-item-name-1">
                <p>{name}</p>
                </div>
                <PolicyOptions id={id}/>
                <p className="backup-item-policy">Clients: {clength}</p>
                <p className="backup-item-backup">Runs: {blength}</p>

        </div>
    )
}

function newPolicyReducer(state, action){
    switch(action.type){
        case'field':{
            return{
                ...state,
                [action.field]:action.value,
                isModified:true
            }
        }
        case 'isUpdatePickerInitialised':{
            console.log("isUpdatePickerInitialised", action.clients, action.rundays);
            return{
                ...state,
                newBackuprundays: action.rundays,
                clients: action.clients,
                isUpdatePickerInitialised: true,
            }
        }
        case 'fieldDays':{
            console.log("days payload: ",action.payload);
            return{
                ...state,
                newBackuprundays:action.payload,
                isModified:true,
                isDaysModified: true,
            }
        }
        case 'fieldClients':{
            console.log(state.clients)
            return{
                ...state, 
                clients:action.value,
                isModified:true,
                isClientsModified:true,
            }
        }
        case 'setTime':{
            return{
                ...state, 
                backupruntime:action.value,
                isModified:true
            }
        }
        
        case 'searchClients':{
            return{
                ...state,
                [action.field]:action.value,
                isModified:true
            }
        }
        case 'setBackupDays':{
            return{
                ...state,
                oldbackuprundays:action.value,
            }
        }
        case 'isClients':{
            return{
                ...state,
                isClients:true,
                isModified:true
            }
        }
        case 'isComplete':{
            return{
                ...state, 
                isNotComplete: true,
                
            }
        }
        case 'isEmpty':{
            return{
                ...state,
                isNotComplete: false,
            }
        }
        case 'isFormComplete':{
            if (state.policyname == "" || state.clients.length == 0 || state.newBackuprundays.length ==0 || state.backupruntime == "" || state.retention == ""){
                console.log("FIELD'S ARE EMPTY", "policy: ", state.policyname, "clients: ", state.clients, "run: ", state.newBackuprundays, "runtime: ", state.backupruntime, "retention: ", state.retention);
                return{
                    ...state,
                    isNotComplete: true,
                }
            }else{
                if (action.value == "PUT" && state.isModified == false){
                    return{
                        ...state,
                        isNotComplete: true,
                    }
                }else{
                    return{
                        ...state,
                        isNotComplete: false,
                    }
                }
                
            }
        }
        case 'editStart':{
            return{
                ...state,
                isEditied:true,
            }
        }
 

        default:
            break;
    }
}

const newPolicyInitialState = {
    error: "",
    isModified: false,
    isDaysModified: false,
    isClientsModified: false,
    isNotComplete: true,
    isUpdatePickerInitialised:false,
    isDate:false,
    isRunTime:false,
    isClients:false,
    isEditied:false,
    policyname: "",
    oldclients:[],
    clients:[],
    //backuprun:[],
    backupruntime:"",
    oldbackuprundays:[],
    newBackuprundays:[],
    retention:"",
    clientsplaceholder:[]
}



//API type means POST or PUT, depending on who is calling
function PolicyConfigurer({clientsList, apiType, apiCall, close}){
    useEffect(() =>{
        // clientsHolder();
        initialiseUpdate();

    },[]);
    const [state, dispatch] = useReducer(newPolicyReducer, newPolicyInitialState);
    const {isModified, isDaysModified, isClientsModified, isNotComplete, isUpdatePickerInitialised, policyname, oldclients, clients, backuprun, backupruntime, oldbackuprundays, newBackuprundays, retention} = state;
    

    //initialise reducer using newPolicyInitialState
    
    
    
    var clientsPlaceholder = clientsList;



    //  console.log(typeof clientsListGetter, " HOLDER ");
    //  var count = 1;
    //  clientsListGetter.map(client =>{
    //      var picker = {name: client.clientname, id: count};
    //      console.log("hello cunt",picker);
    //      clientsPlaceholder.push(picker);
    //      count ++;
    //  })

     console.log(clientsPlaceholder, " HOLDER 2");
    
    //using moment.js for formatting the time

    const initialiseUpdate = () =>{
        
        if(isUpdatePickerInitialised == false && apiType == "PUT"){
            console.log("reached");
            dispatch({type:'isUpdatePickerInitialised', clients:oldclients, rundays:oldbackuprundays});
            
        }
        
    }
    var time = null;
    
    if (backupruntime == ""){
        // time = moment("10:05", 'HH:mm');
        time = null;
        console.log("INITIALISE TIME EMPTY",time);
    }
    else{
        time = moment(backupruntime, 'HH:mm');
        console.log("INITIALISE TIME UPDATED",time);
    }

    // onClick() update clients state
    const onSelectClient = (selectList, selectedItem) =>{
        var dispatchList = [];
        selectList.map(day =>{
            dispatchList.push(day.name);
        });
        dispatch({type:'fieldClients', field:'clients', value: dispatchList});
        dispatch({type:'isFormComplete', value:apiType});
    }

    // onClick() update backuprundays state
    const onSelectDays = (selectList, selectedItem) =>{
        var dispatchList = [];
        selectList.map(day =>{
            dispatchList.push(day.name);
        });
        console.log(selectList);
        dispatch({type:'fieldDays', payload: dispatchList});
        dispatch({type:'isFormComplete', value:apiType});
        
    }

    // onClick() update backupruntime state
    const onTimeChange = (value) =>{
        if (value == null){
            console.log("DELETED");
            dispatch({type:'setTime', value:""});
            dispatch({type:'isFormComplete', value:apiType});
            time = value;
            console.log("MODIFIFED TIME DELETED",backupruntime);
           // test = moment("10:05", 'HH:mm');
            return;
        }
      // console.log(value.format('HH:mm'));
        var test = value.format('HH:mm');
        dispatch({type:'setTime', value:test});
        dispatch({type:'isFormComplete', value:apiType});
        time = value;
        console.log("MODIFIED TIME UPDATED",time);
    }


    const onSubmit = (e) =>{
        e.preventDefault();
        var bodymesage = {"policyname":policyname, "clients":clients, "runtime":backupruntime, "retention":parseInt(retention), "type":"full", "fullbackup":newBackuprundays, "incbackup":[], "BackupRun":[]};
        apiCall(bodymesage);
      }

    const onFieldChange = (field, value) =>{
        dispatch({type:'field',field:field, value:value });
        dispatch({type:'isFormComplete', value:apiType});

    }



    //placeholder for the days selection picker
    var daysPlaceholder = [{name: "Monday", id:1}, {name: "Tuesday", id:2},{name: "Wednesday", id:3},{name: "Thursday", id:4},{name: "Friday", id:5},{name: "Saturday", id:6},{name: "Sunday", id:7}];
    
    return(
        <div className="popup-box">
            
            <div>
                
                <form className="newPolicy-box">
                <span className="close-icon" onClick={close}>x</span>
                <input className="newPolicy-box-title"type="text" placeholder="Enter Policy Name" value={policyname} onChange={ e => onFieldChange('policyname', e.currentTarget.value)}></input>
                <div className="newPolicy-box-retention">
                    <label>Retention Time</label>
                <input  type="number" placeholder="10" value={retention} min={1} max={365} onChange={e => onFieldChange('retention', e.currentTarget.value)}></input>
                </div>
                <div className="newPolicy-box-clients">
                <label>Clients</label>
                <Multiselect options={clientsPlaceholder} selectedValues={oldclients} displayValue="name" onSelect={onSelectClient} onRemove={onSelectClient}></Multiselect>
                </div>
                <div className="newPolicy-box-day-schedule">
                    <label>Select Days</label>
                    <Multiselect options={daysPlaceholder} selectedValues={oldbackuprundays} displayValue="name" onSelect={onSelectDays} onRemove={onSelectDays}>
                    </Multiselect>
                </div>
                <div className="newPolicy-box-time-schedule">
                    <label>Start Time</label>
                    <TimePicker className="" showSecond={false} focusOnOpen={true} format="HH:mm" minuteStep={5} onChange={onTimeChange} value={time}></TimePicker>
                </div>
                <button className="newPolicy-box-submit" disabled={isNotComplete} onClick={onSubmit}>Add</button>
                </form>
            </div>
        </div>
    )

}
const editPolicyInitialState = {
    error: "",
    isModified: false,
    isDaysModified: false,
    isClientsModified: false,
    isNotComplete: true,
    isUpdatePickerInitialised:false,
    isDate:false,
    isLoaded:false,
    isRunTime:false,
    isClients:false,
    isEditied:false,
    policyname: "",
    oldclients:[],
    clients:[],
    //backuprun:[],
    backupruntime:"",
    oldbackuprundays:[],
    newBackuprundays:[],
    retention:"",
    clientsplaceholder:[]
}

function editPolicyReducer(state, action){
    switch (action.type){
        case "initialiseclients":{
            return{
                ...state,
                //oldbackuprundays:action.value,
                oldclients:action.value,
            }
        }
        case 'fieldClients':{
            return {
                ...state,
                clients:action.value,
            }
        }
        case'field':{
            console.log(state.policyname)
            return{
                ...state,
                [action.field]:action.value,
                isModified:true
            }
        }
        case 'isLoaded':{
            return{
                ...state,
                isLoaded:true,
            }
        }
        default:
            break;
    }
}

function EditPolicyConfigurer({clientsList, close, policy, sdays, sclients}){
    const [state, dispatch] = useReducer(editPolicyReducer, editPolicyInitialState);
    const oldClients = sclients;
    useEffect(() =>{
        if(state.isLoaded == false){
            onStartup();
            dispatch({type:"isLoaded"});
        }
        // clientsHolder();
        

    },[]);
   
    


    const onStartup = () =>{

        var tempHolder = []
        var count = 1;
        if (policy.Clients != null){
            policy.Clients.map(client => {
                console.log("client",client);
                var temp = {name:client[0], id:count};
                tempHolder.push(temp);
                
               //editPolicyInitialState.oldclients.push(temp);
               console.log(temp);
                count ++;
            })
        }
        console.log("temp", oldClients);
        dispatch({type:'field', field:"policyname", value:policy.Name});
        dispatch({type:'initialiseclients', value:tempHolder});
        console.log(editPolicyInitialState.clientsplaceholder);
        console.log(editPolicyInitialState.oldbackuprundays, editPolicyInitialState.oldclients)
        
        //var tempHolder = []
        // if (policy.Clients != null){
        //     policy.Clients.map(client => {
        //         console.log("client",client);
        //         var temp = {name:client[0], id:count};
        //         //tempHolder.push(temp);
        //        editPolicyInitialState.oldclients.push(temp);
        //         count ++;
        //     })
        // }
        //oldClients.push(tempHolder)
    }
    console.log(policy);
    const onFieldChange = (field, value) =>{
        dispatch({type:'field',field:field, value:value });
       // dispatch({type:'isFormComplete', value:apiType});

    }
    const onSelectClient = (selectList, selectedItem) =>{
        var dispatchList = [];
        selectList.map(day =>{
            dispatchList.push(day.name);
        });
        dispatch({type:'fieldClients', field:'clients', value: dispatchList});
       // dispatch({type:'isFormComplete', value:apiType});
    }

   return(
    <div className="popup-box">
            
    <div>
        
        <form className="newPolicy-box">
        <span className="close-icon" onClick={close}>x</span>
        <input className="newPolicy-box-title"type="text" placeholder="Enter Policy Name" value={state.policyname} onChange={e => onFieldChange('policyname', e.currentTarget.value)}></input>
        <div className="newPolicy-box-clients">
                <label>Clients</label>
        <Multiselect options={editPolicyInitialState.clientsplaceholder} selectedValues={oldClients} displayValue="name" onSelect={onSelectClient} onRemove={onSelectClient}></Multiselect>
       </div>
       </form>
    </div>
    </div>
   )

}

function EditPolicy({close, policy}){
    // const [state, dispatch] = useReducer(newPolicyReducer, newPolicyInitialState);
    // const {isModified, isDaysModified, isClientsModified,isEditied, isNotComplete, isUpdatePickerInitialised, policyname, oldclients, clients, backuprun, backupruntime, oldbackuprundays, newBackuprundays, retention} = state;
    const selectedPolicy = policy;
    console.log("selected policy", selectedPolicy);
    useEffect(() =>{
    //     if (isEditied == false){
    //         console.log("reached");
            listClients();
            prepInitialPolicyState();
    //         dispatch({type:"editStart"});
    //     }

    })
    var clientsPlaceholder = [];
    // const [Clients, setClients] = useState([]);
    const listClients = async() =>{
    //     //console.log(clientsplaceholder.length);
        
            const data = await fetch('/api/clients/list', {method:"GET"});
            const items = await data.json();
            //clients = items;
            console.log(typeof items); 
            var policies = [];
            var count = 1;
           // setClients(items);
            items.map(client =>{
                var picker = {name: client.clientname, id: count};
                editPolicyInitialState.clientsplaceholder.push(picker);
                var c = client;
                
               // clients.push(c);
                count ++;
            })
            //editPolicyInitialState.oldclients = clientsPlaceholder
    //         console.log(clientsPlaceholder);       
    }
    // var prepDays = [];
    var prepClients = [];
    const prepInitialPolicyState = () =>{
        
        var count = 1;
        selectedPolicy.fullbackup.map(day => {
            var temp = {name:day, id:count};
            editPolicyInitialState.oldbackuprundays.push(temp);
            count ++;
        })
        count = 1;
        //console.log(policy);
        if (selectedPolicy.Clients != null){
            selectedPolicy.Clients.map(client => {
                console.log("client",client);
                var temp = {name:client[0], id:count};
                prepClients.push(temp);
                //editPolicyInitialState.oldclients.push(temp);
                count ++;
            })
       }
 

        // dispatch({type:"field", field:"policyname", value: selectedPolicy.Name});
        // dispatch({type:"field", field:"retention", value:selectedPolicy.retention});
        // dispatch({type:"field", field:"oldbackuprundays", value:prepDays});
        // dispatch({type:"field", field:"oldclients", value:prepClients});
        // console.log("policyname: ", selectedPolicy.Name, oldbackuprundays);
      //  dispatch({type:"setTime", value:policy.})
    }
    // const updatePolicy = async(bodymesage) =>{
    //     var clientIDs = [];
    //     console.log(bodymesage.clients);
    //     console.log(Clients);
    //     var clientnames = bodymesage.clients;
    //     var i = 0;
    //     var j =0;
    //     Clients.map(client =>{
    //         bodymesage.clients.map(msg =>{
    //             if(client.clientname == msg){
    //                 var temp = [client.clientname, client.id];
    //                 clientIDs.push(temp);
    //                 console.log(client.clientname);
    //             }
    //         })
            
            
    //     })
    //     bodymesage.clients = clientIDs
    //     console.log("clientID's: ", bodymesage);
    //     try{
    //         const data = await fetch('/api/policies/update', {method:"PUT", body:JSON.stringify(bodymesage)});
    //         const items = await data.json();
    //         console.log(items);
            
    //     }catch{
            
    //     }
    //     close();
       
    // }

    return(
        <div>
            <EditPolicyConfigurer close={close} policy={selectedPolicy} sclients={prepClients}/>

            {/* <PolicyConfigurer  clientsList={clientsPlaceholder} apiType={"PUT"} apiCall={updatePolicy} close={close}/> */}
        </div>
        

    )   

}


function NewPolicy({close}){
    useEffect(() =>{
        listClients();

    },[]);
    var clientsPlaceholder = [];
    const [clients, setClients] = useState([]);
    //var clients = [];
    const listClients = async() =>{
        //console.log(clientsplaceholder.length);
        
            const data = await fetch('/api/clients/list', {method:"GET"});
            const items = await data.json();
            //clients = items;
            console.log(typeof items); 
            var policies = [];
            var count = 1;
            setClients(items);
            items.map(client =>{
                var picker = {name: client.clientname, id: count};
                clientsPlaceholder.push(picker);
                var c = client;
                
               // clients.push(c);
                count ++;
            })
            console.log(clientsPlaceholder);       
    }
    
    const createPolicy = async(bodymesage) =>{
        var clientIDs = [];
        console.log(bodymesage.clients);
        console.log(clients);
        var clientnames = bodymesage.clients;
        var i = 0;
        var j =0;
        clients.map(client =>{
            bodymesage.clients.map(msg =>{
                if(client.clientname == msg){
                    var temp = [client.clientname, client.id];
                    clientIDs.push(temp);
                    console.log(client.clientname);
                }
            })
            
            
        })
        bodymesage.clients = clientIDs
        console.log("clientID's: ", bodymesage);
        try{
            const data = await fetch('/api/policies/create', {method:"POST", body:JSON.stringify(bodymesage)});
            const items = await data.json();
            console.log(items);
            
        }catch{
            
        }
        close();
       
    }

    return(
        <div>

            <PolicyConfigurer  clientsList={clientsPlaceholder} apiType={"POST"} apiCall={createPolicy} close={close}/>
        </div>
        

    )   
}


function PolicyOptions({id}){
    const dispatch = useContext(PolicyDispatchContext);
    const deletePolicy = () =>{
        fetch(`/api/policies/delete/${id}`, {method:"DELETE"}).then(response => {
           console.log(response);
        
        })
    }
    return(
        <div className="dropdown">
            <button className="backup-options">...</button>
            <div className="dropdowncontent">
            {/* <Link to={`/policy/${id}`} className="dropdown-item">View Policies</Link>
            <Link to={`/backup/${id}`} className="dropdown-item">View Backups</Link> */}
            <li className="dropdown-item" >View Clients</li>
            <li className="dropdown-item" onClick={() =>{
                dispatch({type:"selectPolicyForEditing", payload:id})
            }}>Edit Policy</li>
            <li className="dropdown-item" onClick={() =>{
                deletePolicy()
                dispatch({type: 'hidePolicy', payload:id})
            }}>Delete Policy</li>
            <li className="dropdown-item">Run Policy</li>
            {/* <li onClick={deleteClient} className="dropdown-item">Delete Client</li> */}
            
            </div>
    
        </div>
    );
}