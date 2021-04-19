import React from 'react';
import './App.css';
import {Link} from 'react-router-dom'

function nav(){
    const navStyle ={
        color:"white"
    }
    const pid = "all"
    return(
        <nav>
            <ul className="nav-links">
                <Link style={navStyle} className="nav-link-item" to="/client"><li>Clients</li></Link>
                <Link style={navStyle} className="nav-link-item" to={`/policy/${pid}`}><li>Policies</li></Link>
                <Link style={navStyle} className="nav-link-item" to="/backup"><li>Backups</li></Link>
                
                
                <li>Logout</li>
            </ul>
        </nav>
    );
}
export default nav;

