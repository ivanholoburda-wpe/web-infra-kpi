import React from 'react';
import Button from './Button';
import type Site from "../interfaces/site.interface.ts";

interface SiteListProps {
    sites: Site[];
    onEdit: (site: Site) => void;
    onDelete: (id: number) => void;
}

const SiteList: React.FC<SiteListProps> = ({sites, onEdit, onDelete}) => {
    return (
        <div className="site-list">
            {sites.map((site) => (
                <div key={site.ID} className="site-item">
                    <h2 className="site-name">{site.Name}</h2>
                    <p><strong>URL:</strong> {site.Url}</p>
                    <p><strong>HTTP Status:</strong> {site.HttpStatus}</p>
                    <p><strong>Response Time:</strong> {site.ResponseTime} ms</p>
                    <Button variant="edit" onClick={() => onEdit(site)}>
                        Edit
                    </Button>
                    <Button variant="delete" onClick={() => onDelete(site.ID)}>
                        Delete
                    </Button>
                </div>
            ))}
        </div>
    );
};

export default SiteList;