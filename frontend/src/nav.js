import React, { useState } from 'react';
import './App.css';
import {Link} from 'react-router-dom'
import NewClient from './components/NewClient';

const Nav = (props) =>{
    const navStyle ={
        color:"white"
    }
    const [isOpen, setIsOpen] = useState(false);

    const togglePopup = props.newClient

    const pid = "all"
    return(
        <nav>
            <ul className="nav-links">
                <div className="dropdown">
                    <Link style={navStyle} className="nav-link-item" to="/client">
                        <li>Clients</li>
                    </Link>
                    {/* <div className="dropdowncontent">
                        <Link className="dropdown-item" to="/client">View Clients</Link>
                        <li className="dropdown-item" onClick={togglePopup} >New Client</li> */}
                        {/* {isOpen && <NewClient toggle={togglePopup} />} */}
                        {/* {this.state.seen ? <NewClient toggle={this.newClient} /> : null} */}

                    
                    {/* </div> */}
                
                </div>
                
                <Link style={navStyle} className="nav-link-item" to={`/policy/${pid}`}><li>Policies</li></Link>
                <Link style={navStyle} className="nav-link-item" to="/backup"><li>Backups</li></Link>
                
                
                <li>Logout</li>
            </ul>
        </nav>
    );
}
export default Nav;

