import {useState, useEffect} from 'react'
import './App.scss'
import {createOrUpdateSite, deleteSite, fetchSites} from "./services/api.ts";
import SiteForm from "./components/SiteForm.tsx";
import SiteList from "./components/SiteList.tsx";
import type Site from "./interfaces/site.interface.ts";

function App() {
    const [sites, setSites] = useState<any[]>([]);
    const [selectedSite, setSelectedSite] = useState<any>(null);

    useEffect(() => {
        fetchSites().then(setSites);
    }, []);

    const handleCreateOrUpdate = async (site: Partial<Site>) => {
        await createOrUpdateSite(site);
        fetchSites().then(setSites);
        setSelectedSite(null);
    };

    const handleDelete = async (id: number) => {
        await deleteSite(id);
        fetchSites().then(setSites);
    };

    return (
        <div className="container">
            <h1 className="title">Site Monitor</h1>
            <SiteForm onSave={handleCreateOrUpdate} site={selectedSite}/>
            <SiteList sites={sites} onEdit={setSelectedSite} onDelete={handleDelete}/>
        </div>
    );
}

export default App
