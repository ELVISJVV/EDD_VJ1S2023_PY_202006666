import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './components/Login/Login';
import Administrator from './components/Administrator/Administrator'
import { Reportes } from './components/Reportes/Reportes';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/admin" element={<Administrator />} />
          <Route path="/reportes" element={<Reportes />} />
        </Routes>
      </Router>
      
    </div>
  );
}

export default App;
