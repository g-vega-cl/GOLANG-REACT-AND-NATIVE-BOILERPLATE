import logo from './logo.svg';
import './App.css';
import React, { useState, useEffect } from 'react';

function App() {
  const [fetchedData, setFetchedData] = useState("");

  useEffect(() => {
    const fetchSomeData = async () => {
      try {
        const response = await fetch('https://rezume-backend.herokuapp.com/');
        const json = await response.json();
        setFetchedData(json.Message);
      } catch (error) {
        console.error(error);
      }
    };
    fetchSomeData();
  }, []);
  return (
    <div className='App'>
      <header className='App-header'>
        <img src={logo} className='App-logo' alt='logo' />
        <p>{fetchedData}</p>
      </header>
    </div>
  );
}

export default App;
