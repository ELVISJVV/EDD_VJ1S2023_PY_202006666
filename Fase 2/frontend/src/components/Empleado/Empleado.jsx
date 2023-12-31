import '../../css/Administrator.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export const Empleado = () => {
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/", "_self")
    }

    

    const aplicarFiltros = async (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/filtros", "_self")
    }

    const generarFacturas = async (e) => {
        e.preventDefault();
        window.open("/factura", "_self")
    }

    const verFacturas = async (e) => {
        e.preventDefault();
        window.open("/verfactura", "_self")
    }

    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                    <br />
                    <center><button className="w-50 btn btn-outline-primary" onClick={aplicarFiltros}>Aplicacion de Filtros</button></center>
                    <br />
                    <center><button className="w-50 btn btn-outline-primary" onClick={generarFacturas}>Generar Facturas</button></center>
                    <br />
                    <center><button className="w-50 btn btn-outline-primary" onClick={verFacturas}>Ver Facturas</button></center>
                    <br />
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br />
                    {/* <p className="mt-5 mb-3 text-muted">EDD 201700918</p> */}
                    <br />
                </form>
            </div>
        </div>
    );
}