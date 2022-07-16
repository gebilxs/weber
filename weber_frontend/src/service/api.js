
import axios from 'axios';
axios.defaults.baseURL = "/api/";
axios.interceptors.request.use((config) => {
	let loginResult = JSON.parse(localStorage.getItem("loginResult"));
	if (loginResult) {
		const token = loginResult.token
		config.headers.Authorization = `Bearer ${token}`;
		config.headers['content-type'] = 'application/json'
	}
	return config;
}, (error) => {
	return Promise.reject(error);
});

axios.interceptors.response.use(
	response => {
		if (response.status === 200) {
			return Promise.resolve(response.data);
		} else {
			return Promise.reject(response.data);
		}
	},
	(error) => {
		console.log('error', error);
	}
);

export default axios;