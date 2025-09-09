import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api/v1';

export const fetchSites = async () => {
    const response = await axios.get(`${BASE_URL}/sites`);
    return response.data;
};

export const createOrUpdateSite = async (site: any) => {
    const method = site.id ? 'put' : 'post';
    const url = site.id ? `${BASE_URL}/sites/${site.id}` : `${BASE_URL}/sites`;
    await axios({ method, url, data: site });
};

export const deleteSite = async (id: number) => {
    await axios.delete(`${BASE_URL}/sites/${id}`);
};