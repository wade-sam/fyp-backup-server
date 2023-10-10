import React, {useState} from 'react';
import PolicyItem from './components/policy-item';
const PolicyList = (props) => {
    console.log("HELLO");
    return(
        <div className="grid-container-list" >
            {props.policies.map(policy =>(
                <PolicyItem key={policy.id} policy={policy} />
                //setView(<PolicyItem key={policy.id} id={policy.id} name={policy.policyname} clients={policy.clients} type={policy.type} runs={policy.BackupRun} />)
                //<PolicyItem key={policy.id} id={policy.id} name={policy.policyname} clients={policy.clients} type={policy.type} backuprun={policy.BackupRun} setPolicy={}/>
            
        ))}
       </div>
    );

}

export default PolicyList