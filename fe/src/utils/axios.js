import axios from 'axios';

const fetch = axios.create({
  baseURL: process.env.REACT_APP_BASE_API_URL,
});

fetch.interceptors.response.use(
  function (response) {
    return response.data;
  },
  function (error) {
    if (error.response) {
      console.log('Error response: ', error.response);
      console.log(error.response.data);
      console.log(error.response.status);
      console.log(error.response.headers);
    } else if (error.request) {
      console.log('Error request: ', error.request);
    } else {
      console.log('Error', error.message);
    }
    return Promise.reject(error);
  }
);

export { fetch };
