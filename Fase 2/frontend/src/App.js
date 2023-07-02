import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './components/Login/Login';
import Administrator from './components/Administrator/Administrator'
import { Reportes } from './components/Reportes/Reportes';
import { Empleado } from './components/Empleado/Empleado';
import { Filtros } from './components/Filtros/Filtros';
import { GenerarFactura } from './components/GenerarFactura/GenerarFactura';
import { Factura } from './components/Facturas/Facturas';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/admin" element={<Administrator />} />
          <Route path="/reportes" element={<Reportes />} />
          <Route path="/empleado" element={<Empleado />} />
          <Route path="/filtros" element={<Filtros />} />
          <Route path="/factura" element={<GenerarFactura />} />
          <Route path="/verfactura" element={<Factura/>}/>
        </Routes>
      </Router>
      
    </div>
  );
}

export default App;
