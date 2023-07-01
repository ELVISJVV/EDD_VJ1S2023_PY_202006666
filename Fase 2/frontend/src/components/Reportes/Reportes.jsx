import React, { useState, useEffect } from 'react';
import '../../css/Administrator.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import { useNavigate } from 'react-router-dom';
export const Reportes = () => {
    const navigate = useNavigate();
    const [imagen, setImagen] = useState()
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        navigate('/admin')
    }

    const validar = (data) => {
        // console.log(data)
        setImagen(data.imagen.Imagenbase64)
    }

    const reporteGrafo = async (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-grafo', {
        })
            .then(response => response.json())
            .then(data => validar(data));
    }

    const reporteArbol = async (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol', {
        })
            .then(response => response.json())
            .then(data => validar(data));
    }

    const reporteBlockchain = async (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-bloque', {
        })
            .then(response => response.json())
            .then(data => validar(data));
    }

    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Reportes Administrador</h1>
                    <br />
                    <center><button className="w-50 btn btn-outline-primary" onClick={reporteGrafo}>Grafo</button></center>
                    <br />
                    <center><button className="w-50 btn btn-outline-primary" onClick={reporteArbol}>Arbol AVL</button></center>
                    <br />
                    <center><button className="w-50 btn btn-outline-primary" onClick={reporteBlockchain}>Facturas</button></center>
                    <br />
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br />
                    <center>
                        {imagen && (
                            <img
                                src={imagen}
                                width="350"
                                height="350"
                                alt="some value"
                            />
                        )}
                        {/* if (imagen === undefined) {
                            <img src={imagen} width="350" height="350" alt='some value' />
                        } */}
                    
                    {/* <img src={imagen} width="350" height="350" alt='some value' /> */}
                    </center>
                    <br />
                    
                    <br />
                </form>
            </div>
        </div>
    );
}