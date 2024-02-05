import React, { useState, createContext,useEffect } from 'react';

export const StoreContext = createContext(null);

export default function Context({children}) {
  const [response, setResponse] = useState('');
  const [scanResult, setScanResult] = useState(0);
  const [formDetails, setFormDetails] = useState({});

  useEffect(() => {
    if (response !== '') {
      localStorage.setItem('res', JSON.stringify(response));
    }
  }, [response]);
  
  useEffect(() => {
    const storedResponse = localStorage.getItem('res');
    if (storedResponse) {
      setResponse(JSON.parse(storedResponse));
    }
  }, []);
  const values = { response, setResponse, scanResult, setScanResult,formDetails,setFormDetails };

  return (
    <StoreContext.Provider value={values}>{children}</StoreContext.Provider>
  );
}


