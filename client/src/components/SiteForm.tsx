import React from 'react';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import Button from './Button';
import * as Yup from 'yup';
import type Site from "../interfaces/site.interface.ts";

interface SiteFormProps {
    onSave: (site: Site) => void;
    site?: Site;
}

const SiteForm: React.FC<SiteFormProps> = ({ onSave, site }) => {
    const validationSchema = Yup.object({
        name: Yup.string().required('Site Name is required'),
        url: Yup.string().url('Must be a valid URL').required('URL is required'),
        http_status: Yup.number().min(100).max(599).required('HTTP Status is required'),
        response_time: Yup.number().min(0).required('Response Time is required'),
    });

    // Initialize form values, ensuring all fields are populated when site is provided
    const initialValues: Partial<Site> = {
        id: site?.id || undefined,
        name: site?.name || '',
        url: site?.url || '',
        http_status: site?.http_status || 0,
        response_time: site?.response_time || 0,
    };

    const handleSubmit = (values: Partial<Site>) => {
        // Ensure all required fields are included in the submission
        const submitValues: Site = {
            id: values.id || undefined, // Include id if present (for updates)
            name: values.name!,
            url: values.url!,
            http_status: values.http_status!,
            response_time: values.response_time!,
        };
        onSave(submitValues);
    };

    return (
        <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
            enableReinitialize // Add this to reinitialize form when site prop changes
        >
            {({ isSubmitting }) => (
                <Form className="site-form">
                    <div>
                        <Field
                            type="text"
                            name="name"
                            placeholder="Site Name"
                            className="input"
                        />
                        <ErrorMessage name="name" component="div" className="error" />
                    </div>
                    <div>
                        <Field
                            type="text"
                            name="url"
                            placeholder="URL"
                            className="input"
                        />
                        <ErrorMessage name="url" component="div" className="error" />
                    </div>
                    <div>
                        <Field
                            type="number"
                            name="http_status"
                            placeholder="HTTP Status"
                            className="input"
                        />
                        <ErrorMessage name="http_status" component="div" className="error" />
                    </div>
                    <div>
                        <Field
                            type="number"
                            name="response_time"
                            placeholder="Response Time (ms)"
                            className="input"
                        />
                        <ErrorMessage name="response_time" component="div" className="error" />
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