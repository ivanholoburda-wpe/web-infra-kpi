import React from 'react';
import {Formik, Form, Field, ErrorMessage} from 'formik';
import Button from './Button';
import * as Yup from 'yup';
import type Site from "../interfaces/site.interface.ts";

interface SiteFormProps {
    onSave: (site: Site) => void;
    site?: Site;
}

const SiteForm: React.FC<SiteFormProps> = ({onSave, site}) => {
    const validationSchema = Yup.object({
        Name: Yup.string().required('Site Name is required'),
        // Url: Yup.string().url('Must be a valid URL').required('URL is required'),
        HttpStatus: Yup.number().min(100).max(599).required('HTTP Status is required'),
        ResponseTime: Yup.number().min(0).required('Response Time is required'),
    });

    const initialValues: Partial<Site> = site || {Name: '', Url: '', HttpStatus: 0, ResponseTime: 0};

    const handleSubmit = (values: Partial<Site>) => {
        const submitValues: Site = {...values} as Site;
        if (!submitValues.ID) {
            const newSite: Omit<Site, 'ID'> = {
                Name: submitValues.Name!,
                Url: submitValues.Url!,
                HttpStatus: submitValues.HttpStatus!,
                ResponseTime: submitValues.ResponseTime!
            };
            onSave(newSite as Site);
        } else {
            onSave(submitValues);
        }
    };

    return (
        <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
        >
            {({isSubmitting}) => (
                <Form className="site-form">
                    <div>
                        <Field
                            type="text"
                            name="Name"
                            placeholder="Site Name"
                            className="input"
                        />
                        <ErrorMessage name="Name" component="div" className="error"/>
                    </div>
                    <div>
                        <Field
                            type="text"
                            name="Url"
                            placeholder="URL"
                            className="input"
                        />
                        <ErrorMessage name="Url" component="div" className="error"/>
                    </div>
                    <div>
                        <Field
                            type="number"
                            name="HttpStatus"
                            placeholder="HTTP Status"
                            className="input"
                        />
                        <ErrorMessage name="HttpStatus" component="div" className="error"/>
                    </div>
                    <div>
                        <Field
                            type="number"
                            name="ResponseTime"
                            placeholder="Response Time (ms)"
                            className="input"
                        />
                        <ErrorMessage name="ResponseTime" component="div" className="error"/>
                    </div>
                    <Button type="submit" variant={site ? 'update' : 'add'} disabled={isSubmitting}>
                        {site ? 'Update' : 'Add'} Site
                    </Button>
                </Form>
            )}
        </Formik>
    );
};

export default SiteForm;