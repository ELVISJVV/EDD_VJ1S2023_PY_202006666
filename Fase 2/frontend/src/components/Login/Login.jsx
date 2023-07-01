import style from './Login.module.css';
import imageLogin from '../../img/login.svg';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {   
    const navigate = useNavigate();

    const [userData, setUserData] = useState({
        user: '',
        password: ''
    });

    const handleChange = (event) => {
        const property = event.target.name;
        const value = event.target.value;

        setUserData({
            ...userData,
            [property]: value
        });
    }

    const handleSUbmit = (event) => {
        event.preventDefault();
        fetch('http://localhost:3001/login', {
            method: 'POST',
            body: JSON.stringify({
                Username: userData.user,
                Password: userData.password
            }),
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validate(data))
    }

    const validate = (data) => {
        console.log(data)
        if (data.status === 'admin'){
            navigate('/admin')
        }else if(data.status === 'employee'){
            // navigate('/user')
        }else{
            alert('Usuario o contraseña incorrecta')
        }
    }
    
    return (


        <div className={style.login}>
            <figure className={style.login__picture}>

                <img src={imageLogin} alt="" className={style.login__img} />
            </figure>

            <form onSubmit={handleSUbmit} className={style.login__form} >
                <h2
                    className={style.login__title}
                >INICIAR SESIÓN</h2>

                <input
                    type="text"
                    name="user"
                    className={style.login__input}
                    placeholder="Ingresa tu usuario"
                    onChange={handleChange} />

                <input
                    type="password"
                    name="password"
                    className={style.login__input}
                    placeholder="Ingresa tu contraseña"
                    onChange={handleChange} />

                <button
                    type="submit"
                    className={style.login__cta}
                >INGRESAR </button>
            </form>
        </div>

    )
}

export default Login;