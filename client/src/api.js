import http from './http-common';

const getRam = () => {
    return http.get('/modulo_ram');
}

const getAllData = () => {
    return http.get('/');
}

const getOneDose = () => {
    return http.get('/one_dose');
}

const getTwoDose = () => {
    return http.get('/two_dose');
}

const getNinos = () => {
    return http.get('/ninos');
}

const getAdolescentes = () => {
    return http.get('/adolescentes');
}

const getJovenes = () => {
    return http.get('/jovenes');
}

const getAdultos = () => {
    return http.get('/adultos');
}

const getVejez = () => {
    return http.get('/vejez');
}

const getNombres = () => {
    return http.get('/nombres');
}

export default {
    getRam,
    getAllData,
    getOneDose,
    getTwoDose,
    getNinos,
    getAdolescentes,
    getJovenes,
    getAdultos,
    getVejez,
    getNombres
}