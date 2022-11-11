import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";

function App() {
  let Elem = <h1> Hello </h1>
  return ( 
  <BrowserRouter>
    <Routes>
      <Route path="/" element ={ <Elem /> } />
    </Routes>
  </BrowserRouter>
  )
}

export default App;
