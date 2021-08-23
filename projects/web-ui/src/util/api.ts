import aspida from '@aspida/fetch';
import def from '../api/$api';

const fetchConfig = {
  baseURL: import.meta.env.PROD
    ? '/api'
    : `${window.location.protocol}//${window.location.hostname}:${import.meta.env.VITE_API_PORT}/api`,
  throwHttpErrors: true,
};

const api = def(aspida(fetch, fetchConfig));
export default api;
