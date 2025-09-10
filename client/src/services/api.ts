import axios from 'axios';
import type Site from "../interfaces/site.interface.ts";

const BASE_URL = 'http://localhost:8080/api/v1';

export const fetchSites = async () => {
    const response = await axios.get(`${BASE_URL}/sites`);
    return response.data;
};

export const createOrUpdateSite = async (site: Partial<Site>) => {
    const method = site.id ? 'put' : 'post';
    const url = site.id ? `${BASE_URL}/sites/${site.id}` : `${BASE_URL}/sites`;
    console.log(site);
    await axios({method, url, data: site});
};

export const deleteSite = async (id: number) => {
    await axios.delete(`${BASE_URL}/sites/${id}`);
};