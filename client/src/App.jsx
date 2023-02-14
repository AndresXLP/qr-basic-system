import React, { useState } from 'react';
import axios from "axios";
import { Button, TextField } from "@mui/material";

axios.defaults.baseURL = "http://localhost:8080/api/qr-system"
const App = () => {
  const [content, setContent] = useState("");
  const [image, setImage] = useState(null)

  const handleChange = (e) => {
    e.preventDefault()
    setContent(e.target.value)
  }
  const handleSubmit = async () => {
    try {
      const response = await axios.post("/generate", { content })
      setImage(response.data)
    } catch (e) {
      console.log(e.response)
    }
  }
  return (
    <main>
      <center>
        <div>
          <TextField style={ { width: '500px', marginTop: '1rem' } } label="Content for the QR Code / URL or TEXT"
                     id="fullWidth"
                     value={ content }
                     onChange={ handleChange }/>
        </div>
        <div>
          <Button style={ { marginTop: '.5rem' } } variant="contained" onClick={ handleSubmit }>Generate QR
            Code</Button>
        </div>
        <div className="qr-code">
          { image && <img src={ `data:image/png;base64,${ image }` } /> }
        </div>
      </center>
    </main>
  );
};

export default App;