import React, {useState, useEffect, useContext} from 'react';
import CheckboxTree from 'react-checkbox-tree';
import "react-checkbox-tree/src/less/react-checkbox-tree.less";
import DirectoryScan from './DirectoryScan';

//import {ClientContext, fetchAll} from '../ClientContext';
const ClientDirectoryScanPage = (props) =>{

    useEffect(() =>{
        getScan();
      },[]);
    const [newClient, setNewClient] = useState(true);
    const [editView, setEditView] = useState(false);
    const [scanResult, setScanResult] = useState({});
    const [ignorePath, setIgnorePath] = useState([]);
   // const {clients} = useContext(ClientContext);
    const [client, setClient] = useState([]);
    const [dir, setData] = useState([]);
    const visible = props.toggle;
    var nodes = [];

    const [ready, setReady] = useState(false);

    
    // const findClient = () =>{
    //     var count = 0;
    //     console.log(clients.length);
    //     while (count <= clients.length){
    //        // console.log("clientdirectory",clients[0])
    //        console.log(count);
    //         if (clients[count].id == props.clientid){
    //             console.log("FOUND", clients[count]);
    //             setClient(clients[count]);
    //             console.log(client);
    //             break;
                
    //         } 
    //         count ++;
    //     }
        
    // }
    const getScan = async() =>{
        console.log("CLIENT", props.client);
        const data = await fetch(`http://localhost:8000/api/clients/scan/${props.client.consumerID}`, {method:"GET"});
        const scan = await data.json();
        //nodes = scan;
        //console.log(nodes);
        //var count;
        // for(count = 0; count < scan.length; count ++){

        // }
       // traverse(scan);
        
        nodes.push(scan);
        setData(nodes);
        //nodes = scan;
       
        console.log(nodes);
       // console.log(typeof nodes[0].children);
        setReady(!ready);
    }

    const [state, setState] = useState({
        checked: [],
        expanded: [],
        
    })

  

    const test = [{
        value: "mars",
        label: "Mars",
        children: [
          { value: 'phobos', label: 'Phobos', children:[{value: 'sam', label:'Sam'}]},
          { value: 'deimos', label: 'Deimos' },
        ],
      }];
      console.log(state.data);
      const t = nodes;
    // const test2 = [{
    //     value: "/home/sam/Documents",
    //     label: "Documents",
    //     children: [
    //         {
    //             value: "/home/sam/Documents/backup-server",
    //             label: "backup-server",
    //             children: [
    //                 {
    //                     value: "/home/sam/Documents/backup-server/fyp-backup-server",
    //                     label: "fyp-backup-server",
    //                     children: [
    //                         {
    //                             value: "/home/sam/Documents/backup-server/fyp-backup-server/frontend",
    //                             label: "frontend",
    //                             children: [
    //                                 {
    //                                     value: "/home/sam/Documents/backup-server/fyp-backup-server/frontend/public",
    //                                     label: "public",
    //                                     children: []
    //                                 },]
    //                             },]
    //                         },]
    //         },]
    // },]
    return(
        <React.Fragment>
        <div className="popup-box">
            <span className="close-icon" onClick={props.toggle}>x</span>
            <div className="box">
                <div>
                    <h1>Directory Scan Result</h1>
                    <div>
                        {ready &&
                        // <CheckboxTree
                        //     nodes={dir}
                        //    // key={state.data.value}
                        //     checked={state.checked}
                        //     expanded={state.expanded}
                        //     onCheck={checked => setState({ checked })}
                        //     onExpand={expanded => setState({ expanded })}
                        // />
                        <DirectoryScan scan={dir} />
      
                        }
                    </div>
                    <button>Save</button>   
                </div>

        
        
            
            </div>
            
        </div>
        </React.Fragment>
    )
}
export default ClientDirectoryScanPage