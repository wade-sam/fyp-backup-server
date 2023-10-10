import React, {useState, useReducer, useContext, useEffect} from 'react';
import CheckboxTree from 'react-checkbox-tree';

import 'react-checkbox-tree/lib/react-checkbox-tree.css';

function clientsReducer(state, action){
    switch(action.type){
        case "addBackups":{
            console.log(action.payload)
            return{
                ...state,
                backups:action.payload,
            }
        }
        case 'popupfinished':{
            return{
                ...state,
                // newClient:false,
                // editScan: false,
            }
        }
        default:
            break;
    }
    

    
}
const BackupStateContext = React.createContext();
const BackupDispatchContext = React.createContext();
function setBackupState(backup){
    
    const returnBackup = {
        ID: backup.clientrunid,
        Name: backup.runname,
        ClientName: backup.clientname,
        PolicyName: backup.policyname,
        Status: backup.runstatus,
        totalfiles: backup.totalfiles,
        visible:true
    }
    return returnBackup
}
const initialState = {
    isLoading:true,
    error:'',
    backupSelected:{},
    backups: []
}
export default function BackupPage(){
    const [state,dispatch] = useReducer(clientsReducer, initialState);
    useEffect(() =>{
        fetchItems();
    },[]);

    const fetchItems = async() =>{
        const data = await fetch('/api/backups/list', {method:"GET"});
        const items = await data.json();
        console.log(items); 
        var backups = [];
        items.map(backup =>{
            var returnBackup = setBackupState(backup);
            backups.push(returnBackup);
            console.log(returnBackup);
        })
        dispatch({type:'addBackups', payload:backups})
        console.log(initialState.backups);
    }

    const toggleNewPopup = () => {
        dispatch({type:"popupfinished"});
        fetchItems();
    //  state.newClient = !state.newClient;
  }

    
  console.log(initialState.backups.length);

    return(
        <BackupDispatchContext.Provider value={dispatch}>
            <BackupStateContext.Provider value={state}>
            <div className="grid-container">
                    <div className="searchbox">
                  
                        <input type="text" ></input>
                        {/* <button className="newClient" onClick={() =>{dispatch({type:"newClient"})}}>New</button> */}
                        
                    </div>
                   
                
                    
                    <div className="grid-container-list">
                    {/* {state.newClient && <NewClient close={toggleNewPopup}/>}
                    {state.editScan && <EditDirectoryScan close={toggleNewPopup} client={state.clientSelected}/>} */}
                    {state.backups.map(item =>(
                      
                        item.visible && <BackupBox key={item.ID} id={item.ID} totalfiles={item.totalfiles}name={item.Name} clientname={item.ClientName} policyname={item.PolicyName} runs={item.backups} visible={item.visible}></BackupBox>
                    ))}
                    </div>
                </div>
            </BackupStateContext.Provider>
        </BackupDispatchContext.Provider>
    )
}

function BackupBox({id,name, clientname, policyname, status, totalfiles, visible}){
    if (visible == false){
        return null;
    }

    return(
        <div className="backup-item-grid-container">
            <p className="backup-label-1">Full</p>
                <div className="backup-item-name-1">
                <p>{clientname[0]}</p>
                </div>
                {/* <ClientOptions id={clientname}/> */}
                <p className="backup-item-policy">TotalFiles: {totalfiles}</p>
                <p className="backup-item-backup">Backups:</p>

        </div>
    )
}