import React from 'react';
import Button from './Button';
import type Site from "../interfaces/site.interface.ts";

interface SiteListProps {
    sites: Site[];
    onEdit: (site: Site) => void;
    onDelete: (id: number) => void;
}

const SiteList: React.FC<SiteListProps> = ({sites, onEdit, onDelete}) => {

    console.log(sites);
    return (
        <div className="site-list">
            {sites.map((site) => (
                <div key={site.id} className="site-item">
                    <h2 className="site-name">{site.name}</h2>
                    <p><strong>URL:</strong> {site.url}</p>
                    <p><strong>HTTP Status:</strong> {site.http_status}</p>
                    <p><strong>Response Time:</strong> {site.response_time} ms</p>
                    {/*<Button variant="edit" onClick={() => onEdit(site)}>*/}
                    {/*    Edit*/}
                    {/*</Button>*/}
                    <Button variant="delete" onClick={() => onDelete(site.id)}>
                        Delete
                    </Button>
                </div>
            ))}
        </div>
    );
};

export default SiteList;