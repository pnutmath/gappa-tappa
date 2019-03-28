import React from "react";
import logo from '../../logo.svg';
import "./Header.scss";

const Header = () => (
    <div className="header">
        <img src={logo} className="App-logo" alt="logo" />
        <h2>Gappa-Tappa</h2>
    </div>
)

export default Header;