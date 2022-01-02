import React from 'react';
import { Link } from 'react-router-dom';

import './styles/Navbar.css';

class Navbar extends React.Component {
  render() {
    return (
      <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
        <div className="container-fluid">
          <Link className="navbar-brand" to="/">Proyecto 2</Link>
          <div className="collapse navbar-collapse">
            <ul className="navbar-nav">
              <li className="nav-item">
                <Link className="nav-link active" to="/one-dose">Una Dosis</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link active" to="/two-dose">Dos Dosis</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link active" to="/all-data">Mongo data</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link active" to="/info-data">Redis data</Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    );
  }
}

export default Navbar;