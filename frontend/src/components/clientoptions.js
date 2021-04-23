import React from 'react';
import { Link } from 'react-router-dom'

class ClientOptions extends React.Component{
    render(){
        const id = this.props.id;
        const onClick = this.props.onClick;
        console.log(id)

        const deleteClient = () =>{
             fetch(`http://localhost:8000/api/clients/delete/${id}`, {method:"DELETE"}).then(response => {
                console.log(response);
                onClick()
             })
         }
        return(
            <div className="dropdown">
            <button className="backup-options">...</button>
            <div className="dropdowncontent">
            <Link to={`/policy/${id}`} className="dropdown-item">View Policies</Link>
            <Link to={`/backup/${id}`} className="dropdown-item">View Backups</Link>
            <li onClick={deleteClient} className="dropdown-item">Delete Client</li>
            </div>
    
        </div>
        );
    }
}
export default ClientOptions;
