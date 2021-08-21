import aspida from '@aspida/fetch';
import api from '../api/$api';

const fetchConfig = {
  baseURL: import.meta.env.PROD
    ? '/api'
    : `${window.location.protocol}//${window.location.hostname}:${import.meta.env.VITE_API_PORT}/api`,
  throwHttpErrors: true,
};

const apiClient = api(aspida(fetch, fetchConfig));
export default apiClient;
