const { LOGIN } = require("./action-types")
import axios from 'axios';


export const login = (character) => {
    return {
        type: ADD_FAV,
        payload: character
    }

}
